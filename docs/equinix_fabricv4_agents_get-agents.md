## equinix fabricv4 agents get-agents

Execute get-agents operation

### Synopsis

Execute the get-agents operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 agents get-agents [flags]
```

### Options

```
  -h, --help             help for get-agents
      --limit int        limit field
      --offset int       offset field
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

