// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sjones4/eucalyptus-sdk-go/service/euserv"
	"github.com/spf13/cobra"
	"log"
)

// modifyServiceCmd represents the modify service command
var modifyServiceCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify service state",
	Run: func(cmd *cobra.Command, args []string) {
		svc := GetServicesSvc()
		stateString, err := cmd.Flags().GetString("state")
		if err != nil {
			log.Fatalf("Error modifying service: %s", err.Error())
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("Error modifying service: %s", err.Error())
		}
		request := svc.ModifyServiceRequest(&euserv.ModifyServiceInput{
			Name:  &name,
			State: euserv.StateEnum(stateString),
		})
		resp, err := request.Send(context.Background())
		if err != nil {
			log.Fatalf("Error modifying service: %s", err.Error())
		}
		fmt.Printf("MODIFIED\t%t\n", aws.BoolValue(resp.ModifyServiceOutput.Metadata.Return))
	},
}

func init() {
	servicesCmd.AddCommand(modifyServiceCmd)

	modifyServiceCmd.Flags().String("name", "", "Service name")
	modifyServiceCmd.Flags().String("state", "", "Service target state")
	_ = modifyServiceCmd.MarkFlagRequired("name")
	_ = modifyServiceCmd.MarkFlagRequired("state")
}
