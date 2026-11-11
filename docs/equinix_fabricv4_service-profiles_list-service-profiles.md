## equinix fabricv4 service-profiles list-service-profiles

Execute list-service-profiles operation

### Synopsis

Execute the list-service-profiles operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 service-profiles list-service-profiles [flags]
```

### Options

```
  -h, --help                help for list-service-profiles
      --limit int           limit field
      --offset int          offset field
      --request string      JSON payload for additional optional fields not exposed as flags
      --view-point string   view-point field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

