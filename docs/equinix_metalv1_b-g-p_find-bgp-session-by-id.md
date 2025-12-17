## equinix metalv1 b-g-p find-bgp-session-by-id

Execute find-bgp-session-by-id operation

### Synopsis

Execute the find-bgp-session-by-id operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 b-g-p find-bgp-session-by-id [flags]
```

### Options

```
  -h, --help             help for find-bgp-session-by-id
      --id string        id (required)
      --include string   include field (JSON or string)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 b-g-p](equinix_metalv1_b-g-p.md)	 - Manage b-g-p resources

