## equinix fabricv4 agent-templates get-agent-template-by-uuid

Execute get-agent-template-by-uuid operation

### Synopsis

Execute the get-agent-template-by-uuid operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 agent-templates get-agent-template-by-uuid [flags]
```

### Options

```
  -h, --help             help for get-agent-template-by-uuid
      --limit int        limit field
      --offset int       offset field
      --request string   JSON payload for additional optional fields not exposed as flags
      --uuid string      uuid (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 agent-templates](equinix_fabricv4_agent-templates.md)	 - Manage agent-templates resources

