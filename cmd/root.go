// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
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

type CheckedCommandCallback func(cmd *CheckedCommand)

type CheckedCommand struct {
	cmd *cobra.Command
	err error
}

func (cmd *CheckedCommand) GetFlagString(name string) (value string) {
	if cmd.err == nil {
		value, cmd.err = cmd.cmd.Flags().GetString(name)
	}
	return
}

func (cmd *CheckedCommand) GetFlagInt(name string) (value int) {
	if cmd.err == nil {
		value, cmd.err = cmd.cmd.Flags().GetInt(name)
	}
	return
}

func (cmd *CheckedCommand) GetFlagBool(name string) (value bool) {
	if cmd.err == nil {
		value, cmd.err = cmd.cmd.Flags().GetBool(name)
	}
	return
}

func WithCheckedCommand(cmd *cobra.Command, callback CheckedCommandCallback) error {
	ccmd := &CheckedCommand{cmd: cmd}
	callback(ccmd)
	return ccmd.err
}

func DoInput(cmd *cobra.Command, callback CheckedCommandCallback) {
	err := WithCheckedCommand(cmd, callback)
	DoCommandError(cmd, err)
}

func DoCommandError(cmd *cobra.Command, err error) {
	if err != nil {
		log.Fatalf("Error in %s commmand: %s\n", cmd.Name(), err.Error())
	}
}

func GetEndpoint(envKey string, dnsLabel string, defaultEndpoint string) string {
	endpoint := viper.GetString(ConfigEndpointUrl)

	if endpoint == "" {
		endpoint = os.Getenv(envKey)
	}

	endpointUrlSuffix := viper.GetString(ConfigEndpointUrlSuffix)
	endpointProtocol := viper.GetString(ConfigEndpointProtocol)
	if endpoint == "" && endpointUrlSuffix != "" {
		endpoint = strings.Join([]string{endpointProtocol, "://", dnsLabel, ".", endpointUrlSuffix}, "")
	}

	if endpoint == "" {
		endpoint = defaultEndpoint
	}

	endpoint = strings.TrimSuffix(endpoint, "/")

	return endpoint
}

func GetAwsConfig(endpoint string) aws.Config {
	cfg, err := external.LoadDefaultAWSConfig(
		external.WithEndpointResolverFunc(func(aws.EndpointResolver) aws.EndpointResolver {
			return aws.ResolveWithEndpointURL(endpoint)
		}),
		external.WithSharedConfigProfile(viper.GetString(ConfigProfile)),
		external.WithRegion(viper.GetString(ConfigRegion)),
	)
	if err != nil {
		log.Fatalf("Error with default configuration: %s", err.Error())
	}
	if debugLogging {
		cfg.LogLevel = aws.LogDebugWithHTTPBody
	}
	return cfg
}
