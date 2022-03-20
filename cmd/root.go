/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/micrictor/typo-generator/pkg/mapping"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "typo-generator target-key",
	Short: "Generate typos for a given character",
	Run:   getTypo,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("layout", "l", "qwerty", "Keyboard layout to generate typos for")
}

func getTypo(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("Expected target-key to generate typo for as an argument")
	}
	var targetCharacter rune = rune(args[0][0])
	layout, _ := cmd.Flags().GetString("layout")

	fmt.Printf("Finding typos for character '%c' using keyboard layout %s\n", targetCharacter, layout)

	m, err := mapping.New(layout)
	if err != nil {
		log.Fatalf("Failed to get keyboard map for layout!\n%v", err)
	}

	typos, err := m.FindTypos(targetCharacter)
	if err != nil {
		log.Fatalf("Failed to get typos for character!\n%v", err)
	}

	typoStrings := func(arr []rune) []string {
		var out []string
		for _, item := range arr {
			out = append(out, string(item))
		}
		return out
	}(typos)

	outputStr := strings.Join(typoStrings, "\n\t")

	log.Printf("Likely typos for '%c' on keyboard layout '%s':\n\t%s", targetCharacter, layout, outputStr)
}
