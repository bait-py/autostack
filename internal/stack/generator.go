package stack

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// StackConfig defines the configuration to generate a stack
type StackConfig struct {
	Name           string
	ProjectDir     string
	Files          map[string]string // relative path -> content
	Dirs           []string          // directories to create
	AutoStart      bool              // run docker-compose up -d automatically
	Ports          map[string]string // service -> port (to display in summary)
	Description    string            // stack description
	EnvVars        []StackEnvVars    // configurable environment variables
	ConfigurePorts []StackPort       // configurable ports
}

// ApplyEnvVars replaces environment variable placeholders in files
func (config *StackConfig) ApplyEnvVars(values map[string]string) {
	if len(values) == 0 {
		return
	}

	// Replace placeholders in all files
	for path, content := range config.Files {
		for key, value := range values {
			placeholder := "{{" + key + "}}"
			content = strings.ReplaceAll(content, placeholder, value)
		}
		config.Files[path] = content
	}
}

// ApplyPorts replaces port placeholders in files
func (config *StackConfig) ApplyPorts(values map[string]string) {
	if len(values) == 0 {
		return
	}

	// Replace port placeholders and update display ports
	for path, content := range config.Files {
		for serviceName, hostPort := range values {
			placeholder := "{{PORT_" + strings.ToUpper(serviceName) + "}}"
			content = strings.ReplaceAll(content, placeholder, hostPort)
		}
		config.Files[path] = content
	}

	// Update ports map for display
	for serviceName, hostPort := range values {
		config.Ports[serviceName] = hostPort
	}
}

// GenerateStack creates all necessary files and directories for a stack
func GenerateStack(config StackConfig) error {
	fmt.Printf("Generating files for %s stack...\n", config.Name)

	// Create main project directory
	if err := os.MkdirAll(config.ProjectDir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	// Create subdirectories
	for _, dir := range config.Dirs {
		fullPath := filepath.Join(config.ProjectDir, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %w", dir, err)
		}
	}

	// Create files
	var fileNames []string
	for relPath, content := range config.Files {
		fullPath := filepath.Join(config.ProjectDir, relPath)
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("error writing %s: %w", relPath, err)
		}
		fileNames = append(fileNames, relPath)
	}

	// Run docker-compose up -d if enabled
	if config.AutoStart {
		if err := startDockerCompose(config.ProjectDir); err != nil {
			fmt.Printf("\nWARNING: Error starting Docker Compose: %v\n", err)
			fmt.Println("You can start it manually with: docker-compose up -d")
		}
	}

	// Show summary
	printSuccess(config, fileNames)

	return nil
}

// startDockerCompose runs docker-compose up -d in the specified directory
func startDockerCompose(projectDir string) error {
	fmt.Println("\nStarting Docker services...")

	return exec.Command("sh", "-c", fmt.Sprintf("cd %s && docker-compose up -d", projectDir)).Run()
}

// printSuccess shows a formatted success message
func printSuccess(config StackConfig, files []string) {
	fmt.Println("\n=== Stack Created Successfully ===")
	fmt.Printf("Name: %s\n", config.Name)

	if config.Description != "" {
		fmt.Printf("Description: %s\n", config.Description)
	}

	fmt.Printf("Directory: %s\n", config.ProjectDir)
	fmt.Println("\nGenerated files:")
	for _, file := range files {
		fmt.Printf("  - %s\n", file)
	}

	if !config.AutoStart {
		fmt.Println("\nTo start the stack:")
		fmt.Printf("  cd %s\n", config.ProjectDir)
		fmt.Println("  docker-compose up -d")
	} else {
		fmt.Println("\nStack is starting...")
	}

	if len(config.Ports) > 0 {
		fmt.Println("\nAccess URLs:")
		for service, port := range config.Ports {
			fmt.Printf("  %s: http://localhost:%s\n", service, port)
		}
	}
	fmt.Println()
}
