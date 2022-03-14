package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "dev-tools",
	Short: "A collection of tools for developers",
	Long:  "",
}

var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "A tool to encode or decode JWTs",
}

func Execute() {
	rootCmd.AddCommand(jwtCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
