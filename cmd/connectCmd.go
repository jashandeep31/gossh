package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jashandeep31/gossh/server"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect [index]",
	Short: "Connect to the server using index id",
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid index: %v\n", err)
			return
		}
		err = server.Connect(index)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error connecting server: %v\n", err)
			return
		}
	},
}
