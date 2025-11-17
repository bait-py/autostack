# AutoStack

Fast and efficient Docker stack generator for development environments.

## Overview

AutoStack generates pre-configured Docker Compose stacks with a single command. Each stack includes all necessary configuration files, documentation, and is ready to deploy immediately.

## Installation

```bash
# Build from source
go build -o autostack

# (Optional) Install globally
sudo mv autostack /usr/local/bin/
```

## Quick Start

```bash
# List available stacks
./autostack list

# Create a stack
./autostack create lamp

# Start the stack
cd lamp-stack
docker-compose up -d
```

## Available Stacks

| Stack | Services | Default Ports | Use Case |
|-------|----------|---------------|----------|
| `lamp` | Apache + PHP 8.2<br>MySQL 8.0<br>phpMyAdmin | 8080 (Web)<br>3306 (MySQL)<br>8081 (phpMyAdmin) | Web development with PHP and MySQL |
| `mariadb` | MariaDB<br>phpMyAdmin | 3306 (MariaDB)<br>8080 (phpMyAdmin) | Database development with MariaDB |
| `observability` | Prometheus<br>Grafana<br>Node Exporter | 9090 (Prometheus)<br>3000 (Grafana)<br>9100 (Node Exporter) | Monitoring and metrics visualization |

## Configuration

During stack creation, you can configure:

### Environment Variables
- Database credentials
- Application settings
- Security configurations

All variables have sensible defaults. Press Enter to use default values or input custom values.

### Ports
- Customize exposed ports for each service
- Avoid conflicts with existing services
- Default ports are pre-configured for immediate use

### Auto-start
- Option to automatically start services after creation
- Uses `docker-compose up -d` command
- Can be manually started later if declined

## Usage Examples

### Basic Stack Creation
```bash
./autostack create lamp
# Follow prompts, press Enter for defaults
cd lamp-stack
docker-compose up -d
```

### Custom Configuration
```bash
./autostack create mariadb
# Enter custom values when prompted:
# - Database name: myproject_db
# - User: developer
# - Password: secure_password
# - Ports: use defaults (press Enter)
cd mariadb-stack
docker-compose up -d
```

### Using Aliases
```bash
./autostack create obs  # Short for 'observability'
```

## Stack Details

### LAMP Stack
Includes:
- Apache web server with PHP 8.2
- MySQL 8.0 database
- phpMyAdmin for database management
- Pre-configured PHP connection example
- Apache logs directory
- Persistent MySQL data

Generated files:
- `docker-compose.yml` - Service configuration
- `www/index.php` - PHP example with MySQL connection
- `README.md` - Stack documentation
- `.gitignore` - Git ignore rules

### MariaDB Stack
Includes:
- MariaDB latest version
- phpMyAdmin for database management
- Persistent data storage
- Pre-configured connection settings

Generated files:
- `docker-compose.yml` - Service configuration
- `README.md` - Stack documentation
- `.gitignore` - Git ignore rules

### Observability Stack
Includes:
- Prometheus for metrics collection
- Grafana for visualization
- Node Exporter for system metrics
- Pre-configured Prometheus datasource in Grafana
- Sample Prometheus configuration

Generated files:
- `docker-compose.yml` - Service configuration
- `prometheus/prometheus.yml` - Prometheus config
- `grafana/provisioning/datasources/prometheus.yml` - Grafana datasource
- `README.md` - Stack documentation
- `.gitignore` - Git ignore rules

## Default Credentials

### LAMP Stack
- MySQL Root Password: `rootpassword`
- Database: `lamp_db`
- User: `lamp_user`
- Password: `lamp_password`

### MariaDB Stack
- Root Password: `rootpassword`
- Database: `mydb`
- User: `myuser`
- Password: `mypassword`

### Observability Stack
- Grafana Admin: `admin`
- Grafana Password: `admin`

## Commands Reference

| Command | Description |
|---------|-------------|
| `./autostack list` | Display all available stacks |
| `./autostack create <stack>` | Create a new stack |
| `./autostack --help` | Show help information |

## Stack Management

### Start Services
```bash
cd <stack-directory>
docker-compose up -d
```

### Stop Services
```bash
docker-compose down
```

### View Logs
```bash
docker-compose logs -f
```

### Remove Everything (including data)
```bash
docker-compose down -v
```

## Requirements

- Docker
- Docker Compose
- Go 1.22+ (for building from source)

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
└── README.md
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add or modify stack definitions
4. Submit a pull request

## License

MIT License
