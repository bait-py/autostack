package stack

import (
	"errors"
	"fmt"
)

// Create creates a stack based on the specified name
func Create(name string) error {
	switch name {
	case "lamp":
		return createLamp()
	case "observability", "obs":
		return createObservability()
	case "mariadb":
		return createMariaDB()
	default:
		return errors.New("stack not recognized: " + name)
	}
}

// ListStacks shows all available stacks
func ListStacks() {
	stacks := []struct {
		Name        string
		Alias       string
		Description string
	}{
		{"lamp", "", "LAMP stack with Apache, MySQL, PHP and phpMyAdmin"},
		{"observability", "obs", "Prometheus + Grafana + Node Exporter for monitoring"},
		{"mariadb", "", "MariaDB + phpMyAdmin for databases"},
	}

	fmt.Println("\nAvailable stacks:")
	for _, s := range stacks {
		if s.Alias != "" {
			fmt.Printf("  %s (%s) - %s\n", s.Name, s.Alias, s.Description)
		} else {
			fmt.Printf("  %s - %s\n", s.Name, s.Description)
		}
	}
	fmt.Println()
}
