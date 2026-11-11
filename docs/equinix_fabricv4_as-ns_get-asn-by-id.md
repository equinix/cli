## equinix fabricv4 as-ns get-asn-by-id

Execute get-asn-by-id operation

### Synopsis

Execute the get-asn-by-id operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 as-ns get-asn-by-id [flags]
```

### Options

```
  -h, --help             help for get-asn-by-id
      --id string        id (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 as-ns](equinix_fabricv4_as-ns.md)	 - Manage as-ns resources

