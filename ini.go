package goini

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Contain the object for a parsed ini file.
type Content struct {
	path     []string
	keys     map[string]string
	sections map[string]map[string]string
}

// Parse the file content keys.
func (c *Content) parseKeys(data []string) error {
	// Reset the keys map.
	c.keys = make(map[string]string)
	c.sections = make(map[string]map[string]string)

	// Create a new rw mutex.
	var mutex sync.RWMutex

	for i, row := range data {
		// Check if the entry is a comment, an empty line or a linebreak.
		if strings.HasPrefix(row, "#") || strings.HasPrefix(row, ";") || row == "" || row == "\r" {
			continue
		}

		// Check if the row is a section header.
		if strings.Contains(row, "[") && strings.Contains(row, "]") {
			// Create the new section.
			return c.newSection(strings.Split(strings.Split(row, "[")[1], "]")[0], data[i+1:]...)
		}

		// Split the row by the key and value.
		content := strings.Split(row, "=")

		// Check if the content matches the required specification.
		if len(content) != 2 {
			return fmt.Errorf("failed to parse row at %s", row)
		}

		// Get the key and value from the rows.
		var (
			key, value string = content[0], content[1]
		)

		// Check if the key has already been defined.
		if _, ok := c.keys[key]; ok {
			return fmt.Errorf("%s used multiple times", key)
		}

		mutex.Lock()
		// Set the key and value inside of the map.
		c.keys[key] = strings.TrimSuffix(value, "\r")
		mutex.Unlock()
	}

	return nil
}

// Create a new section inside of the map.
func (c *Content) newSection(name string, data ...string) error {
	// Create a new rw mutex.
	var mutex sync.RWMutex

	// Check if the section already exists.
	if _, ok := c.sections[name]; ok {
		return fmt.Errorf("section redefined at %s", name)
	}

	mutex.Lock()

	// Create a new section.
	c.sections[name] = make(map[string]string)

	mutex.Unlock()

	for i, row := range data {
		// Check if the entry is a comment, an empty line or a linebreak.
		if strings.HasPrefix(row, "#") || strings.HasPrefix(row, ";") || row == "" || row == "\r" {
			continue
		}

		// Check if the row is a section header.
		if strings.Contains(row, "[") && strings.Contains(row, "]") {
			// Create the new section.
			return c.newSection(strings.Split(strings.Split(row, "[")[1], "]")[0], data[i+1:]...)
		}

		// Split the row by the key and value.
		content := strings.Split(row, "=")

		// Check if the content matches the required specification.
		if len(content) != 2 {
			return fmt.Errorf("failed to parse row at %s", row)
		}

		// Get the key and value from the rows.
		var (
			key, value string = content[0], content[1]
		)

		// Check if the key has already been defined.
		if _, ok := c.keys[key]; ok {
			return fmt.Errorf("%s used multiple times", key)
		}

		mutex.Lock()

		// Set the key and value inside of the map.
		c.sections[name][key] = strings.TrimSuffix(value, "\r")

		mutex.Unlock()
	}

	return nil
}

// Reload the keys inside of the ini configuration.
func (c *Content) ReloadKeys() error {
	// Read the file path.
	bytes, err := os.ReadFile(filepath.Join(c.path...))
	if err != nil {
		return err
	}

	// Parse the file keys.
	return c.parseKeys(strings.Split(string(bytes), "\n"))
}

// Fetch the value from the map using the key name.
func (c *Content) GetValueFromKey(name string) *string {
	// Fetch the value from the map.
	value, exists := c.keys[name]
	if !exists {
		return nil
	}

	return &value
}

// Fetch the value from the map using the key name.
func (c *Content) GetValueFromSectionKey(section, name string) *string {
	// Fetch the value from the map.
	value, exists := c.sections[section][name]
	if !exists {
		return nil
	}

	return &value
}

// Create a new ini parser object.
func New(path ...string) (*Content, error) {
	// Create a new ini parser object.
	var content *Content = &Content{path: path, keys: make(map[string]string)}

	// Read the file path.
	bytes, err := os.ReadFile(filepath.Join(path...))
	if err != nil {
		return nil, err
	}

	// Parse the file keys.
	if err := content.parseKeys(strings.Split(string(bytes), "\n")); err != nil {
		return nil, err
	}

	return content, nil
}
