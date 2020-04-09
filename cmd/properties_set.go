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

// setPropertiesCmd represents the set command
var setPropertiesCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a property value by name",
	Run: func(cmd *cobra.Command, args []string) {
		svc := GetPropertiesSvc()
		propertyName, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("Error modifying property %s\n", err.Error())
		}
		propertyValue, err := cmd.Flags().GetString("value")
		if err != nil {
			log.Fatalf("Error modifying property: %s\n", err.Error())
		}
		request := svc.ModifyPropertyValueRequest(&euprop.ModifyPropertyValueInput{
			Name:  &propertyName,
			Value: &propertyValue,
		})
		resp, err := request.Send(context.Background())
		if err != nil {
			log.Fatalf("Error modifying property: %s\n", err.Error())
		}
		fmt.Printf("PROPERTY\t%s\n", *resp.Name)
		fmt.Printf("OLDVALUE\t%s\n", *resp.OldValue)
		fmt.Printf("VALUE\t%s\n", *resp.Value)
	},
}

func init() {
	propertiesCmd.AddCommand(setPropertiesCmd)

	setPropertiesCmd.Flags().String("name", "", "Name of the property to set (required)")
	setPropertiesCmd.Flags().String("value", "", "Value for the property (required)")
	_ = setPropertiesCmd.MarkFlagRequired("name")
	_ = setPropertiesCmd.MarkFlagRequired("value")
}
