package main

import (
	"fmt"

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

	rootCmd.Execute()
}
