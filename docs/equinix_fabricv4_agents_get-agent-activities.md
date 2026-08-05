## equinix fabricv4 agents get-agent-activities

Execute get-agent-activities operation

### Synopsis

Execute the get-agent-activities operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 agents get-agent-activities [flags]
```

### Options

```
  -h, --help             help for get-agent-activities
      --limit int        limit field
      --offset int       offset field
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

* [equinix fabricv4 agents](equinix_fabricv4_agents.md)	 - Manage agents resources

