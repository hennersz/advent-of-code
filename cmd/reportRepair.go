/*
Copyright © 2020 Henry Mortimer <henry@morti.net>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hennersz/advent-of-code/2020/day1"
	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command.
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Solves day 1: Report Repair",
	Run: func(cmd *cobra.Command, args []string) {
		reader := getInput("inFile", cmd)
		defer reader.Close()

		var res int
		var err error

		if b, _ := cmd.Flags().GetBool("triple"); b {
			res, err = day1.SolveTriple(reader)
		} else {
			res, err = day1.Solve(reader)
		}

		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Result: %d\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
	day1Cmd.PersistentFlags().BoolP("triple", "t", false, "Solve for triple")
}
