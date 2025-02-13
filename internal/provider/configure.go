package provider

import (
	"time"

	"context"
	"fmt"
	"github.com/d-strobel/terraform-provider-windows/internal/generate/provider_windows"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/d-strobel/gowindows"
	"github.com/d-strobel/gowindows/connection/ssh"
	"github.com/d-strobel/gowindows/connection/winrm"
)

const (
	// SSH environment variables.
	envSSHUsername       string = "WIN_SSH_USERNAME"
	envSSHPassword       string = "WIN_SSH_PASSWORD"
	envSSHPrivateKey     string = "WIN_SSH_PRIVATE_KEY"
	envSSHPrivateKeyPath string = "WIN_SSH_PRIVATE_KEY_PATH"
	envSSHKnownHostsPath string = "WIN_SSH_KNOWN_HOSTS_PATH"
	envSSHPort           string = "WIN_SSH_PORT"
	envSSHInsecure       string = "WIN_SSH_INSECURE"

	// WinRM environment variables.
	envWinRMUsername string = "WIN_WINRM_USERNAME"
	envWinRMPassword string = "WIN_WINRM_PASSWORD"
	envWinRMPort     string = "WIN_WINRM_PORT"
	envWinRMTimeout  string = "WIN_WINRM_TIMEOUT"
	envWinRMInsecure string = "WIN_WINRM_INSECURE"
	envWinRMUseTLS   string = "WIN_WINRM_USE_TLS"

	// SSH default values.
	defaultSSHPort     int  = 22
	defaultSSHInsecure bool = false

	// WinRM default values.
	defaultWinRMPort     int           = 5986
	defaultWinRMTimeout  time.Duration = 0
	defaultWinRMInsecure bool          = false
	defaultWinRMUseTLS   bool          = true
)

