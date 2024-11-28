# pingLogger

PingLogger is a Go-based program that monitors the reachability of specified targets via ICMP pings. It logs changes in reachability status to an SQLite database and displays updates in the terminal.

## Features

Load targets (name and address) from a JSON file.
Perform ICMP ping tests to check target reachability.
Log status changes (reachable/unreachable) to:
Terminal output.
An SQLite database.
Modular and testable design.

## Installation

### Prerequisites
* Go (version 1.19 or higher).
* SQLite3.
* Root or elevated permissions (for sending privileged ICMP pings).

### Clone the Repository
```bash
git clone https://github.com/33015/pingLogger.git
cd pingLogger
```

### Install Dependencies
```bash
go mod tidy
```

## Usage

Configure Targets
Create a targets.json file in the root directory:

```json
[
  {
    "name": "Google",
    "address": "google.com"
  },
  {
    "name": "Localhost",
    "address": "127.0.0.1"
  }
]
```

### Run the Program
go run main.go

## Testing
The program includes unit tests for core functionality.

Run all tests:

```bash
go test ./...
```

### Example Test Structure
* Config Tests: Validate target loading from targets.json.
* Database Tests: Ensure correct logging to SQLite.
* Ping Tests: Test the checkPing functionality with mocked and real environments.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request for improvements or bug fixes.

## Acknowledgments

Uses go-ping for ICMP ping functionality.
Built with Go's standard library for SQLite and JSON handling.