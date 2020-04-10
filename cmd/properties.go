// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"github.com/sjones4/eucalyptus-sdk-go/service/euprop"
	"github.com/spf13/cobra"
)

const (
	PropertiesDnsLabel        = "properties"
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
	return GetEndpoint(EnvPropertiesServiceUrl, PropertiesDnsLabel, DefaultPropertiesEndpoint)
}

func GetPropertiesSvc() *euprop.Client {
	propertiesEndpoint := GetPropertiesEndpoint()
	cfg := GetAwsConfig(propertiesEndpoint)
	return euprop.New(cfg)
}
