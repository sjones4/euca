// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"github.com/sjones4/eucalyptus-sdk-go/service/euserv"
	"github.com/spf13/cobra"
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
