## equinix fabricv4 agents patch-agent-by-uuid

Execute patch-agent-by-uuid operation

### Synopsis

Execute the patch-agent-by-uuid operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 agents patch-agent-by-uuid [flags]
```

### Options

```
      --agent-patch-request-additional-properties string   agent-patch-request-additional-properties (JSON)
      --agent-patch-request-op string                      agent-patch-request-op
      --agent-patch-request-path string                    agent-patch-request-path
      --agent-patch-request-value string                   agent-patch-request-value (JSON)
  -h, --help                                               help for patch-agent-by-uuid
      --request string                                     JSON payload for additional optional fields not exposed as flags
      --uuid string                                        uuid (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 agents](equinix_fabricv4_agents.md)	 - Manage agents resources

