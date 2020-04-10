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

// describeServicesCmd represents the describe command
var describeServicesCmd = &cobra.Command{
	Use:   "describe",
	Short: "List service states",
	Run: func(cmd *cobra.Command, args []string) {
		input := &euserv.DescribeServicesInput{}
		DoInput(cmd, func(ccmd *CheckedCommand) {
			input.ListAll = aws.Bool(ccmd.GetFlagBool("all"))
			input.Filters = ServiceFilters(ccmd)
		})
		svc := GetServicesSvc()
		request := svc.DescribeServicesRequest(input)
		response, err := request.Send(context.Background())
		DoCommandError(cmd, err)
		for _, service := range response.DescribeServicesOutput.ServiceStatuses {
			fmt.Printf("SERVICE\t%s\t%s\t%s\t%s\n",
				aws.StringValue(service.ServiceId.Type),
				aws.StringValue(service.ServiceId.Partition),
				aws.StringValue(service.ServiceId.Name),
				aws.StringValue(service.LocalState))
		}
	},
}

func init() {
	servicesCmd.AddCommand(describeServicesCmd)

	describeServicesCmd.Flags().BoolP("all", "a", false, "Show all services")

	ServiceFilterFlags(describeServicesCmd)
}
