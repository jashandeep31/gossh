package cmd

import (
	"fmt"

	"github.com/jashandeep31/gossh/server"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new server to the data",
	Run: func(cmd *cobra.Command, args []string) {
		var name, host, user string

		// Prompt for the server name
		fmt.Print("Enter server name {example: main server}: ")
		_, err := fmt.Scanf("%s", &name)
		if err != nil {
			fmt.Println("Error reading server name:", err)
			return
		}

		// Prompt for server Host
		fmt.Print("Enter server host: ")
		_, err = fmt.Scanf("%s", &host)
		if err != nil {
			fmt.Println("Error reading server host:", err)
			return
		}

		// Prompt for server User
		fmt.Print("Enter server user: ")
		_, err = fmt.Scanf("%s", &user)
		if err != nil {
			fmt.Println("Error reading server user:", err)
			return
		}

		newServer := server.Server{
			Name: name,
			Host: host,
			User: user,
		}
		server.AddServer(newServer)

	},
}
