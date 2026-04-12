# GLM CLI

A command-line interface for launching Claude Code with GLM (ChatGLM) settings via BigModel API, using temporary session-based configuration.

## Features

- 🚀 **Session-Based Launch**: Launch Claude with GLM settings temporarily (no persistent config changes)
- 🎯 **Model Selection**: Choose different GLM models at launch time (glm-5.1, glm-5, glm-4.7, glm-4.6, glm-4.5, glm-4.5-air, etc.)
- 🔀 **Flag Passthrough**: Pass any claude CLI flags directly through glm (e.g., `--allowedTools`, `--verbose`)
- ⚡ **YOLO Mode**: Skip permission prompts with `--yolo` flag for faster workflows
- 📦 **Auto-Install**: Install Claude Code with built-in npm dependency checking
- 🔄 **Auto-Update**: Check for and install updates with SHA256 checksum verification
- ⚙️ **Token Management**: Securely manage your authentication token
- 📋 **Model Listing**: View all available GLM models with `glm models`
- 🛠️ **Configuration Management**: View, set, and reset GLM settings with `glm config`

## Installation

### Quick Install (Recommended)

**Automatic Installer:**
```bash
curl -fsSL https://raw.githubusercontent.com/rodrigorodrigo/glm/main/install.sh | bash
```

**Alternative - Manual Quick Install:**
```bash
# Create user bin directory and download GLM CLI
mkdir -p ~/.local/bin
curl -L -o ~/.local/bin/glm "https://github.com/rodrigorodrigo/glm/releases/download/v1.4.0/glm-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m | sed 's/x86_64/amd64/')"
chmod +x ~/.local/bin/glm

# Add to PATH (one-time setup)
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

Both methods will:
- Detect your operating system and architecture
- Download the latest binary release
- Install to your user directory
- Set up PATH for easy access

### Manual Installation

#### Option 1: Download Pre-built Binary

1. Go to the [releases page](https://github.com/rodrigorodrigo/glm/releases)
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
git clone https://github.com/rodrigorodrigo/glm.git
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
glm  # Will prompt for token if not found
```

### Option 2: Manual Token Setup
```bash
glm token set              # Enter your token interactively
glm token set --token "X"  # Non-interactive (scripts, CI)
```

### Option 3: Environment Variable
```bash
export ANTHROPIC_AUTH_TOKEN="your_token_here"
glm
```

**Token Priority Order:**
1. Environment variable `ANTHROPIC_AUTH_TOKEN`
2. Config file `~/.glm/config.json`
3. Interactive prompt

## Usage

### Launch Claude with GLM (Primary Usage)

Launch Claude with the default model (glm-5.1):
```bash
glm
```

Launch Claude with a specific model:
```bash
glm --model glm-4.5-air
glm -m glm-4.5-air
```

Launch Claude in YOLO mode (skip permission prompts):
```bash
glm --yolo
glm --yolo --model glm-4.5-air
```

Pass additional flags directly to claude:
```bash
glm --allowedTools "Bash,Read,Write"
glm --verbose
glm --yolo --allowedTools "Bash,Read"
```

**How it works:**
- Sets temporary environment variables for the Claude session
- No persistent changes to Claude's configuration files
- Settings only apply to the launched Claude session
- To use Claude without GLM, just run `claude` directly

### List Available Models

```bash
glm models
```

Displays all available GLM models with the current default marked.

### Configuration Management

View current configuration:
```bash
glm config show
```

Change the default model:
```bash
glm config set --model glm-4.5-air
```

Change the API base URL:
```bash
glm config set --base-url https://custom.api.endpoint/v1
```

Change the authentication token:
```bash
glm config set --token "your_token_here"
```

Reset all settings to defaults:
```bash
glm config reset
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
glm token set --token "your_token"  # Non-interactive
```

View current token (masked):
```bash
glm token show
```

Clear stored token:
```bash
glm token clear
```

### Update GLM

Check for updates:
```bash
glm update --check
```

Update to latest version:
```bash
glm update
```

Update without confirmation:
```bash
glm update --force
```

### Help

Get help for any command:
```bash
glm --help
glm models --help
glm config --help
glm token --help
glm update --help
```

## Commands Reference

| Command | Description | Example |
|---------|-------------|---------|
| `glm` | Launch Claude with GLM (temporary config) | `glm --model glm-5` |
| `glm --yolo` | Launch with permission prompts skipped | `glm --yolo` |
| `glm --<flag>` | Pass any flag through to claude | `glm --allowedTools "Bash"` |
| `glm models` | List available GLM models | `glm models` |
| `glm config show` | Show current configuration | `glm config show` |
| `glm config set` | Set configuration values | `glm config set --model glm-5` |
| `glm config reset` | Reset config to defaults | `glm config reset` |
| `glm install claude` | Install Claude Code | `glm install claude` |
| `glm token set` | Set authentication token | `glm token set --token "X"` |
| `glm token show` | Show current token (masked) | `glm token show` |
| `glm token clear` | Clear stored token | `glm token clear` |
| `glm update` | Update GLM to latest version | `glm update` |
| `glm update --check` | Check for updates only | `glm update --check` |

## Available Models

