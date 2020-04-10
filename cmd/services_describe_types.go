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

// describeServiceTypesCmd represents the describe service types command
var describeServiceTypesCmd = &cobra.Command{
	Use:   "describe-types",
	Short: "List the available service types",
	Run: func(cmd *cobra.Command, args []string) {
		svc := GetServicesSvc()
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			log.Fatalf("Error describing service types: %s", err.Error())
		}
		request := svc.DescribeAvailableServiceTypesRequest(&euserv.DescribeAvailableServiceTypesInput{
			Verbose: &verbose,
		})
		resp, err := request.Send(context.Background())
		if err != nil {
			log.Fatalf("Error describing service types: %s", err.Error())
		}
		for _, serviceType := range resp.DescribeAvailableServiceTypesOutput.Available {
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
