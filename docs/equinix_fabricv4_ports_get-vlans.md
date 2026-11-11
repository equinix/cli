## equinix fabricv4 ports get-vlans

Get Vlans

### Synopsis

The API provides capability to retrieve Vlans for a Port.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 ports get-vlans [flags]
```

### Options

```
  -h, --help[=true|false]   help for get-vlans
      --limit int           limit field
      --offset int          offset field
      --port-uuid string    Port UUID (required)
      --request string      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

