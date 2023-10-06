/*
Copyright © 2023 annonch
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// benchCmd represents the bench command
var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		const totalRuns = 100
		var elapsedRuns = 0
		fmt.Println("bench called")
		fmt.Printf("hello %v : %v world\n", totalRuns, elapsedRuns)

		var priority float32
		var events []string

		events = append(events, "chris")
		priority = 123.3
		fmt.Println(priority)

		//for _, event := range events {
		for i := 0; i < 10; i++ {
			events = append(events, "Chris")
		}

		fmt.Printf("events: %v, type: %T, length %v\n", events, events, len(events))
	},
}

func init() {
	rootCmd.AddCommand(benchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// benchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	benchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
