## equinix metalv1 s-s-h-keys update-s-s-h-key

Execute update-s-s-h-key operation

### Synopsis

Execute the update-s-s-h-key operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 s-s-h-keys update-s-s-h-key [flags]
```

### Options

```
  -h, --help                                           help for update-s-s-h-key
      --include string                                 include field (JSON or string)
      --param-1 string                                 param-1 (required)
      --request string                                 JSON payload for additional optional fields not exposed as flags
      --s-s-h-key-input-additional-properties string   s-s-h-key-input-additional-properties (JSON)
      --s-s-h-key-input-key string                     s-s-h-key-input-key
      --s-s-h-key-input-label string                   s-s-h-key-input-label
      --s-s-h-key-input-tags string                    s-s-h-key-input-tags (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 s-s-h-keys](equinix_metalv1_s-s-h-keys.md)	 - Manage s-s-h-keys resources

