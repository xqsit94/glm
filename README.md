# GLM CLI

A command-line interface for managing GLM (ChatGLM) settings with Claude Code, enabling easy switching between different GLM models via BigModel API.

## Features

- 🚀 **Enable/Disable GLM**: Quickly configure Claude to use GLM models
- 🔧 **Model Management**: Switch between different GLM models (glm-4.5, glm-4.5-air, etc.)
- 📦 **Auto-Install**: Install Claude Code with built-in npm dependency checking
- ⚙️ **Easy Configuration**: Simple commands to manage your GLM settings

## Installation

### Quick Install (Recommended)

Install GLM CLI with a simple command:

```bash
# Create user bin directory and download GLM CLI
mkdir -p ~/.local/bin
curl -L -o ~/.local/bin/glm "https://github.com/xqsit94/glm/releases/download/v1.0.2/glm-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m | sed 's/x86_64/amd64/')"
chmod +x ~/.local/bin/glm

# Add to PATH (one-time setup)
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

**Alternative - Automatic Installer:**
```bash
curl -fsSL https://raw.githubusercontent.com/xqsit94/glm/main/install.sh | bash
```

Both methods will:
- Detect your operating system and architecture
- Download the latest binary release
- Install to your user directory
- Set up PATH for easy access

### Manual Installation

#### Option 1: Download Pre-built Binary

1. Go to the [releases page](https://github.com/xqsit94/glm/releases)
2. Download the binary for your platform:
   - macOS Intel: `glm-darwin-amd64`
   - macOS Apple Silicon: `glm-darwin-arm64`
   - Linux x64: `glm-linux-amd64`
   - Linux ARM64: `glm-linux-arm64`
3. Make it executable and move to PATH:
   ```bash
   chmod +x glm-*
   sudo mv glm-* /usr/local/bin/glm
   ```

#### Option 2: Build from Source

**Prerequisites:**
- Go 1.24 or later
- Your GLM API token

```bash
git clone https://github.com/xqsit94/glm.git
cd glm
go mod tidy
go build -o glm
sudo mv glm /usr/local/bin/
```

## Authentication Setup

The GLM CLI supports multiple ways to provide your Anthropic API token:

### Option 1: Interactive Setup (Recommended)
On first run, the CLI will automatically prompt you to set up your token:
```bash
glm enable  # Will prompt for token if not found
```

### Option 2: Manual Token Setup
```bash
glm token set  # Enter your token securely
```

### Option 3: Environment Variable
```bash
export ANTHROPIC_AUTH_TOKEN="your_token_here"
glm enable
```

**Token Priority Order:**
1. Environment variable `ANTHROPIC_AUTH_TOKEN`
2. Config file `~/.glm/config.json`
3. Interactive prompt

## Usage

### Enable GLM

Enable GLM with the default model (glm-4.5):
```bash
glm enable
```

Enable GLM with a specific model:
```bash
glm enable --model glm-4.5-air
glm enable -m glm-4.5-air
```

### Change Model

Change the model when GLM is already enabled:
```bash
glm set --model glm-4.5-air
glm set -m glm-4.5-air
```

### Disable GLM

Remove GLM configuration and restore default Claude settings:
```bash
glm disable
```

### Install Claude Code

Install Claude Code via npm (with automatic Node.js detection):
```bash
glm install claude
```

### Manage Authentication Token

Set your API token:
```bash
glm token set
```

View current token (masked):
```bash
glm token show
```

Clear stored token:
```bash
glm token clear
```

### Quick Start

Run GLM with default settings (enables GLM and starts Claude):
```bash
glm
```

### Help

Get help for any command:
```bash
glm --help
glm enable --help
glm set --help
glm install --help
```

## Commands Reference

| Command | Description | Example |
|---------|-------------|---------|
| `glm` | Quick start (enable + run claude) | `glm` |
| `glm enable` | Enable GLM settings for Claude | `glm enable --model glm-4.5` |
| `glm disable` | Disable GLM settings | `glm disable` |
| `glm set` | Change GLM model | `glm set --model glm-4.5-air` |
| `glm install claude` | Install Claude Code | `glm install claude` |
| `glm token set` | Set authentication token | `glm token set` |
| `glm token show` | Show current token (masked) | `glm token show` |
| `glm token clear` | Clear stored token | `glm token clear` |

## Available Models

- `glm-4.5` (default)
- `glm-4.5-air`
- Any other GLM model supported by BigModel API

## Configuration Files

The CLI manages the following files:
- `~/.claude/settings.json` - Claude configuration file
- `~/.glm/config.json` - Your authentication token

## How It Works

1. **Enable**: Creates/updates `~/.claude/settings.json` with BigModel API configuration
2. **Disable**: Removes the settings file and cleans up empty directories
3. **Set**: Updates the model in existing configuration without re-authentication
4. **Install**: Checks for npm and installs Claude Code globally

## Example Workflow

```bash
# Install GLM CLI
curl -fsSL https://raw.githubusercontent.com/xqsit94/glm/main/install.sh | bash

# First time setup
glm install claude
glm token set  # Enter your token securely
glm enable

# Or use the quick start
glm  # Enables GLM and starts Claude in one command

# Switch models as needed
glm set --model glm-4.5-air
glm set --model glm-4.5

# When done
glm disable
```

## Troubleshooting

### Installation Issues

#### curl not found
If you get a "curl not found" error:
- **macOS**: Install Xcode Command Line Tools: `xcode-select --install`
- **Linux**: Install curl: `sudo apt install curl` (Ubuntu/Debian) or `sudo yum install curl` (CentOS/RHEL)

#### Permission denied during installation
If the installer fails with permission errors:
```bash
# Download and run manually with explicit sudo
curl -fsSL https://raw.githubusercontent.com/xqsit94/glm/main/install.sh -o install.sh
chmod +x install.sh
sudo ./install.sh
```

#### Binary not found for your platform
If no binary is available for your platform:
1. Check the [releases page](https://github.com/xqsit94/glm/releases) for available binaries
2. Build from source using the manual installation instructions

### Runtime Issues

#### npm not found
If you get an npm error when running `glm install claude`:
1. Install Node.js from https://nodejs.org/
2. Restart your terminal
3. Run `glm install claude` again

#### Authentication token not found
Set up your token using any of these methods:
- `glm token set` (recommended)
- Set environment variable: `export ANTHROPIC_AUTH_TOKEN="your_token"`

#### Settings not taking effect
Try:
1. Restart your Claude Code session
2. Verify the settings file exists: `cat ~/.claude/settings.json`
3. Re-enable GLM: `glm disable && glm enable`

#### Command not found after installation
If `glm` command is not found after installation:
1. Check if `/usr/local/bin` is in your PATH: `echo $PATH`
2. Add to PATH if missing (add to `.bashrc`, `.zshrc`, etc.):
   ```bash
   export PATH="/usr/local/bin:$PATH"
   ```
3. Restart your terminal or run: `source ~/.bashrc` (or `.zshrc`)

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## Support

For issues and feature requests, please create an issue in the repository.
