## equinix fabricv4 tags list-tags

Execute list-tags operation

### Synopsis

Execute the list-tags operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 tags list-tags [flags]
```

### Options

```
  -h, --help             help for list-tags
      --limit int        limit field
      --offset int       offset field
      --request string   JSON payload for additional optional fields not exposed as flags
      --type_ string     type_ field (JSON or string)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 tags](equinix_fabricv4_tags.md)	 - Manage tags resources

