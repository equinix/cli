## equinix fabricv4 ports get-port-by-uuid

Get Port by uuid

### Synopsis

Get Port By uuid returns details of assigned and available Equinix Fabric port for the specified user credentials. The metro code attribute in the response shows the origin of the proposed connection.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 ports get-port-by-uuid [flags]
```

### Options

```
  -h, --help             help for get-port-by-uuid
      --port-id string   Port UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

