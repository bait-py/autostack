# AutoStack

**Professional Docker stack generator for rapid development environment setup.**

AutoStack is a command-line tool that generates production-ready Docker Compose configurations with interactive customization. Ideal for developers who need quick, reproducible development environments.

## Overview

* Interactive configuration of environment variables and ports
* Multiple pre-configured stacks
* Sensible defaults with easy customization

## Features

* Generates pre-configured Docker Compose stacks with a single command
* Includes all necessary config files and documentation
* Auto-start capability for quick deployment
* Clean, generated documentation for each stack

## Installation

### Binary Installation (Recommended)

```bash
# Download and install latest version
curl -L https://github.com/YOUR_USERNAME/autostack/releases/latest/download/autostack -o autostack
chmod +x autostack
sudo mv autostack /usr/local/bin/
```

Or install for current user only:

```bash
mv autostack ~/.local/bin/
```

### Build from Source

Requires Go 1.22+

```bash
git clone https://github.com/YOUR_USERNAME/autostack.git
cd autostack
go build -o autostack
sudo mv autostack /usr/local/bin/
```

## Quick Start

### List available stacks

```bash
autostack list
```

### Create a stack with defaults

```bash
autostack create lamp
```

You will be prompted to configure environment variables and ports. Press Enter to accept defaults or type your values.

### Start the created stack

```bash
cd lamp-stack
docker-compose up -d
```

## Available Stacks

| Stack             | Command                                                    | Services                                | Default Ports                                           | Use Case                                 |
| ----------------- | ---------------------------------------------------------- | --------------------------------------- | ------------------------------------------------------- | ---------------------------------------- |
| **LAMP**          | `autostack create lamp`                                    | Apache + PHP 8.2, MySQL 8.0, phpMyAdmin | 8080 (Web), 3306 (MySQL), 8081 (phpMyAdmin)             | PHP web development, WordPress, Laravel  |
| **MariaDB**       | `autostack create mariadb`                                 | MariaDB, phpMyAdmin                     | 3306 (MariaDB), 8080 (phpMyAdmin)                       | Database development, MySQL alternative  |
| **Observability** | `autostack create observability` or `autostack create obs` | Prometheus, Grafana, Node Exporter      | 9090 (Prometheus), 3000 (Grafana), 9100 (Node Exporter) | System monitoring, metrics visualization |

---

## Usage

### Basic Usage

Create a stack with interactive prompts:

```bash
autostack create lamp
```

Follow prompts to customize environment variables and ports or press Enter to use defaults.

### Managing stacks

* Start a stack:

  ```bash
  cd lamp-stack
  docker-compose up -d
  ```

* Stop a stack:

  ```bash
  cd lamp-stack
  docker-compose down
  ```

* View logs:

  ```bash
  cd lamp-stack
  docker-compose logs -f
  ```

* Remove stack and data:

  ```bash
  cd lamp-stack
  docker-compose down -v
  rm -rf ../lamp-stack
  ```

## Configuration Details

### During stack creation, configurable options include:

* Environment variables (e.g., database credentials)
* Service ports for each container
* Option to auto-start stack after creation

All variables have sensible defaults and can be customized or accepted by pressing Enter.

## Troubleshooting

### Port Already in Use

If you get port conflicts:

```bash
sudo lsof -i :8080
```

Stop the conflicting service or recreate stack with different ports.

### Permission Denied

If permission errors happen with volumes:

```bash
cd lamp-stack
sudo chown -R $USER:$USER .
```

### Docker Compose Not Found

Install Docker Compose:

```bash
# Ubuntu/Debian
sudo apt install docker-compose

# macOS
brew install docker-compose
```

## Project Structure

```
autostack/
├── cmd/                 # CLI commands
├── internal/stack/      # Stack implementations
│   ├── lamp.go
│   ├── mariadb.go
│   ├── observability.go
│   └── templates.go
├── main.go
lamp-stack/
├── docker-compose.yml
├── www/
│   └── index.php
├── mysql/
├── logs/
└── README.md
mariadb-stack/
├── docker-compose.yml
├── mariadb/
└── README.md
observability-stack/
├── docker-compose.yml
├── prometheus/
│   └── prometheus.yml
├── grafana/
│   └── provisioning/
└── README.md
```

## Contributing

1. Fork the repo
2. Create a new feature branch
3. Add or modify stack definitions in `internal/stack/`
4. Update `ListStacks()` function
5. Submit a pull request

## License

MIT License — see LICENSE file for details.

## Author

Created for rapid development environment provisioning.

GitHub: [https://github.com/bait-py/autostack](https://github.com/bait-py/autostack)

Issues: [https://github.com/bait-py/autostack/issues](https://github.com/bait-py/autostack/issues)