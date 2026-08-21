## equinix metalv1 vlans update-virtual-network

Updates the virtual network

### Synopsis

Updates the virtual network.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 vlans update-virtual-network [flags]
```

### Options

```
      --exclude string                                              exclude field (JSON or string)
  -h, --help                                                        help for update-virtual-network
      --id string                                                   Virtual Network UUID (required)
      --include string                                              include field (JSON or string)
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

* [equinix metalv1 vlans](equinix_metalv1_vlans.md)	 - Manage vlans resources

