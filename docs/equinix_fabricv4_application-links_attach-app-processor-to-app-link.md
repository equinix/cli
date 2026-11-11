## equinix fabricv4 application-links attach-app-processor-to-app-link

Execute attach-app-processor-to-app-link operation

### Synopsis

Execute the attach-app-processor-to-app-link operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 application-links attach-app-processor-to-app-link [flags]
```

### Options

```
      --app-link-attach-processor-request-additional-properties string   app-link-attach-processor-request-additional-properties (JSON)
      --app-link-attach-processor-request-config string                  app-link-attach-processor-request-config
  -h, --help                                                             help for attach-app-processor-to-app-link
      --param-1 string                                                   param-1 (required)
      --param-2 string                                                   param-2 (required)
      --request string                                                   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-links](equinix_fabricv4_application-links.md)	 - Manage application-links resources

