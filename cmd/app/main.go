package main

import (
	"fmt"
	"os"

	"github.com/AFelipeTrujillo/vidgo/internal/Infrastructure/Delivery/Console"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "vidgo",
		Short: "Vidgo is a CLI tool for video manipulation",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Vidgo is ready to process some videos!")
		},
	}

	rootCmd.AddCommand(Console.NewClipCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
