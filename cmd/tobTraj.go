/*
Copyright © 2020 Henry Mortimer <henry@morti.net>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hennersz/advent-of-code/2020/tobTraj"
	"github.com/spf13/cobra"
)

// tobTrajCmd represents the tobTraj command.
var tobTrajCmd = &cobra.Command{
	Use:   "tt",
	Short: "Solves day 3",
	Run: func(cmd *cobra.Command, args []string) {
		input := getInput("inFile", cmd)
		defer input.Close()
		slopes := getInput("slopes", cmd)
		defer slopes.Close()

		res, err := tobTraj.Solve(input, slopes)
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Result: %d\n", res)
	},
}

func init() {
	rootCmd.AddCommand(tobTrajCmd)
	tobTrajCmd.PersistentFlags().StringP("slopes", "s", "", "Slopes file path (required)")
	tobTrajCmd.MarkPersistentFlagRequired("slopes")
}
