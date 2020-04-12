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

// modifyServiceCmd represents the modify service command
var modifyServiceCmd = &cobra.Command{
	Use:           "modify",
	Short:         "Modify service state",
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		input := &euserv.ModifyServiceInput{}
		DoInput(cmd, func(ccmd *CheckedCommand) {
			input.Name = aws.String(ccmd.GetFlagString("name"))
			input.State = euserv.StateEnum(ccmd.GetFlagString("state"))
		})
		svc := GetServicesSvc()
		request := svc.ModifyServiceRequest(input)
		response, err := request.Send(context.Background())
		DoCommandError(cmd, err)
		modified := aws.BoolValue(response.ModifyServiceOutput.Metadata.Return)
		fmt.Printf("MODIFIED\t%t\n", modified)
		if !modified {
			cmd.SilenceUsage = true
			return errors.New("modify failed")
		}
		return nil
	},
}

func init() {
	servicesCmd.AddCommand(modifyServiceCmd)

	modifyServiceCmd.Flags().String("name", "", "Service name")
	modifyServiceCmd.Flags().String("state", "", "Service target state")
	_ = modifyServiceCmd.MarkFlagRequired("name")
	_ = modifyServiceCmd.MarkFlagRequired("state")
}
