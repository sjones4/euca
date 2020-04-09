// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	ConfigEndpointProtocol  = "endpoint-protocol"
	ConfigEndpointUrl       = "endpoint-url"
	ConfigEndpointUrlSuffix = "endpoint-url-suffix"
	ConfigProfile           = "profile"
	ConfigRegion            = "region"
)

var (
	cfgFile      string
	debugLogging bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "euca",
	Short:   "A Command Line Interface for Eucalyptus Clouds",
	Version: Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVar(&debugLogging, "debug", false, "Enable debug logging")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.euca/cli.yaml)")

	rootCmd.PersistentFlags().String(ConfigEndpointProtocol, "https", "Protocol to use for derived endpoint URL")
	rootCmd.PersistentFlags().String(ConfigEndpointUrl, "", "Override command's default URL with the given URL")
	rootCmd.PersistentFlags().String(ConfigEndpointUrlSuffix, "", "URL suffix to append to service label")
	rootCmd.PersistentFlags().String(ConfigProfile, "", "AWS configuration profile to use")
	rootCmd.PersistentFlags().String(ConfigRegion, "eucalyptus", "AWS region to use")

	_ = viper.BindPFlag(ConfigEndpointProtocol, rootCmd.PersistentFlags().Lookup(ConfigEndpointProtocol))
	_ = viper.BindPFlag(ConfigEndpointUrlSuffix, rootCmd.PersistentFlags().Lookup(ConfigEndpointUrlSuffix))
	_ = viper.BindPFlag(ConfigProfile, rootCmd.PersistentFlags().Lookup(ConfigProfile))
	_ = viper.BindPFlag(ConfigRegion, rootCmd.PersistentFlags().Lookup(ConfigRegion))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(fmt.Sprintf("%s/.euca", home))
		viper.SetConfigName("cli")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading configuration %s: %s\n", viper.ConfigFileUsed(), err.Error())
	}
}

func GetAwsConfig(endpoint string) (aws.Config, error) {
	cfg, err := external.LoadDefaultAWSConfig(
		external.WithEndpointResolverFunc(func(aws.EndpointResolver) aws.EndpointResolver { return aws.ResolveWithEndpointURL(endpoint) }),
		external.WithSharedConfigProfile(viper.GetString(ConfigProfile)),
		external.WithRegion(viper.GetString(ConfigRegion)),
	)
	if err == nil && debugLogging {
		cfg.LogLevel = aws.LogDebugWithHTTPBody
	}
	return cfg, err
}
