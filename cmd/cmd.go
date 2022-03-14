package cmd

import (
	"github.com/robbailey3/dev-tools/tcp"
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
	rootCmd.AddCommand(tcp.Commands())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
