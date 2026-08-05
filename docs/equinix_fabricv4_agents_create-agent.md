## equinix fabricv4 agents create-agent

Execute create-agent operation

### Synopsis

Execute the create-agent operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 agents create-agent [flags]
```

### Options

```
      --agent-post-request-additional-properties string                  agent-post-request-additional-properties (JSON)
      --agent-post-request-agent-template-additional-properties string   agent-post-request-agent-template-additional-properties (JSON)
      --agent-post-request-agent-template-uuid string                    agent-post-request-agent-template-uuid
      --agent-post-request-configuration-additional-properties string    agent-post-request-configuration-additional-properties (JSON)
      --agent-post-request-configuration-prompt string                   agent-post-request-configuration-prompt
      --agent-post-request-description string                            agent-post-request-description
      --agent-post-request-enabled                                       agent-post-request-enabled
      --agent-post-request-name string                                   agent-post-request-name
      --agent-post-request-project-additional-properties string          agent-post-request-project-additional-properties (JSON)
      --agent-post-request-project-project-id string                     Subscriber-assigned project ID
      --agent-post-request-type string                                   agent-post-request-type
  -h, --help                                                             help for create-agent
      --request string                                                   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 agents](equinix_fabricv4_agents.md)	 - Manage agents resources

