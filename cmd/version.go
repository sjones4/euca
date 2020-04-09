// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	Version = "2020.4a"
)

// versionCmd outputs the software version number
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
