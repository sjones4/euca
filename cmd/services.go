// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sjones4/eucalyptus-sdk-go/service/euserv"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strings"
)

const (
	BootstrapDnsLabel                = "bootstrap"
	DefaultBootstrapServicesEndpoint = "http://127.0.0.1:8773/services/Empyrean"
	EnvBootstrapServiceUrl           = "EUCA_BOOTSTRAP_URL"
)

// servicesCmd is the parent for all service management service commands
var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Eucalyptus cloud service management service",
}

func init() {
	rootCmd.AddCommand(servicesCmd)
}

func GetServicesEndpoint() string {
	return GetEndpoint(EnvBootstrapServiceUrl, BootstrapDnsLabel, DefaultBootstrapServicesEndpoint)
}

func GetServicesSvc() *euserv.Client {
	servicesEndpoint := GetServicesEndpoint()
	cfg := GetAwsConfig(servicesEndpoint)
	return euserv.New(cfg)
}

func ServiceFilterFlags(cmd *cobra.Command) {
	cmd.Flags().String("filter:certificate-usage", "", "Filter results by certificate usage")
	cmd.Flags().String("filter:host", "", "Filter results by host")
	cmd.Flags().String("filter:internal", "", "Filter results by internal service flag (true|false)")
	cmd.Flags().String("filter:partition", "", "Filter results by zone or partition")
	cmd.Flags().String("filter:public", "", "Filter results by public service flag (true|false)")
	cmd.Flags().String("filter:service-group", "", "Filter results by membership of service group")
	cmd.Flags().String("filter:service-group-member", "", "Filter results by service group membership (true|false)")
	cmd.Flags().String("filter:service-type", "", "Filter results on service type")
	cmd.Flags().String("filter:state", "", "Filter results by service state")
}

func ServiceFilters(ccmd *CheckedCommand) []euserv.Filter {
	filterMap := map[string][]string{}
	ccmd.cmd.Flags().Visit(func(flag *pflag.Flag) {
		if strings.HasPrefix(flag.Name, "filter:") {
			filterName := strings.TrimPrefix(flag.Name, "filter:")
			filterValues, ok := filterMap[filterName]
			if !ok {
				filterValues = []string{}
			}
			filterValues = append(filterValues, flag.Value.String())
			filterMap[filterName] = filterValues
		}
	})
	filters := []euserv.Filter{}
	for name, values := range filterMap {
		filters = append(filters, euserv.Filter{Name: aws.String(name), Values: values})
	}
	return filters
}
