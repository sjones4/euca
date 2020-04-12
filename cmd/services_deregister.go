// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sjones4/eucalyptus-sdk-go/service/euserv"
	"github.com/spf13/cobra"
)

// deregisterServiceCmd represents the deregister service command
var deregisterServiceCmd = &cobra.Command{
	Use:           "deregister",
	Short:         "Deregister a service",
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		input := &euserv.DeregisterServiceInput{}
		DoInput(cmd, func(ccmd *CheckedCommand) {
			input.Name = aws.String(ccmd.GetFlagString("name"))
			input.Type = aws.String(ccmd.GetFlagString("type"))
		})
		svc := GetServicesSvc()
		request := svc.DeregisterServiceRequest(input)
		response, err := request.Send(context.Background())
		DoCommandError(cmd, err)
		deregistered := aws.BoolValue(response.RegistrationMetadata.ReponseMetadata.Return)
		fmt.Printf("DEREGISTERED\t%t\n", deregistered)
		if deregistered {
			for _, serviceId := range response.DeregisterServiceOutput.DeregisteredServices {
				fmt.Printf("SERVICE\t%s\t%s\t%s\n",
					aws.StringValue(serviceId.Type),
					aws.StringValue(serviceId.Partition),
					aws.StringValue(serviceId.Name))
			}
		} else {
			for _, message := range response.RegistrationMetadata.StatusMessages {
				fmt.Printf("MESSAGE\t%s\n", aws.StringValue(message.Entry))
			}
			cmd.SilenceUsage = true
			return errors.New("deregistration failed")
		}
		return nil
	},
}

func init() {
	servicesCmd.AddCommand(deregisterServiceCmd)

	deregisterServiceCmd.Flags().String("name", "", "Service name")
	deregisterServiceCmd.Flags().String("type", "", "Service type")
	_ = deregisterServiceCmd.MarkFlagRequired("name")
}
