/*
Copyright © 2020 Henry Mortimer <henry@morti.net>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hennersz/advent-of-code/2020/passPhil"
	"github.com/spf13/cobra"
)

// passPhilCmd represents the passPhil command
var passPhilCmd = &cobra.Command{
	Use:   "pp",
	Short: "Solves day 2: password philosophy",
	Run: func(cmd *cobra.Command, args []string) {
		reader := getInput("inFile", cmd)
		defer reader.Close()

		var res int
		var err error

		if b, _ := cmd.Flags().GetBool("position"); b {
			res, err = passPhil.SolvePosition(reader)
		} else {
			res, err = passPhil.SolveRange(reader)
		}

		if err != nil {
			fmt.Printf("Error ocurred: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Result: %d\n", res)
	},
}

func init() {
	rootCmd.AddCommand(passPhilCmd)
	passPhilCmd.PersistentFlags().BoolP("position", "p", false, "Solve for position")
}
