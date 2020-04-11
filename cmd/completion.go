// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// completionCmd is for bash command completion script generation
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate a bash command completion script",
	Long: `To load completion run

. <(euca completion)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(euca completion)
`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = rootCmd.GenBashCompletion(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
