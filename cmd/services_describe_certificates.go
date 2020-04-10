// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sjones4/eucalyptus-sdk-go/service/euserv"
	"github.com/spf13/cobra"
	"strings"
)

// describeServiceCertificatesCmd represents the describe service certificates command
var describeServiceCertificatesCmd = &cobra.Command{
	Use:   "describe-certificates",
	Short: "List service certificates",
	Run: func(cmd *cobra.Command, args []string) {
		input := &euserv.DescribeServiceCertificatesInput{}
		outputCertificate := false
		DoInput(cmd, func(ccmd *CheckedCommand) {
			input.Format = euserv.CertificateFormatEnum(ccmd.GetFlagString("format"))
			input.FingerprintDigest = aws.String(ccmd.GetFlagString("digest"))
			input.Filters = ServiceFilters(ccmd)

			outputCertificate = ccmd.GetFlagBool("certificate")
		})
		svc := GetServicesSvc()
		request := svc.DescribeServiceCertificatesRequest(input)
		response, err := request.Send(context.Background())
		DoCommandError(cmd, err)
		for _, serviceCertificate := range response.DescribeServiceCertificatesOutput.ServiceCertificates {
			fmt.Printf("CERTIFICATE\t%s\t%s\t%s\t%s\n",
				aws.StringValue(serviceCertificate.ServiceType),
				aws.StringValue(serviceCertificate.CertificateUsage),
				aws.StringValue(serviceCertificate.CertificateFingerprintDigest),
				aws.StringValue(serviceCertificate.CertificateFingerprint),
			)
			if outputCertificate {
				fmt.Printf("%s\t%s\n",
					strings.ToUpper(aws.StringValue(serviceCertificate.CertificateFormat)),
					aws.StringValue(serviceCertificate.Certificate),
				)
			}
		}
	},
}

func init() {
	servicesCmd.AddCommand(describeServiceCertificatesCmd)

	describeServiceCertificatesCmd.Flags().String("format", "pem", "Certificate format")
	describeServiceCertificatesCmd.Flags().String("digest", "", "Certificate fingerprint digest")
	describeServiceCertificatesCmd.Flags().Bool("certificate", false, "Output full certificates")

	ServiceFilterFlags(describeServiceCertificatesCmd)
}
