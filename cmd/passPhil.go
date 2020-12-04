/*
Copyright © 2020 Henry Mortimer <henry@morti.net>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hennersz/advent-of-code/2020/day2"
	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command.
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Solves day 2: password philosophy",
	Run: func(cmd *cobra.Command, args []string) {
		reader := getInput("inFile", cmd)
		defer reader.Close()

		var res int
		var err error

		if b, _ := cmd.Flags().GetBool("position"); b {
			res, err = day2.SolvePosition(reader)
		} else {
			res, err = day2.SolveRange(reader)
		}

		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Result: %d\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)
	day2Cmd.PersistentFlags().BoolP("position", "p", false, "Solve for position")
}
