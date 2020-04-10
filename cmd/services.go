// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"github.com/sjones4/eucalyptus-sdk-go/service/euserv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

const (
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
	servicesEndpoint := viper.GetString(ConfigEndpointUrl)

	if servicesEndpoint == "" {
		servicesEndpoint = os.Getenv(EnvBootstrapServiceUrl)
	}

	endpointUrlSuffix := viper.GetString(ConfigEndpointUrlSuffix)
	endpointProtocol := viper.GetString(ConfigEndpointProtocol)
	if servicesEndpoint == "" && endpointUrlSuffix != "" {
		servicesEndpoint = strings.Join([]string{endpointProtocol, "://bootstrap.", endpointUrlSuffix}, "")
	}

	if servicesEndpoint == "" {
		servicesEndpoint = DefaultBootstrapServicesEndpoint
	}

	servicesEndpoint = strings.TrimSuffix(servicesEndpoint, "/")

	return servicesEndpoint
}

func GetServicesSvc() *euserv.Client {
	servicesEndpoint := GetServicesEndpoint()
	cfg, err := GetAwsConfig(servicesEndpoint)
	if err != nil {
		log.Fatalf("Error with default configuration: %s", err.Error())
	}
	return euserv.New(cfg)
}
