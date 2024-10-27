package server

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var servers = []Server{}

func init() {
	LoadServer()
}
func GetFile() (*os.File, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not get home directory: %v", err)
	}

	gosshDir := filepath.Join(homeDir, ".gossh")
	if err := os.MkdirAll(gosshDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("could not create directory: %v", err)
	}

	filePath := filepath.Join(gosshDir, JSONFILE)

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create the file with an empty array if it doesn't exist
			file, err = os.Create(filePath)
			if err != nil {
				return nil, fmt.Errorf("could not create file: %v", err)
			}
			// Write an empty JSON array to the file
			if _, err := file.WriteString("[]"); err != nil {
				return nil, fmt.Errorf("could not write empty array to file: %v", err)
			}
			file.Seek(0, 0) // Reset the file pointer to the beginning
		} else {
			return nil, fmt.Errorf("could not open file: %v", err)
		}
	}
	return file, nil
}

func LoadServer() ([]Server, error) {

	file, err := GetFile()
	if err != nil {
		return nil, fmt.Errorf("file isn't working")
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&servers); err != nil {
		return nil, fmt.Errorf("could not decode JSON: %v", err)
	}
	return servers, nil
}

func AddServer(server Server) error {
	servers = append(servers, server)
	err := UpdateJsonData()
	if err != nil {
		return fmt.Errorf("failed to update")
	}
	return nil
}

func DeleteServer(id int) error {
	if id < 0 || id >= len(servers) {
		return fmt.Errorf("invalid server ID: %d", id)
	}
	servers = append(servers[:id], servers[id+1:]...)

	err := UpdateJsonData()
	if err != nil {
		return fmt.Errorf("failed to del")
	}
	return nil
}

func UpdateJsonData() error {
	file, err := os.OpenFile(JSONFILE, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("could not open servers.json for writing: %v", err)
	}
	defer file.Close() // Ensure the file is closed after writing

	// Encode the updated servers slice back to JSON
	if err := json.NewEncoder(file).Encode(servers); err != nil {
		return fmt.Errorf("could not write to servers.json: %v", err)
	}
	return nil
}

func Connect(id int) error {
	if id < 0 || id >= len(servers) {
		return fmt.Errorf("invalid server ID: %d", id)
	}
	server := servers[id]
	cmdStr := fmt.Sprintf("ssh %s@%s", server.User, server.Host)

	cmd := exec.Command("bash", "-c", cmdStr)

	// Set the command's output to the terminal's stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Important: set the input to allow interaction

	// Start the command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}

	return nil
}
