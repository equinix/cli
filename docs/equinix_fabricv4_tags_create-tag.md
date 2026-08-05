## equinix fabricv4 tags create-tag

Execute create-tag operation

### Synopsis

Execute the create-tag operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 tags create-tag [flags]
```

### Options

```
  -h, --help                                       help for create-tag
      --request string                             JSON payload for additional optional fields not exposed as flags
      --tag-request-additional-properties string   tag-request-additional-properties (JSON)
      --tag-request-display-name string            tag-request-display-name
      --tag-request-name string                    tag-request-name
      --tag-request-type string                    tag-request-type
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 tags](equinix_fabricv4_tags.md)	 - Manage tags resources

