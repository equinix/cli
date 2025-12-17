## equinix metalv1 ports convert-layer3

Convert to Layer 3

### Synopsis

Converts a bond port to Layer 3. VLANs must first be unassigned.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ports convert-layer3 [flags]
```

### Options

```
  -h, --help                                                     help for convert-layer3
      --id string                                                Port UUID (required)
      --include string                                           include field (JSON or string)
      --port-convert-layer3-input-additional-properties string   port-convert-layer3-input-additional-properties (JSON)
      --port-convert-layer3-input-request_ips string             port-convert-layer3-input-request_ips (JSON array)
      --request string                                           JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ports](equinix_metalv1_ports.md)	 - Manage ports resources

