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

// registerServiceCmd represents the register service command
var registerServiceCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a service",
	Run: func(cmd *cobra.Command, args []string) {
		svc := GetServicesSvc()
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("Error registering service: %s", err.Error())
		}
		serviceType, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Fatalf("Error registering service: %s", err.Error())
		}
		partition, err := cmd.Flags().GetString("partition")
		if err != nil {
			log.Fatalf("Error registering service: %s", err.Error())
		}
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatalf("Error registering service: %s", err.Error())
		}
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatalf("Error registering service: %s", err.Error())
		}
		request := svc.RegisterServiceRequest(&euserv.RegisterServiceInput{
			Name:      &name,
			Type:      &serviceType,
			Host:      &host,
			Partition: &partition,
			Port:      aws.Int64(int64(port)),
		})
		resp, err := request.Send(context.Background())
		if err != nil {
			log.Fatalf("Error registering service: %s", err.Error())
		}
		for _, serviceId := range resp.RegisterServiceOutput.RegisteredServices {
			fmt.Printf("SERVICE\t%s\t%s\t%s\n",
				aws.StringValue(serviceId.Type),
				aws.StringValue(serviceId.Partition),
				aws.StringValue(serviceId.Name))
		}
	},
}

func init() {
	servicesCmd.AddCommand(registerServiceCmd)

	registerServiceCmd.Flags().String("name", "", "Service name")
	registerServiceCmd.Flags().String("type", "", "Service type")
	registerServiceCmd.Flags().String("partition", "", "Zone or group name")
	registerServiceCmd.Flags().String("host", "", "Host for the service")
	registerServiceCmd.Flags().Int("port", 8773, "Port for the service")

	_ = registerServiceCmd.MarkFlagRequired("name")
	_ = registerServiceCmd.MarkFlagRequired("type")
	_ = registerServiceCmd.MarkFlagRequired("host")
}
