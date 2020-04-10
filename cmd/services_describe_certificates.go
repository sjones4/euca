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
	"strings"
)

// describeServiceCertificatesCmd represents the describe service certificates command
var describeServiceCertificatesCmd = &cobra.Command{
	Use:   "describe-certificates",
	Short: "List service certificates",
	Run: func(cmd *cobra.Command, args []string) {
		svc := GetServicesSvc()
		formatString, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Fatalf("Error describing service certificates: %s", err.Error())
		}
		digest, err := cmd.Flags().GetString("digest")
		if err != nil {
			log.Fatalf("Error describing service certificates: %s", err.Error())
		}
		outputCertificate, err := cmd.Flags().GetBool("certificate")
		if err != nil {
			log.Fatalf("Error describing service certificates: %s", err.Error())
		}
		request := svc.DescribeServiceCertificatesRequest(&euserv.DescribeServiceCertificatesInput{
			Format:            euserv.CertificateFormatEnum(formatString),
			FingerprintDigest: &digest,
		})
		resp, err := request.Send(context.Background())
		if err != nil {
			log.Fatalf("Error describing service certificates: %s", err.Error())
		}
		for _, serviceCertificate := range resp.DescribeServiceCertificatesOutput.ServiceCertificates {
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
}
