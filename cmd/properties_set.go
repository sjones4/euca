// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sjones4/eucalyptus-sdk-go/service/euprop"
	"github.com/spf13/cobra"
)

// setPropertiesCmd represents the set command
var setPropertiesCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a property value by name",
	Run: func(cmd *cobra.Command, args []string) {
		input := &euprop.ModifyPropertyValueInput{}
		DoInput(cmd, func(ccmd *CheckedCommand) {
			input.Name = aws.String(ccmd.GetFlagString("name"))
			input.Value = aws.String(ccmd.GetFlagString("value"))
		})
		svc := GetPropertiesSvc()
		request := svc.ModifyPropertyValueRequest(input)
		response, err := request.Send(context.Background())
		DoCommandError(cmd, err)
		fmt.Printf("PROPERTY\t%s\n", *response.Name)
		fmt.Printf("OLDVALUE\t%s\n", *response.OldValue)
		fmt.Printf("VALUE\t%s\n", *response.Value)
	},
}

func init() {
	propertiesCmd.AddCommand(setPropertiesCmd)

	setPropertiesCmd.Flags().String("name", "", "Name of the property to set (required)")
	setPropertiesCmd.Flags().String("value", "", "Value for the property (required)")
	_ = setPropertiesCmd.MarkFlagRequired("name")
	_ = setPropertiesCmd.MarkFlagRequired("value")
}
