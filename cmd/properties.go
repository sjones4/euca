// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"github.com/sjones4/eucalyptus-sdk-go/service/euprop"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	DefaultPropertiesEndpoint = "http://127.0.0.1:8773/services/Properties"
	EnvPropertiesServiceUrl   = "EUCA_PROPERTIES_URL"
)

// propertiesCmd is the parent for all properties service commands
var propertiesCmd = &cobra.Command{
	Use:   "properties",
	Short: "Eucalyptus cloud properties service",
}

func init() {
	rootCmd.AddCommand(propertiesCmd)
}

func GetPropertiesEndpoint() string {
	propertiesEndpoint := viper.GetString(ConfigEndpointUrl)

	if propertiesEndpoint == "" {
		propertiesEndpoint = os.Getenv(EnvPropertiesServiceUrl)
	}

	endpointUrlSuffix := viper.GetString(ConfigEndpointUrlSuffix)
	endpointProtocol := viper.GetString(ConfigEndpointProtocol)
	if propertiesEndpoint == "" && endpointUrlSuffix != "" {
		propertiesEndpoint = strings.Join([]string{endpointProtocol, "://properties.", endpointUrlSuffix}, "")
	}

	if propertiesEndpoint == "" {
		propertiesEndpoint = DefaultPropertiesEndpoint
	}

	propertiesEndpoint = strings.TrimSuffix(propertiesEndpoint, "/")

	return propertiesEndpoint
}

func GetPropertiesSvc() *euprop.Client {
	propertiesEndpoint := GetPropertiesEndpoint()
	cfg, err := GetAwsConfig(propertiesEndpoint)
	if err != nil {
		log.Fatalf("Error with default configuration: %s", err.Error())
	}
	return euprop.New(cfg)
}
