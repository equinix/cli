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
      --cloud-router-command-post-request-additional-properties string           cloud-router-command-post-request-additional-properties (JSON)
      --cloud-router-command-post-request-description string                     Customer-provided Cloud Router Command description
      --cloud-router-command-post-request-name string                            Customer-provided Cloud Router Command name
      --cloud-router-command-post-request-project-additional-properties string   cloud-router-command-post-request-project-additional-properties (JSON)
      --cloud-router-command-post-request-project-project-id string              Subscriber-assigned project ID
      --cloud-router-command-post-request-request-additional-properties string   cloud-router-command-post-request-request-additional-properties (JSON)
      --cloud-router-command-post-request-request-data-bytes PING_COMMAND        Ping Command DataBytes.  This field is only applicable for commands of type PING_COMMAND.
      --cloud-router-command-post-request-request-destination string             Fabric Cloud Router Ping or Traceroute Command Destination
      --cloud-router-command-post-request-request-hops-max TRACEROUTE_COMMAND    Maximum number of hops for the traceroute command. This field is only applicable for commands of type TRACEROUTE_COMMAND.
      --cloud-router-command-post-request-request-probes TRACEROUTE_COMMAND      Number of probes for Fabric Cloud Router Traceroute Command. This field is only applicable for commands of type TRACEROUTE_COMMAND and is not configurable.
      --cloud-router-command-post-request-request-source-connection string       cloud-router-command-post-request-request-source-connection (JSON)
      --cloud-router-command-post-request-request-timeout PING_COMMAND           Timeout in seconds for Fabric Cloud Router Command:   - For PING_COMMAND: Packet timeout duration. The default value is 5.   - For `TRACEROUTE_COMMAND`: Probe timeout duration.     The default value is 2 and it is not configurable.
      --cloud-router-command-post-request-type string                            cloud-router-command-post-request-type
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

