package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jashandeep31/gossh/server"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del [index]",
	Short: "Delete the server by idnex",
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid index: %v\n", err)
			return
		}
		err = server.DeleteServer(index) // Call your DeleteServer function
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting server: %v\n", err)
			return
		}
	},
}
