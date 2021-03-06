## Eucalyptus CLI

[![Build Status](https://dev.azure.com/sjones4/eucalyptus/_apis/build/status/euca-go-ci?branchName=master)](https://dev.azure.com/sjones4/eucalyptus/_build/latest?definitionId=20&branchName=master)

A Eucalyptus Cloud CLI written in Go using the [Eucalyptus SDK for Go](https://github.com/sjones4/eucalyptus-sdk-go)

Binaries for Linux (amd64) and Windows can be downloaded from the Azure Pipelines artifacts. 

### Building

    git checkout https://github.com/sjones4/euca.git
    cd euca
    go build
    ./euca version

### Usage

The CLI currently supports administration for Eucalyptus Cloud properties and services:

    euca properties list --property-prefix dns
    euca properties get --name dns.enabled
    euca properties set --name dns.enabled --value true
    
    euca services describe
    euca services describe-certificates
    euca services describe-types
    euca services modify --name compute-1 --state STOPPED
    euca services register --type compute --partition api --host 10.10.10.10 --name compute-1
    euca services deregister --name compute-1

The following global options (flags) are available:

* --debug        : Output debug logging for service requests and responses
* --endpoint-url : Specify the service endpoint
* --profile      : Specify the AWS SDK profile to use (e.g. for credentials)

For command completion with a bash shell run:

    . <(euca completion)

For command completion support you may need additional dependencies (e.g. the _bash-completion_ rpm on CentOS 7)

### Configuration

The CLI uses the AWS SDK which uses the usual environment variables:

    AWS_CA_BUNDLE
    AWS_CONFIG_FILE
    AWS_SHARED_CREDENTIALS_FILE
    AWS_PROFILE
    AWS_ACCESS_KEY_ID
    AWS_SECRET_ACCESS_KEY
    AWS_SESSION_TOKEN

The following are CLI specific environment variables:

    EUCA_BOOTSTRAP_URL
    EUCA_PROPERTIES_URL

A configuration file can be created in _~/.euca/cli.yaml_ with global options to use for each run. e.g.:

    endpoint-url-suffix: cloud-10-10-10-10.euca.me:8773
    endpoint-protocol: http

The options shown above allow service endpoints to be derived per service.