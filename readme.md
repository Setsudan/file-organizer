# File Organizer

File Organizer sorts files from your Downloads directory into category folders.
It supports both one-shot execution and continuous watch mode.

## Features

- Dual mode runtime: `once` and `watch`
- Safe file move behavior with name conflict protection
- Cross-device fallback (copy then remove if rename cannot cross filesystems)
- Temporary download file filtering
- Expanded categories: images, videos, audio, files, archives, executable, code, others
- Configurable source/target directories through JSON config

## Requirements

- [Go](https://go.dev/) 1.23.4 or newer
- Windows, macOS, or Linux

## Build

```bash
go build -o file-organizer ./cmd/file-organizer
```

## Usage

Run one pass:

```bash
go run ./cmd/file-organizer --mode=once
```

Run continuously:

```bash
go run ./cmd/file-organizer --mode=watch
```

Useful flags:

- `--config <path>`: optional config JSON
- `--source <path>`: override source directory
- `--interval 3s`: watch loop interval
- `--stable-for 5s`: minimum file age before processing
- `--verbose`: debug logging

## Default target folders

By default, all category folders are created inside `~/Downloads`:

- `~/Downloads/Downloaded Images`
- `~/Downloads/Downloaded Videos`
- `~/Downloads/Downloaded Audio`
- `~/Downloads/Downloaded Files`
- `~/Downloads/Downloaded Archives`
- `~/Downloads/Downloaded Executables`
- `~/Downloads/Downloaded Code`
- `~/Downloads/Downloaded Others`

## Optional config file

```json
{
  "sourceDir": "C:/Users/you/Downloads",
  "targets": {
    "images": "C:/Users/you/Downloads",
    "videos": "C:/Users/you/Downloads",
    "audio": "C:/Users/you/Downloads",
    "files": "C:/Users/you/Downloads",
    "archives": "C:/Users/you/Downloads",
    "executable": "C:/Users/you/Downloads",
    "code": "C:/Users/you/Downloads",
    "others": "C:/Users/you/Downloads"
  },
  "tempExtensions": [".crdownload", ".part", ".tmp", ".opdownload", ".download"]
}
```

## Development

```bash
make fmt
make vet
make test
make build
```

## License

MIT. See [LICENSE](LICENSE).
