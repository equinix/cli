## equinix fabricv4 ports add-to-lag

Add to Lag

### Synopsis

Add Physical Ports to Virtual Port.<font color="red"> <sup color='red'>Preview</sup></font>

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 ports add-to-lag [flags]
```

### Options

```
      --bulk-physical-port-additional-properties string   bulk-physical-port-additional-properties (JSON)
      --bulk-physical-port-data string                    add physical ports to virtual port (JSON array)
  -h, --help                                              help for add-to-lag
      --port-id string                                    Port UUID (required)
      --request string                                    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

