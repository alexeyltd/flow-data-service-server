package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "starter",
	Short: "Quick starter",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
