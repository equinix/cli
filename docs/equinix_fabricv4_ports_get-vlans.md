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
  -h, --help               help for get-vlans
      --port-uuid string   Port UUID (required)
      --request string     JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

