package stack

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StackEnvVars defines configurable environment variables for a stack
type StackEnvVars struct {
	VarName     string
	Description string
	Default     string
	Value       string
}

// StackPort defines a configurable port for a service
type StackPort struct {
	ServiceName string
	Description string
	Default     string
	HostPort    string // The port on the host machine
	Internal    string // The internal container port (usually fixed)
}

// PromptAutoStart asks the user if they want to auto-start the stack
func PromptAutoStart() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nAuto-start the stack after creation? (Y/n): ")

	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	// Default is "yes"
	if response == "" || response == "y" || response == "yes" {
		return true
	}
	return false
}

// PromptEnvVars prompts the user for environment variable values
func PromptEnvVars(vars []StackEnvVars) map[string]string {
	if len(vars) == 0 {
		return nil
	}

	fmt.Println("\n=== Environment Variables Configuration ===")
	fmt.Println("Press Enter to use default values")
	fmt.Println(strings.Repeat("-", 60))

	reader := bufio.NewReader(os.Stdin)
	result := make(map[string]string)

	for i, v := range vars {
		fmt.Printf("\n[%d/%d] %s\n", i+1, len(vars), v.Description)
		fmt.Printf("      Default: %s\n", v.Default)
		fmt.Printf("      Enter value: ")

		value, _ := reader.ReadString('\n')
		value = strings.TrimSpace(value)

		if value == "" {
			result[v.VarName] = v.Default
		} else {
			result[v.VarName] = value
		}
	}

	fmt.Println(strings.Repeat("-", 60))
	return result
}

// PromptPorts prompts the user for port configurations
func PromptPorts(ports []StackPort) map[string]string {
	if len(ports) == 0 {
		return nil
	}

	fmt.Println("\n=== Port Configuration ===")
	fmt.Println("Press Enter to use default values")
	fmt.Println(strings.Repeat("-", 60))

	reader := bufio.NewReader(os.Stdin)
	result := make(map[string]string)

	for i, p := range ports {
		fmt.Printf("\n[%d/%d] %s\n", i+1, len(ports), p.Description)
		fmt.Printf("      Default: %s\n", p.Default)
		fmt.Printf("      Enter port: ")

		value, _ := reader.ReadString('\n')
		value = strings.TrimSpace(value)

		if value == "" {
			result[p.ServiceName] = p.Default
		} else {
			result[p.ServiceName] = value
		}
	}

	fmt.Println(strings.Repeat("-", 60))
	return result
}

// ConfirmConfiguration shows the configuration and asks for confirmation
func ConfirmConfiguration(envVars map[string]string, ports map[string]string) bool {
	hasConfig := len(envVars) > 0 || len(ports) > 0
	if !hasConfig {
		return true
	}

	fmt.Println("\n=== Configuration Summary ===")

	if len(envVars) > 0 {
		fmt.Println("\nEnvironment Variables:")
		for key, value := range envVars {
			// Partially hide passwords
			displayValue := value
			if strings.Contains(strings.ToLower(key), "password") ||
				strings.Contains(strings.ToLower(key), "secret") {
				if len(value) > 4 {
					displayValue = value[:2] + "****" + value[len(value)-2:]
				} else {
					displayValue = "****"
				}
			}
			fmt.Printf("      %s: %s\n", key, displayValue)
		}
	}

	if len(ports) > 0 {
		fmt.Println("\nPorts:")
		for service, port := range ports {
			fmt.Printf("  %s: %s\n", service, port)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nConfirm configuration? (Y/n): ")

	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	return response == "" || response == "y" || response == "yes"
}
