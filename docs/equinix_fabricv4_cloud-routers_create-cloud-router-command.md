## equinix fabricv4 cloud-routers create-cloud-router-command

Initiate Command

### Synopsis

This API provides capability to initiate Command

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers create-cloud-router-command [flags]
```

### Options

```
      --cloud-router-command-post-request-additional-properties string           cloud-router-command-post-request-additional-properties (required) (JSON)
      --cloud-router-command-post-request-description string                     cloud-router-command-post-request-description
      --cloud-router-command-post-request-name string                            cloud-router-command-post-request-name
      --cloud-router-command-post-request-project-additional-properties string   cloud-router-command-post-request-project-additional-properties (required) (JSON)
      --cloud-router-command-post-request-project-project-id string              cloud-router-command-post-request-project-project-id (required)
      --cloud-router-command-post-request-request-additional-properties string   cloud-router-command-post-request-request-additional-properties (required) (JSON)
      --cloud-router-command-post-request-request-data-bytes int                 cloud-router-command-post-request-request-data-bytes
      --cloud-router-command-post-request-request-destination string             cloud-router-command-post-request-request-destination (required)
      --cloud-router-command-post-request-request-hops-max int                   cloud-router-command-post-request-request-hops-max
      --cloud-router-command-post-request-request-probes int                     cloud-router-command-post-request-request-probes
      --cloud-router-command-post-request-request-source-connection string       cloud-router-command-post-request-request-source-connection (JSON)
      --cloud-router-command-post-request-request-timeout int                    cloud-router-command-post-request-request-timeout
      --cloud-router-command-post-request-type string                            cloud-router-command-post-request-type (required)
  -h, --help                                                                     help for create-cloud-router-command
      --request string                                                           JSON payload for additional optional fields not exposed as flags
      --router-id string                                                         Router UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

