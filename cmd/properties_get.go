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

// getPropertiesCmd represents the get command
var getPropertiesCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a property value by name",
	Run: func(cmd *cobra.Command, args []string) {
		svc := GetPropertiesSvc()
		propertyName, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("Error getting property: %s", err.Error())
		}
		request := svc.DescribePropertiesRequest(&euprop.DescribePropertiesInput{
			Properties: []string{propertyName},
		})
		resp, err := request.Send(context.Background())
		if err != nil {
			log.Fatalf("Error getting property: %s", err.Error())
		}
		for _, property := range resp.DescribePropertiesOutput.Properties {
			if property.Name != nil && *property.Name == propertyName {
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
