package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jashandeep31/gossh/server"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gossh",
	Short: "Gossh is a custom SSH manager CLI tool",
	Long: `Gossh helps you manage SSH connections by adding, listing,
connecting to, and removing saved SSH server entries with ease.`,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the servers",
	Run: func(cmd *cobra.Command, args []string) {
		servers, err := server.LoadServer()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading servers: %v\n", err)
			return
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Index", "Name", "Host", "User"})
		for i, s := range servers {
			table.Append([]string{strconv.Itoa(i), s.Name, s.Host, s.User})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(delCmd)
	rootCmd.AddCommand(connectCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
