// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"context"
	"fmt"
	"github.com/sjones4/eucalyptus-sdk-go/service/euprop"
	"github.com/spf13/cobra"
)

// getPropertiesCmd represents the get command
var getPropertiesCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a property value by name",
	Run: func(cmd *cobra.Command, args []string) {
		input := &euprop.DescribePropertiesInput{}
		DoInput(cmd, func(ccmd *CheckedCommand) {
			input.Properties = []string{ccmd.GetFlagString("name")}
		})
		svc := GetPropertiesSvc()
		request := svc.DescribePropertiesRequest(input)
		response, err := request.Send(context.Background())
		DoCommandError(cmd, err)
		for _, property := range response.DescribePropertiesOutput.Properties {
			if property.Name != nil && *property.Name == input.Properties[0] {
				fmt.Printf("PROPERTY\t%s\t%s\n", *property.Name, *property.Value)
				fmt.Printf("DESCRIPTION\t%s\n", *property.Description)
			}
		}
	},
}

func init() {
	propertiesCmd.AddCommand(getPropertiesCmd)

	getPropertiesCmd.Flags().String("name", "", "Name of the property to get (required)")
	_ = getPropertiesCmd.MarkFlagRequired("name")
}