- `glm-5.1` (default)
- `glm-5`
- `glm-4.7`
- `glm-4.6`
- `glm-4.5`
- `glm-4.5-air`
- Any other GLM model supported by BigModel API

## Configuration Files

The CLI manages the following files:
- `~/.glm/config.json` - Your authentication token, default model, and preferences

**Note:** GLM no longer modifies `~/.claude/settings.json`. All configuration is passed via temporary environment variables.

## How It Works

1. **Launch (`glm`)**: Launches Claude Code with temporary environment variables:
   - `ANTHROPIC_BASE_URL=https://open.bigmodel.cn/api/anthropic`
   - `ANTHROPIC_AUTH_TOKEN=<your_token>`
   - `ANTHROPIC_MODEL=<selected_model>`

2. **Session-Based**: Settings only exist for the launched Claude session. No persistent file modifications.

3. **Token Storage**: Your authentication token is securely stored in `~/.glm/config.json` for convenience.

4. **Install**: Checks for npm and installs Claude Code globally.

5. **Update**: Downloads and replaces the GLM binary with the latest version from GitHub, with SHA256 checksum verification.

## Example Workflow

```bash
# Install GLM CLI
curl -fsSL https://raw.githubusercontent.com/rodrigorodrigo/glm/main/install.sh | bash

# First time setup
glm install claude        # Install Claude Code
glm token set            # Enter your token securely

# Check available models
glm models

# Launch Claude with GLM (default model: glm-5.1)
glm

# Launch with specific model
glm --model glm-4.5-air

# Launch in YOLO mode (skip permission prompts)
glm --yolo

# Pass additional flags to claude
glm --allowedTools "Bash,Read,Write"

# Change your default model
glm config set --model glm-5

# View current configuration
glm config show

# Use Claude without GLM
claude

# Check for updates
glm update --check

# Update to latest version
glm update
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
curl -fsSL https://raw.githubusercontent.com/rodrigorodrigo/glm/main/install.sh -o install.sh
chmod +x install.sh
sudo ./install.sh
```

#### Binary not found for your platform
If no binary is available for your platform:
1. Check the [releases page](https://github.com/rodrigorodrigo/glm/releases) for available binaries
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

#### Claude still using default settings
The session-based configuration means:
- Settings only apply to Claude sessions launched via `glm`
- If you run `claude` directly, it uses default settings
- This is intentional - use `glm` to launch with GLM settings

#### Command not found after installation
If `glm` command is not found after installation:
1. Check if `/usr/local/bin` or `~/.local/bin` is in your PATH: `echo $PATH`
2. Add to PATH if missing (add to `.bashrc`, `.zshrc`, etc.):
   ```bash
   export PATH="$HOME/.local/bin:$PATH"
   ```
3. Restart your terminal or run: `source ~/.bashrc` (or `.zshrc`)

#### Update fails with permission error
If `glm update` fails with permission denied:
```bash
sudo glm update
```

## Migration from Previous Versions

If you're upgrading from version 1.0.x or 1.2.x:

### ⚠️ IMPORTANT: Remove Old Configuration File

**You MUST remove the old persistent configuration file to avoid conflicts:**

```bash
rm -f ~/.claude/settings.json
```

**Why this is required:**
- Version 1.0.x created a persistent `~/.claude/settings.json` file that made Claude always use GLM settings
- This conflicts with the current session-based approach
- **Without removing this file:** Running `claude` directly will still use GLM settings (not the default)
- **After removing this file:**
  - `glm` → Uses GLM settings (temporary, session-based)
  - `claude` → Uses default Claude settings (no GLM)

### Changes in v1.4.0:

1. **Removed commands**: `glm enable` and `glm disable` have been removed — use `glm` to launch and `claude` directly for non-GLM usage
2. **New commands**: `glm models`, `glm config show/set/reset`
3. **Non-interactive token**: `glm token set --token "X"` for scripts and CI
4. **Checksum verification**: Update downloads are verified with SHA256

## Changelog

### v1.4.0

#### 🚨 Breaking Changes
- **Module path changed** from `xqsit94/glm` to `rodrigorodrigo/glm`
- **Deprecated commands removed**: `glm enable` and `glm disable` no longer exist

#### ✨ New Features
- **`glm models`** — List all available GLM models with the current default highlighted
- **`glm config show`** — View current configuration (model, base URL, token status)
- **`glm config set`** — Change default model (`--model`), API base URL (`--base-url`), or token (`--token`)
- **`glm config reset`** — Reset all configuration to default values
- **Non-interactive token setting** — `glm token set --token "X"` for use in scripts and CI pipelines
- **SHA256 checksum verification** — Binary downloads during `glm update` are verified against checksums

#### 🐛 Fixes
- Updater now points to the correct fork repository (`rodrigorodrigo/glm`)
- All GitHub URLs updated to the correct repository

#### 📦 Internal
- Go module path updated to `rodrigorodrigo/glm`
- Removed deprecated enable/disable command code

### v1.2.1
- Default model changed to `glm-5.1`

### v1.2.0
- Flag passthrough support
- YOLO mode

### v1.1.0
- Session-based launch (no persistent config changes)
- Token management commands

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## Support

For issues and feature requests, please create an issue in the [repository](https://github.com/rodrigorodrigo/glm/issues).
