// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"context"
	"fmt"
	"github.com/sjones4/eucalyptus-sdk-go/service/euprop"
	"log"

	"github.com/spf13/cobra"
)

// listPropertiesCmd represents the list command
var listPropertiesCmd = &cobra.Command{
	Use:   "list",
	Short: "List property names and values",
	Run: func(cmd *cobra.Command, args []string) {
		svc := GetPropertiesSvc()
		propertyPrefix, err := cmd.Flags().GetString("property-prefix")
		if err != nil {
			log.Fatalf("Error describing properties: %s", err.Error())
		}
		var requestProperties []string
		if propertyPrefix != "" {
			requestProperties = []string{propertyPrefix}
		}
		request := svc.DescribePropertiesRequest(&euprop.DescribePropertiesInput{
			Properties: requestProperties,
		})
		resp, err := request.Send(context.Background())
		if err != nil {
			log.Fatalf("Error describing properties: %s", err.Error())
		}
		for _, property := range resp.DescribePropertiesOutput.Properties {
			fmt.Printf("PROPERTY\t%s\t%s\n", *property.Name, *property.Value)
		}
	},
}

func init() {
	propertiesCmd.AddCommand(listPropertiesCmd)

	listPropertiesCmd.Flags().String("property-prefix", "", "Property name prefix match")
}
