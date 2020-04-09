## Eucalyptus CLI

### Building

    git checkout https://github.com/sjones4/euca.git
    cd euca
    go build
    ./euca version

### Usage

The CLI currently supports the Eucalyptus Cloud properties administrative service:

    euca properties list --property-prefix dns
    euca properties get --name dns.enabled
    euca properties set --name dns.enabled --value true

The following global options (flags) are available:

* --debug        : Output debug logging for service requests and responses
* --endpoint-url : Specify the service endpoint
* --profile      : Specify the AWS SDK profile to use (e.g. for credentials)

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

    EUCA_PROPERTIES_URL

A configuration file can be created in _~/.euca/cli.yaml_ with global options to use for each run. e.g.:

    endpoint-url-suffix: cloud-10-10-10-10.euca.me:8773
    endpoint-protocol: http

The options shown above allow service endpoints to be derived per service.