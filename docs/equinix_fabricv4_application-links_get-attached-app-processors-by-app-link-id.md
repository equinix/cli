## equinix fabricv4 application-links get-attached-app-processors-by-app-link-id

Execute get-attached-app-processors-by-app-link-id operation

### Synopsis

Execute the get-attached-app-processors-by-app-link-id operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 application-links get-attached-app-processors-by-app-link-id [flags]
```

### Options

```
      --attachment-status string   attachment-status field (JSON or string)
  -h, --help                       help for get-attached-app-processors-by-app-link-id
      --limit int                  limit field
      --offset int                 offset field
      --order string               order field
      --param-1 string             param-1 (required)
      --request string             JSON payload for additional optional fields not exposed as flags
      --style string               style field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-links](equinix_fabricv4_application-links.md)	 - Manage application-links resources

