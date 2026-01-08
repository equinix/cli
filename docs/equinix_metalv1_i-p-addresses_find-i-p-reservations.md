## equinix metalv1 i-p-addresses find-i-p-reservations

Execute find-i-p-reservations operation

### Synopsis

Execute the find-i-p-reservations operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 i-p-addresses find-i-p-reservations [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-i-p-reservations
      --include string   include field (JSON or string)
      --page int         page field
      --param-1 string   param-1 (required)
      --per-page int     per-page field
      --request string   JSON payload for additional optional fields not exposed as flags
      --types string     types field (JSON or string)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 i-p-addresses](equinix_metalv1_i-p-addresses.md)	 - Manage i-p-addresses resources

