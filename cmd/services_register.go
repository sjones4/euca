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

// registerServiceCmd represents the register service command
var registerServiceCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a service",
	Run: func(cmd *cobra.Command, args []string) {
		input := &euserv.RegisterServiceInput{}
		DoInput(cmd, func(ccmd *CheckedCommand) {
			input.Name = aws.String(ccmd.GetFlagString("name"))
			input.Type = aws.String(ccmd.GetFlagString("type"))
			input.Host = aws.String(ccmd.GetFlagString("host"))
			input.Port = aws.Int64(int64(ccmd.GetFlagInt("port")))
			input.Partition = aws.String(ccmd.GetFlagString("partition"))
		})
		svc := GetServicesSvc()
		request := svc.RegisterServiceRequest(input)
		response, err := request.Send(context.Background())
		DoCommandError(cmd, err)
		for _, serviceId := range response.RegisterServiceOutput.RegisteredServices {
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
