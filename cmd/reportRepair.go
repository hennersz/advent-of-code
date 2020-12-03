/*
Copyright © 2020 Henry Mortimer <henry@morti.net>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hennersz/advent-of-code/2020/reportRepair"
	"github.com/spf13/cobra"
)

// reportRepairCmd represents the reportRepair command
var reportRepairCmd = &cobra.Command{
	Use:   "rr",
	Short: "Solves day 1: Report Repair",
	Run: func(cmd *cobra.Command, args []string) {
		reader := getInput("inFile", cmd)
		defer reader.Close()

		var res int
		var err error

		if b, _ := cmd.Flags().GetBool("triple"); b {
			res, err = reportRepair.SolveTriple(reader)
		} else {
			res, err = reportRepair.Solve(reader)
		}

		if err != nil {
			fmt.Printf("Error ocurred: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Result: %d\n", res)
	},
}

func init() {
	rootCmd.AddCommand(reportRepairCmd)
	reportRepairCmd.PersistentFlags().BoolP("triple", "t", false, "Solve for triple")
}
