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

// describeServiceTypesCmd represents the describe service types command
var describeServiceTypesCmd = &cobra.Command{
	Use:   "describe-types",
	Short: "List the available service types",
	Run: func(cmd *cobra.Command, args []string) {
		input := &euserv.DescribeAvailableServiceTypesInput{}
		DoInput(cmd, func(ccmd *CheckedCommand) {
			input.Verbose = aws.Bool(ccmd.GetFlagBool("verbose"))
		})
		svc := GetServicesSvc()
		request := svc.DescribeAvailableServiceTypesRequest(input)
		response, err := request.Send(context.Background())
		DoCommandError(cmd, err)
		for _, serviceType := range response.DescribeAvailableServiceTypesOutput.Available {
			fmt.Printf("SVCTYPE\t%s\t%t\t%s\n",
				aws.StringValue(serviceType.ComponentName),
				len(serviceType.ServiceGroupMembers) > 0,
				aws.StringValue(serviceType.Description),
			)
		}
	},
}

func init() {
	servicesCmd.AddCommand(describeServiceTypesCmd)

	describeServiceTypesCmd.Flags().Bool("verbose", false, "List all available service types")
}