// Configure sets up the provider client.
// This includes the connection to the Windows system via WinRM or SSH.
func (p *WindowsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data provider_windows.WindowsModel
	var client *gowindows.Client

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	// Add log masking for sensitive fields.
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "Password", "PrivateKey")

	// Check the WinRM config and setup the connection.
	if !data.Winrm.IsNull() {
		tflog.Info(ctx, "Configure Windows remote client with WinRM connection")

		config := &winrm.Config{}

		// Endpoint
		config.Host = data.Endpoint.ValueString()

		// Username
		config.Username = data.Winrm.Username.ValueString()
		if data.Winrm.Username.IsNull() {
			config.Username = os.Getenv(envWinRMUsername)
		}

		// Password
		config.Password = data.Winrm.Password.ValueString()
		if data.Winrm.Password.IsNull() {
			config.Password = os.Getenv(envWinRMPassword)
		}

		// Port
		config.Port = defaultWinRMPort
		if !data.Winrm.Port.IsNull() {
			config.Port = int(data.Winrm.Port.ValueInt64())
		} else if os.Getenv(envWinRMPort) != "" {
			winrmPort, err := strconv.Atoi(os.Getenv(envWinRMPort))
			if err != nil {
				resp.Diagnostics.AddError("Environment variable conversion error", fmt.Sprintf("Failed to convert environment variable '%s' to integer. Error: %s", envWinRMPort, err))
			}
			config.Port = winrmPort
		}

		// Timeout
		config.Timeout = defaultWinRMTimeout
		if !data.Winrm.Timeout.IsNull() {
			config.Timeout = time.Duration(data.Winrm.Timeout.ValueInt64())
		} else if os.Getenv(envWinRMTimeout) != "" {
			winrmTimeout, err := strconv.Atoi(os.Getenv(envWinRMTimeout))
			if err != nil {
				resp.Diagnostics.AddError("Environment variable conversion error", fmt.Sprintf("Failed to convert environment variable '%s' to integer. Error: %s", envWinRMTimeout, err))
			}
			config.Timeout = time.Duration(winrmTimeout)
		}

		// UseTLS (https)
		config.UseTLS = defaultWinRMUseTLS
		if !data.Winrm.UseTls.IsNull() {
			config.UseTLS = data.Winrm.UseTls.ValueBool()
		} else if os.Getenv(envWinRMUseTLS) != "" {
			winrmUseTls, err := strconv.ParseBool(os.Getenv(envWinRMUseTLS))
			if err != nil {
				resp.Diagnostics.AddError("Environment variable conversion error", fmt.Sprintf("Failed to convert environment variable '%s' to integer. Error: %s", envWinRMUseTLS, err))
			}
			config.UseTLS = winrmUseTls
		}

		// Insecure
		config.Insecure = defaultWinRMInsecure
		if !data.Winrm.Insecure.IsNull() {
			config.Insecure = data.Winrm.Insecure.ValueBool()
		} else if os.Getenv(envWinRMInsecure) != "" {
			winrmInsecure, err := strconv.ParseBool(os.Getenv(envWinRMInsecure))
			if err != nil {
				resp.Diagnostics.AddError("Environment variable conversion error", fmt.Sprintf("Failed to convert environment variable '%s' to bool. Error: %s", envWinRMInsecure, err))
			}
			config.UseTLS = winrmInsecure
		}

		tflog.Debug(ctx, "Setup WinRM connection with the following config", map[string]interface{}{
			"Host":     config.Host,
			"Port":     config.Port,
			"Username": config.Username,
			"Password": config.Password,
			"Timeout":  config.Timeout,
			"UseTLS":   config.UseTLS,
		})

		// Connect to the Windows system.
		conn, err := winrm.NewConnection(config)
		if err != nil {
			resp.Diagnostics.AddError("Unable to setup connection via WinRM", err.Error())
			return
		}

		// Setup the gowindows client.
		client = gowindows.NewClient(conn)
	}

	// Check the SSH config and setup the connection.
	if !data.Ssh.IsNull() {
		tflog.Info(ctx, "Configure Windows remote client with SSH connection")

		config := &ssh.Config{}

		// Endpoint
		config.Host = data.Endpoint.ValueString()

		// Username
		if !data.Ssh.Username.IsNull() {
			config.Username = data.Ssh.Username.ValueString()
		} else if os.Getenv(envSSHUsername) != "" {
			config.Username = os.Getenv(envSSHUsername)
		}

		// Password
		if !data.Ssh.Password.IsNull() {
			config.Password = data.Ssh.Password.ValueString()
		} else if os.Getenv(envSSHPassword) != "" {
			config.Password = os.Getenv(envSSHPassword)
		}

		// Private Key
		if !data.Ssh.PrivateKey.IsNull() {
			config.PrivateKey = data.Ssh.PrivateKey.ValueString()
		} else if os.Getenv(envSSHPrivateKey) != "" {
			config.PrivateKey = os.Getenv(envSSHPrivateKey)
		}

		// Private Key path
		if !data.Ssh.PrivateKeyPath.IsNull() {
			config.PrivateKeyPath = data.Ssh.PrivateKeyPath.ValueString()
		} else if os.Getenv(envSSHPrivateKeyPath) != "" {
			config.PrivateKeyPath = os.Getenv(envSSHPrivateKeyPath)
		}

		// Port
		config.Port = defaultSSHPort
		if !data.Ssh.Port.IsNull() {
			config.Port = int(data.Ssh.Port.ValueInt64())
		} else if os.Getenv(envSSHPort) != "" {
			sshPort, err := strconv.Atoi(os.Getenv(envSSHPort))
			if err != nil {
				resp.Diagnostics.AddError("Environment variable conversion error", fmt.Sprintf("Failed to convert environment variable '%s' to integer. Error: %s", envSSHPort, err))
			}
			config.Port = sshPort
		}

		// Insecure
		config.Insecure = defaultSSHInsecure
		if !data.Ssh.Insecure.IsNull() {
			config.Insecure = data.Ssh.Insecure.ValueBool()
		} else if os.Getenv(envSSHInsecure) != "" {
			sshInsecure, err := strconv.ParseBool(os.Getenv(envSSHInsecure))
			if err != nil {
				resp.Diagnostics.AddError("Environment variable conversion error", fmt.Sprintf("Failed to convert environment variable '%s' to bool. Error: %s", envSSHInsecure, err))
			}
			config.Insecure = sshInsecure
		}

		// Known hosts path
		if !data.Ssh.KnownHostsPath.IsNull() {
			config.KnownHostsPath = data.Ssh.KnownHostsPath.ValueString()
		} else if os.Getenv(envSSHKnownHostsPath) != "" {
			config.KnownHostsPath = os.Getenv(envSSHKnownHostsPath)
		}

		tflog.Debug(ctx, "Setup SSH connection with the following config", map[string]interface{}{
			"Host":           config.Host,
			"Port":           config.Port,
			"Username":       config.Username,
			"Password":       config.Password,
			"PrivateKey":     config.PrivateKey,
			"PrivateKeyPath": config.PrivateKeyPath,
			"KnownHostsPath": config.KnownHostsPath,
			"Insecure":       config.Insecure,
		})

		// Connect to the Windows system.
		conn, err := ssh.NewConnection(config)
		if err != nil {
			resp.Diagnostics.AddError("Unable to setup connection via SSH", err.Error())
			return
		}

		// Setup the gowindows client.
		client = gowindows.NewClient(conn)
	}

	// Do not setup client if any errors occur
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the client in the provider.
	resp.DataSourceData = client
	resp.ResourceData = client
}
