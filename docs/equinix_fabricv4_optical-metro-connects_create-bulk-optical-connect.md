## equinix fabricv4 optical-metro-connects create-bulk-optical-connect

Create Dual Diverse Optical Metro Connect Service

### Synopsis

Create a dual diverse pair of circuits on separate optical paths.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 optical-metro-connects create-bulk-optical-connect [flags]
```

### Options

```
      --bulk-optical-connect-request-additional-properties string   bulk-optical-connect-request-additional-properties (JSON)
      --bulk-optical-connect-request-data string                    The two connections forming the diverse pair — one PRIMARY and one             SECONDARY. (JSON array)
  -h, --help                                                        help for create-bulk-optical-connect
      --request string                                              JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 optical-metro-connects](equinix_fabricv4_optical-metro-connects.md)	 - Manage optical-metro-connects resources

