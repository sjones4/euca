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

// deregisterServiceCmd represents the deregister service command
var deregisterServiceCmd = &cobra.Command{
	Use:   "deregister",
	Short: "Deregister a service",
	Run: func(cmd *cobra.Command, args []string) {
		svc := GetServicesSvc()
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("Error deregistering service: %s", err.Error())
		}
		serviceType, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Fatalf("Error deregistering service: %s", err.Error())
		}
		request := svc.DeregisterServiceRequest(&euserv.DeregisterServiceInput{
			Name: &name,
			Type: &serviceType,
		})
		resp, err := request.Send(context.Background())
		if err != nil {
			log.Fatalf("Error deregistering service: %s", err.Error())
		}
		for _, serviceId := range resp.DeregisterServiceOutput.DeregisteredServices {
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
