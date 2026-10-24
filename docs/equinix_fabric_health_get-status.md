## equinix fabric health get-status

Get service status

### Synopsis

GET All service health statys with an option query parameter to return all Equinix Fabric customer in which the customer has a presence.

Use --request flag to provide optional JSON payload fields.

```
equinix fabric health get-status [flags]
```

### Options

```
  -h, --help             help for get-status
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabric health](equinix_fabric_health.md)	 - Manage health resources

