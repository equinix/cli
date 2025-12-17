## equinix metalv1 v-r-fs delete-bgp-dynamic-neighbor-by-id

Execute delete-bgp-dynamic-neighbor-by-id operation

### Synopsis

Execute the delete-bgp-dynamic-neighbor-by-id operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 v-r-fs delete-bgp-dynamic-neighbor-by-id [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for delete-bgp-dynamic-neighbor-by-id
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

* [equinix metalv1 v-r-fs](equinix_metalv1_v-r-fs.md)	 - Manage v-r-fs resources

