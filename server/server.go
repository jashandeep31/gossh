// server/server.go
package server

import "fmt"

const JSONFILE = "servers.json"

type Server struct {
	Name string `json:"name"`
	Host string `json:"host"`
	User string `json:"user"`
}

func (s Server) Connect() {
	fmt.Printf("Connecting to %s@%s...\n", s.User, s.Host)
	// Implement SSH connection logic
}
