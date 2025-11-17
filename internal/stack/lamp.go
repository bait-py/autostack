package stack

// createLamp creates a LAMP stack (Linux, Apache, MySQL, PHP)
func createLamp() error {
	// Define configurable environment variables
	envVars := []StackEnvVars{
		{
			VarName:     "MYSQL_ROOT_PASSWORD",
			Description: "MySQL root password",
			Default:     "rootpassword",
		},
		{
			VarName:     "MYSQL_DATABASE",
			Description: "MySQL database name",
			Default:     "lamp_db",
		},
		{
			VarName:     "MYSQL_USER",
			Description: "MySQL user",
			Default:     "lamp_user",
		},
		{
			VarName:     "MYSQL_PASSWORD",
			Description: "MySQL user password",
			Default:     "lamp_password",
		},
	}

	// Define configurable ports
	ports := []StackPort{
		{
			ServiceName: "web",
			Description: "Apache web server port",
			Default:     "8080",
			Internal:    "80",
		},
		{
			ServiceName: "mysql",
			Description: "MySQL database port",
			Default:     "3306",
			Internal:    "3306",
		},
		{
			ServiceName: "phpmyadmin",
			Description: "phpMyAdmin web interface port",
			Default:     "8081",
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
		Name:        "LAMP",
		Description: "LAMP stack with Apache, MySQL 8.0, PHP 8.2 and phpMyAdmin",
		ProjectDir:  "lamp-stack",
		AutoStart:   autoStart,
		Ports: map[string]string{
			"Web Application": portValues["web"],
			"phpMyAdmin":      portValues["phpmyadmin"],
			"MySQL":           portValues["mysql"],
		},
		Dirs: []string{
			"www",
			"mysql",
			"logs",
		},
		Files: map[string]string{
			"docker-compose.yml": DockerComposeLAMP,
			"www/index.php":      IndexPHP,
			"README.md":          ReadmeLAMP,
			".gitignore":         GitignoreLAMP,
		},
		EnvVars:        envVars,
		ConfigurePorts: ports,
	}

	// Apply environment variables and ports to templates
	config.ApplyEnvVars(envValues)
	config.ApplyPorts(portValues)

	return GenerateStack(config)
}
