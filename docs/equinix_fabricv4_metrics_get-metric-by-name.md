## equinix fabricv4 metrics get-metric-by-name

Execute get-metric-by-name operation

### Synopsis

Execute the get-metric-by-name operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 metrics get-metric-by-name [flags]
```

### Options

```
  -h, --help             help for get-metric-by-name
      --limit int        limit field
      --name string      name field
      --offset int       offset field
      --request string   JSON payload for additional optional fields not exposed as flags
      --value string     value field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 metrics](equinix_fabricv4_metrics.md)	 - Manage metrics resources

