/*
Copyright © 2020 Henry Mortimer <henry@morti.net>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hennersz/advent-of-code/2020/day3"
	"github.com/spf13/cobra"
)

// day3Cmd represents the day3 command.
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Solves day 3",
	Run: func(cmd *cobra.Command, args []string) {
		input := getInput("inFile", cmd)
		defer input.Close()
		slopes := getInput("slopes", cmd)
		defer slopes.Close()

		res, err := day3.Solve(input, slopes)
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Result: %d\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
	day3Cmd.PersistentFlags().StringP("slopes", "s", "", "Slopes file path (required)")
	day3Cmd.MarkPersistentFlagRequired("slopes")
}
