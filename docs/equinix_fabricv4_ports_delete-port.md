## equinix fabricv4 ports delete-port

Delete a single port

### Synopsis

The API provides capability to delete a single port

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 ports delete-port [flags]
```

### Options

```
      --dry-run[=true|false]   dry-run field
  -h, --help[=true|false]      help for delete-port
      --port-id string         Port UUID (required)
      --request string         JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

