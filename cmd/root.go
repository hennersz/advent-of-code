package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "advent-of-code",
	Short: "advent of code runner",
	Long:  `A runner for advent of code solutions`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("inFile", "i", "", "Input file path (required)")
	rootCmd.MarkPersistentFlagRequired("inFile")
}

func getInput(flag string, cmd *cobra.Command) *os.File {
	inFilePath, err := cmd.Flags().GetString(flag)
	if err != nil {
		fmt.Printf("Error ocurred: %v\n", err)
		os.Exit(1)
	}

	reader, err := os.Open(inFilePath)
	if err != nil {
		fmt.Printf("Error ocurred: %v\n", err)
		os.Exit(1)
	}

	return reader
}
