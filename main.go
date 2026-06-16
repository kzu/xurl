package main

import (
	"fmt"
	"os"

	"github.com/kzu/xurl/auth"
	"github.com/kzu/xurl/cli"
	"github.com/kzu/xurl/config"
)

func main() {
	// Create a new config from environment variables
	config := config.NewConfig()
	auth := auth.NewAuth(config)

	// Create the root command
	rootCmd := cli.CreateRootCommand(config, auth)

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
