/*
Copyright © 2024 Rohit Kumar <kumar1rohit@outlook.com>
*/
package cmd

import (
	"os"

	"github.com/rohit1kumar/reacher-go/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "reacher-go <to_email>",
	Short: "Verify email addresses without sending emails, reacher-go <to_email>",
	Long: `reacher-go verifies email address validity without sending actual emails.
Example:
reacher-go test@example.com
	`,
	Example: `reacher-go test@example.com`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.CheckEmail(args[0])
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
