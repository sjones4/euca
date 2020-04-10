// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"context"
	"fmt"
	"github.com/sjones4/eucalyptus-sdk-go/service/euprop"
	"github.com/spf13/cobra"
)

// listPropertiesCmd represents the list command
var listPropertiesCmd = &cobra.Command{
	Use:   "list",
	Short: "List property names and values",
	Run: func(cmd *cobra.Command, args []string) {
		input := &euprop.DescribePropertiesInput{}
		DoInput(cmd, func(ccmd *CheckedCommand) {
			propertyPrefix := ccmd.GetFlagString("property-prefix")
			if propertyPrefix != "" {
				input.Properties = []string{propertyPrefix}
			}
		})
		svc := GetPropertiesSvc()
		request := svc.DescribePropertiesRequest(input)
		response, err := request.Send(context.Background())
		DoCommandError(cmd, err)
		for _, property := range response.DescribePropertiesOutput.Properties {
			fmt.Printf("PROPERTY\t%s\t%s\n", *property.Name, *property.Value)
		}
	},
}

func init() {
	propertiesCmd.AddCommand(listPropertiesCmd)

	listPropertiesCmd.Flags().String("property-prefix", "", "Property name prefix match")
}
