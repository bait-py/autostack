package stack

// DockerComposeLAMP contains the template for LAMP stack
const DockerComposeLAMP = `version: '3.8'

services:
  # Apache Web Server with PHP
  web:
    image: php:8.2-apache
    container_name: lamp_web
    ports:
      - "{{PORT_WEB}}:80"
    volumes:
      - ./www:/var/www/html
      - ./logs:/var/log/apache2
    depends_on:
      - db
    networks:
      - lamp-network
    environment:
      - APACHE_DOCUMENT_ROOT=/var/www/html

  # MySQL Database
  db:
    image: mysql:8.0
    container_name: lamp_db
    ports:
      - "{{PORT_MYSQL}}:3306"
    volumes:
      - ./mysql:/var/lib/mysql
    environment:
      MYSQL_ROOT_PASSWORD: {{MYSQL_ROOT_PASSWORD}}
      MYSQL_DATABASE: {{MYSQL_DATABASE}}
      MYSQL_USER: {{MYSQL_USER}}
      MYSQL_PASSWORD: {{MYSQL_PASSWORD}}
    networks:
      - lamp-network

  # phpMyAdmin for database management
  phpmyadmin:
    image: phpmyadmin:latest
    container_name: lamp_phpmyadmin
    ports:
      - "{{PORT_PHPMYADMIN}}:80"
    environment:
      PMA_HOST: db
      PMA_PORT: 3306
      PMA_USER: root
      PMA_PASSWORD: {{MYSQL_ROOT_PASSWORD}}
    depends_on:
      - db
    networks:
      - lamp-network

networks:
  lamp-network:
    driver: bridge
`

// IndexPHP contains a sample PHP file with MySQL connection
const IndexPHP = `<?php
phpinfo();

// MySQL connection test
$host = 'db';
$db   = '{{MYSQL_DATABASE}}';
$user = '{{MYSQL_USER}}';
$pass = '{{MYSQL_PASSWORD}}';

try {
    $pdo = new PDO("mysql:host=$host;dbname=$db", $user, $pass);
    echo "<h2>Successful connection to MySQL!</h2>";
} catch (PDOException $e) {
    echo "<h2>Connection error: " . $e->getMessage() . "</h2>";
}
?>
`

// ReadmeLAMP contains the stack documentation
const ReadmeLAMP = `# LAMP Stack with Docker

## Project structure

- **www/**: Directory for your PHP/HTML files
- **mysql/**: Persistent MySQL data
- **logs/**: Apache logs

## Included services

- **Apache + PHP 8.2**: Port {{PORT_WEB}}
- **MySQL 8.0**: Port {{PORT_MYSQL}}
- **phpMyAdmin**: Port {{PORT_PHPMYADMIN}}

## Configuration

### MySQL
- Host: db (inside Docker) or localhost:{{PORT_MYSQL}} (from your machine)
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

### Access web container
` + "```bash" + `
docker exec -it lamp_web bash
` + "```" + `

## Access URLs

- Web application: http://localhost:{{PORT_WEB}}
- phpMyAdmin: http://localhost:{{PORT_PHPMYADMIN}}

## Notes

- Files in www/ are automatically synced with the container
- MySQL data persists in the mysql/ directory
- To change credentials, edit environment variables in docker-compose.yml
`

// GitignoreLAMP contains files to ignore in git
const GitignoreLAMP = `mysql/
logs/
*.log
`
