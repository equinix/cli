## equinix metalv1 i-p-addresses find-i-p-availabilities

Execute find-i-p-availabilities operation

### Synopsis

Execute the find-i-p-availabilities operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 i-p-addresses find-i-p-availabilities [flags]
```

### Options

```
      --cidr string      cidr field
  -h, --help             help for find-i-p-availabilities
      --param-1 string   param-1 (required)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 i-p-addresses](equinix_metalv1_i-p-addresses.md)	 - Manage i-p-addresses resources

