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

// describeServicesCmd represents the describe command
var describeServicesCmd = &cobra.Command{
	Use:   "describe",
	Short: "List service states",
	Run: func(cmd *cobra.Command, args []string) {
		svc := GetServicesSvc()
		listAll, err := cmd.Flags().GetBool("all")
		if err != nil {
			log.Fatalf("Error describing services: %s", err.Error())
		}
		request := svc.DescribeServicesRequest(&euserv.DescribeServicesInput{
			ListAll: &listAll,
		})
		resp, err := request.Send(context.Background())
		if err != nil {
			log.Fatalf("Error describing services: %s", err.Error())
		}
		for _, service := range resp.DescribeServicesOutput.ServiceStatuses {
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
}
