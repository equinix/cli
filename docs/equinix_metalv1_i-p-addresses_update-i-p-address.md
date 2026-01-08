## equinix metalv1 i-p-addresses update-i-p-address

Execute update-i-p-address operation

### Synopsis

Execute the update-i-p-address operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 i-p-addresses update-i-p-address [flags]
```

### Options

```
      --exclude string                                             exclude field (JSON or string)
  -h, --help                                                       help for update-i-p-address
      --i-p-assignment-update-input-additional-properties string   i-p-assignment-update-input-additional-properties (JSON)
      --i-p-assignment-update-input-customdata string              i-p-assignment-update-input-customdata (JSON)
      --i-p-assignment-update-input-details string                 i-p-assignment-update-input-details
      --i-p-assignment-update-input-tags string                    i-p-assignment-update-input-tags (JSON array)
      --include string                                             include field (JSON or string)
      --param-1 string                                             param-1 (required)
      --request string                                             JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 i-p-addresses](equinix_metalv1_i-p-addresses.md)	 - Manage i-p-addresses resources

