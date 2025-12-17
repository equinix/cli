## equinix metalv1 s-s-h-keys create-s-s-h-key

Execute create-s-s-h-key operation

### Synopsis

Execute the create-s-s-h-key operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 s-s-h-keys create-s-s-h-key [flags]
```

### Options

```
  -h, --help                                                  help for create-s-s-h-key
      --include string                                        include field (JSON or string)
      --request string                                        JSON payload for additional optional fields not exposed as flags
      --s-s-h-key-create-input-additional-properties string   s-s-h-key-create-input-additional-properties (JSON)
      --s-s-h-key-create-input-instances_ids string           List of instance UUIDs to associate SSH key with, when empty array is sent all instances belonging       to entity will be included (JSON array)
      --s-s-h-key-create-input-key string                     s-s-h-key-create-input-key
      --s-s-h-key-create-input-label string                   s-s-h-key-create-input-label
      --s-s-h-key-create-input-tags string                    s-s-h-key-create-input-tags (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 s-s-h-keys](equinix_metalv1_s-s-h-keys.md)	 - Manage s-s-h-keys resources

