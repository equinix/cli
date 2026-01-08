## equinix metalv1 v-l-a-ns create-virtual-network

Execute create-virtual-network operation

### Synopsis

Execute the create-virtual-network operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 v-l-a-ns create-virtual-network [flags]
```

### Options

```
      --exclude string                                              exclude field (JSON or string)
  -h, --help                                                        help for create-virtual-network
      --include string                                              include field (JSON or string)
      --network-id string                                           network-id (required)
      --request string                                              JSON payload for additional optional fields not exposed as flags
      --virtual-network-create-input-additional-properties string   virtual-network-create-input-additional-properties (JSON)
      --virtual-network-create-input-description string             virtual-network-create-input-description
      --virtual-network-create-input-facility string                The UUID (or facility code) for the Facility in which to create this Virtual network.
                                                                    Deprecated
      --virtual-network-create-input-metro string                   The UUID (or metro code) for the Metro in which to create this Virtual Network.
      --virtual-network-create-input-tags string                    virtual-network-create-input-tags (JSON array)
      --virtual-network-create-input-vxlan int                      VLAN ID between 2-3999. Must be unique for the project within the Metro in which this Virtual Network is being created. If no value is specified, the next-available VLAN ID in the range 1000-1999 will be automatically selected.
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 v-l-a-ns](equinix_metalv1_v-l-a-ns.md)	 - Manage v-l-a-ns resources

