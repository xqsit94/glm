# GLM v1.4.0 Release Notes

## 🚨 Breaking Changes

- **Module path changed**: Go module path moved from `xqsit94/glm` to `rodrigorodrigo/glm`. If you import this module, update your imports.
- **Deprecated commands removed**: `glm enable` and `glm disable` have been removed. Use `glm` to launch Claude with GLM settings, or run `claude` directly for default settings.

## ✨ New Features

### `glm models` — List Available Models
View all available GLM models with the current default highlighted:
```bash
$ glm models
Available GLM models:

  * glm-5.1 (default)
    glm-5
    glm-4.7
    glm-4.6
    glm-4.5
    glm-4.5-air
```

### `glm config` — Configuration Management
A new `config` command with subcommands for managing GLM settings:

```bash
# View current configuration
glm config show

# Change default model
glm config set --model glm-4.5-air

# Change API base URL
glm config set --base-url https://custom.endpoint/v1

# Change authentication token
glm config set --token "your_token"

# Reset everything to defaults
glm config reset
```

### Non-Interactive Token Setting
Set your token without interactive prompts — useful for scripts, CI/CD, and automation:
```bash
glm token set --token "your_token_here"
```

### SHA256 Checksum Verification
Binary downloads during `glm update` are now verified against SHA256 checksums to ensure integrity.

## 🐛 Fixes

- **Updater repository**: The updater now correctly points to `rodrigorodrigo/glm` for fetching releases
- **All URLs updated**: Install scripts, release links, and documentation now reference the correct repository

## 📦 Full Command Reference

| Command | Description |
|---------|-------------|
| `glm` | Launch Claude with GLM settings |
| `glm --model <model>` | Launch with a specific model |
| `glm --yolo` | Launch with permission prompts skipped |
| `glm models` | List available GLM models |
| `glm config show` | View current configuration |
| `glm config set --model <m>` | Set default model |
| `glm config set --base-url <url>` | Set API base URL |
| `glm config set --token <t>` | Set authentication token |
| `glm config reset` | Reset config to defaults |
| `glm install claude` | Install Claude Code via npm |
| `glm token set` | Set token (interactive) |
| `glm token set --token "X"` | Set token (non-interactive) |
| `glm token show` | Show token (masked) |
| `glm token clear` | Clear stored token |
| `glm update` | Update to latest version |
| `glm update --check` | Check for available updates |

## ⬆️ Upgrading

```bash
glm update
```

Or download directly from the [releases page](https://github.com/rodrigorodrigo/glm/releases/tag/v1.4.0).

If upgrading from v1.0.x, remove the old config file:
```bash
rm -f ~/.claude/settings.json
```
