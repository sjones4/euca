// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sjones4/eucalyptus-sdk-go/service/euserv"
	"github.com/spf13/cobra"
)

// deregisterServiceCmd represents the deregister service command
var deregisterServiceCmd = &cobra.Command{
	Use:   "deregister",
	Short: "Deregister a service",
	Run: func(cmd *cobra.Command, args []string) {
		input := &euserv.DeregisterServiceInput{}
		DoInput(cmd, func(ccmd *CheckedCommand) {
			input.Name = aws.String(ccmd.GetFlagString("name"))
			input.Type = aws.String(ccmd.GetFlagString("type"))
		})
		svc := GetServicesSvc()
		request := svc.DeregisterServiceRequest(input)
		response, err := request.Send(context.Background())
		DoCommandError(cmd, err)
		for _, serviceId := range response.DeregisterServiceOutput.DeregisteredServices {
			fmt.Printf("SERVICE\t%s\t%s\t%s\n",
				aws.StringValue(serviceId.Type),
				aws.StringValue(serviceId.Partition),
				aws.StringValue(serviceId.Name))
		}
	},
}

func init() {
	servicesCmd.AddCommand(deregisterServiceCmd)

	deregisterServiceCmd.Flags().String("name", "", "Service name")
	deregisterServiceCmd.Flags().String("type", "", "Service type")
	_ = deregisterServiceCmd.MarkFlagRequired("name")
}
