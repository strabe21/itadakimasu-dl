# Itadakimasu-DL

An efficient anime downloader that scrapes anime from various websites (actually one) without using JavaScript-dependent browsers. This lightweight tool requires minimal resources.

## Features

- **Lightweight**: Operates without browser dependencies, making it resource-efficient
- **Fast Downloads**: Multiple concurrent downloads from different servers
- **Flexible Episode Selection**: Download specific episodes or ranges
- **Configurable**: Customizable file paths and naming conventions
- **Multiple Sources**: Currently supports AnimeFlv (more sources planned)
- **Multiple Servers**: Currently supports Streamtape and YourUpload

## Installation

### Prerequisites (for compilation only)

- Go 1.23.5 or higher

### Building from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/itadakimasu-dl
cd itadakimasu-dl

# Build the binary
go build

# Alternatively, run directly without building
go run .
```

## Usage

### Searching and Downloading Anime

```bash
# Basic search
itadakimasu-dl search "anime name"

# Download specific episodes
itadakimasu-dl search "anime name" --episodes="1-5"    # Episodes 1 to 5
itadakimasu-dl search "anime name" --episodes="1,5"    # Episodes 1 and 5
itadakimasu-dl search "anime name" --episodes="1-5,7,9-12"    # Episodes 1-5, 7, and 9-12

# Specify output directory
itadakimasu-dl search "anime name" --output="/path/to/directory"
```

### Shell Completion Scripts

Generate shell completion scripts for easier command-line usage:

```bash
itadakimasu-dl completion bash > /path/to/bash_completion.d/itadakimasu-dl
itadakimasu-dl completion zsh > ~/.zsh/completion/_itadakimasu-dl
itadakimasu-dl completion fish > ~/.config/fish/completions/itadakimasu-dl.fish
itadakimasu-dl completion powershell > itadakimasu-dl.ps1
```

## Configuration

The default configuration is stored in a JSON file:

```json
{
    "downloadPath": "test",
    "animePath": "[NAME]/[ASK]",
    "episodeFile": "[NUMBER]",
    "animeWebs": [
        "animeflv"
    ],
    "links": [
        {
            "stape": {
                "maxConcurrentDownloads": 20,
                "priority": 2
            },
            "yourupload": {
                "maxConcurrentDownloads": 20,
                "priority": 1
            }
        }
    ]
}
```

### Configuration Options

- **downloadPath**: Base directory for all downloads
- **animePath**: Directory structure for anime
  - `[NAME]`: Will be replaced with the anime name
  - `[ASK]`: Prompts for input unless overridden by `--output` flag
- **episodeFile**: Filename format for episodes
  - `[NUMBER]`: Will be replaced with the episode number
  - Example: Setting this to `[NAME]-[NUMBER]` will save files as `AnimeName-1.mp4`
- **animeWebs**: List of supported anime websites to scrape from
- **links**: Server configuration
  - **maxConcurrentDownloads**: Maximum simultaneous downloads per server
  - **priority**: Server priority (lower numbers = higher priority)

## How It Works

Itadakimasu-DL intelligently distributes download tasks among configured servers based on their priority and availability. If a server is busy, it will try another or wait until one becomes available.

## Supported Platforms

- Linux
- macOS
- Windows

## Future Plans

- Add support for "link" usage to download episodes from specific URLs.
- Add support for more anime websites
- Add support for more download servers
- Implement a watch mode for automatic new episode downloads

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
