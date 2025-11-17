package stack

// DockerComposeMariaDB contains the template for MariaDB + phpMyAdmin
const DockerComposeMariaDB = `version: '3.8'

services:
  # MariaDB Database
  mariadb:
    image: mariadb:latest
    container_name: mariadb_db
    ports:
      - "{{PORT_MARIADB}}:3306"
    volumes:
      - ./mariadb:/var/lib/mysql
    environment:
      MYSQL_ROOT_PASSWORD: {{MYSQL_ROOT_PASSWORD}}
      MYSQL_DATABASE: {{MYSQL_DATABASE}}
      MYSQL_USER: {{MYSQL_USER}}
      MYSQL_PASSWORD: {{MYSQL_PASSWORD}}
    networks:
      - mariadb-network
    restart: unless-stopped

  # phpMyAdmin for database management
  phpmyadmin:
    image: phpmyadmin:latest
    container_name: mariadb_phpmyadmin
    ports:
      - "{{PORT_PHPMYADMIN}}:80"
    environment:
      PMA_HOST: mariadb
      PMA_PORT: 3306
      PMA_USER: root
      PMA_PASSWORD: {{MYSQL_ROOT_PASSWORD}}
    depends_on:
      - mariadb
    networks:
      - mariadb-network
    restart: unless-stopped

networks:
  mariadb-network:
    driver: bridge
`

// ReadmeMariaDB contains the stack documentation
const ReadmeMariaDB = `# MariaDB + phpMyAdmin Stack

## Project structure

- **mariadb/**: Persistent MariaDB data

## Included services

- **MariaDB**: Port {{PORT_MARIADB}}
- **phpMyAdmin**: Port {{PORT_PHPMYADMIN}}

## Configuration

### MariaDB
- Host: mariadb (inside Docker) or localhost:{{PORT_MARIADB}} (from your machine)
- Database: {{MYSQL_DATABASE}}
- User: {{MYSQL_USER}}
- Password: {{MYSQL_PASSWORD}}
- Root Password: {{MYSQL_ROOT_PASSWORD}}

### phpMyAdmin
- URL: http://localhost:{{PORT_PHPMYADMIN}}
- User: root
- Password: {{MYSQL_ROOT_PASSWORD}}

## Useful commands

### Start the stack
` + "```bash" + `
docker-compose up -d
` + "```" + `

### Stop the stack
` + "```bash" + `
docker-compose down
` + "```" + `

### View logs
` + "```bash" + `
docker-compose logs -f
` + "```" + `

### Access MariaDB container
` + "```bash" + `
docker exec -it mariadb_db bash
` + "```" + `

### Connect to MariaDB from command line
` + "```bash" + `
docker exec -it mariadb_db mysql -u root -p{{MYSQL_ROOT_PASSWORD}}
` + "```" + `

## Access URLs

- phpMyAdmin: http://localhost:{{PORT_PHPMYADMIN}}
- MariaDB: localhost:{{PORT_MARIADB}}

## Database backup

` + "```bash" + `
docker exec mariadb_db mysqldump -u root -p{{MYSQL_ROOT_PASSWORD}} {{MYSQL_DATABASE}} > backup.sql
` + "```" + `

## Restore backup

` + "```bash" + `
docker exec -i mariadb_db mysql -u root -p{{MYSQL_ROOT_PASSWORD}} {{MYSQL_DATABASE}} < backup.sql
` + "```" + `

## Notes

- MariaDB data persists in the mariadb/ directory
- To change credentials, edit environment variables in docker-compose.yml
- Compatible with standard MySQL clients
`

// GitignoreMariaDB contains files to ignore
const GitignoreMariaDB = `mariadb/
*.sql
*.log
`

// createMariaDB creates a stack with MariaDB and phpMyAdmin
func createMariaDB() error {
	// Define configurable environment variables
	envVars := []StackEnvVars{
		{
			VarName:     "MYSQL_ROOT_PASSWORD",
			Description: "MariaDB root password",
			Default:     "rootpassword",
		},
		{
			VarName:     "MYSQL_DATABASE",
			Description: "MariaDB database name",
			Default:     "mydb",
		},
		{
			VarName:     "MYSQL_USER",
			Description: "MariaDB user",
			Default:     "myuser",
		},
		{
			VarName:     "MYSQL_PASSWORD",
			Description: "MariaDB user password",
			Default:     "mypassword",
		},
	}

	// Define configurable ports
	ports := []StackPort{
		{
			ServiceName: "mariadb",
			Description: "MariaDB database port",
			Default:     "3306",
			Internal:    "3306",
		},
		{
			ServiceName: "phpmyadmin",
			Description: "phpMyAdmin web interface port",
			Default:     "8080",
			Internal:    "80",
		},
	}

	// Prompt for environment variables
	envValues := PromptEnvVars(envVars)

	// Prompt for ports
	portValues := PromptPorts(ports)

	// Confirm configuration
	if !ConfirmConfiguration(envValues, portValues) {
		return nil
	}

	// Prompt for auto-start
	autoStart := PromptAutoStart()

	config := StackConfig{
		Name:        "MariaDB",
		Description: "MariaDB database with phpMyAdmin for web management",
		ProjectDir:  "mariadb-stack",
		AutoStart:   autoStart,
		Ports: map[string]string{
			"phpMyAdmin": portValues["phpmyadmin"],
			"MariaDB":    portValues["mariadb"],
		},
		Dirs: []string{
			"mariadb",
		},
		Files: map[string]string{
			"docker-compose.yml": DockerComposeMariaDB,
			"README.md":          ReadmeMariaDB,
			".gitignore":         GitignoreMariaDB,
		},
		EnvVars:        envVars,
		ConfigurePorts: ports,
	}

	// Apply environment variables and ports to templates
	config.ApplyEnvVars(envValues)
	config.ApplyPorts(portValues)

	return GenerateStack(config)
}
