## equinix metalv1 v-l-a-ns update-virtual-network

Execute update-virtual-network operation

### Synopsis

Execute the update-virtual-network operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 v-l-a-ns update-virtual-network [flags]
```

### Options

```
      --exclude string                                              exclude field (JSON or string)
  -h, --help                                                        help for update-virtual-network
      --include string                                              include field (JSON or string)
      --network-id string                                           network-id (required)
      --request string                                              JSON payload for additional optional fields not exposed as flags
      --virtual-network-update-input-additional-properties string   virtual-network-update-input-additional-properties (JSON)
      --virtual-network-update-input-description string             virtual-network-update-input-description
      --virtual-network-update-input-tags string                    virtual-network-update-input-tags (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 v-l-a-ns](equinix_metalv1_v-l-a-ns.md)	 - Manage v-l-a-ns resources

