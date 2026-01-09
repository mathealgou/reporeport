# reporeport

![Go Version](https://img.shields.io/badge/Go-1.25%2B-00ADD8?logo=go)

Fast CLI that scans a repository, counts files and lines by extension, and infers the project type. Outputs a short, colored summary you can run before diving into unfamiliar codebases.

## Features

- Counts files and lines for common extensions, excluding vendor/lib folders by default
- Infers project type from dominant languages and presence of package manifests (Go, Python, Java, React/TSX, TypeScript, generic fallback)
- Shows percentage breakdown of lines by file type plus total elapsed time
- Optional flag to include `node_modules`, `vendor`, `.venv`, etc. in the report
- Cross-platform: simple build script for Linux and a Windows build helper

## Install

### Linux

```sh
sudo curl -sSL https://github.com/mathealgou/reporeport/releases/download/v1.0.0/reporeport -o /usr/local/bin/reporeport && sudo chmod +x /usr/local/bin/reporeport
```

### Windows

```powershell
Invoke-WebRequest -Uri "https://github.com/mathealgou/reporeport/releases/download/v1.0.0/reporeport.exe" -OutFile "C:\Windows\System32\reporeport.exe"
```

### Pre-built binaries:

Download the latest release from the [Releases](https://github.com/mathealgou/reporeport/releases)

### Clone and build locally:

Requirements: Go 1.25+

```sh
git clone https://github.com/mathealgou/reporeport.git
cd reporeport
go build -o reporeport main.go
```

Linux helper script (moves binary to `/usr/bin`, requires sudo):

```sh
chmod +x build.sh
./build.sh
```

Windows build (drops `dist/reporeport.exe`, requires Bash):

```sh
chmod +x build-windows.sh
./build-windows.sh
```

## Usage

From the repo root (or after placing the binary on your PATH):

```sh
reporeport [--include-libs]
```

Flags currently supported in code:

- `--help`: show CLI help
- `--include-libs`: count files inside lib/vendor/node_modules/.venv, etc.
- `--use-gitignore`: respect `.gitignore` rules when scanning files, skipping ignored files and directories.
- `--verbose`: Output the processing and analysis steps to the console (may slow down execution due to I/O overheard).

## Example output

```bash
Inferred Project Type: React Project
Percentage of Lines by File Type:
ts         | 32.73% | ██████████████████████████████
tsx        | 20.15% | ██████████████████
json       | 16.37% | ███████████████
css        | 13.66% | ████████████
md         | 10.89% | █████████
yml        | 5.02 % | ████
js         | 1.19 % | █

Report generated in: 6.208091ms
```

## How it works

- Walks the current directory tree to collect files
- Filters by allowed extensions and, unless `--include-libs` is set, skips common library/vendor paths
- Tallies counts and percentages per extension and by lines
- Infers a project type from the top extensions and presence of `package.json` (React/TSX and TS heuristics)
- Gathers project characteristics like docs, config files, etc.

## Contributing

Issues and PRs are welcome. Please keep the CLI fast, with minimal dependencies. Before sending changes, format with `go fmt ./...` and ensure a clean build with `go build ./...`.

## License

MIT License. See `LICENSE` file for details.
