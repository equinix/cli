## equinix fabricv4 ports get-ports

Get All Ports

### Synopsis

Get All Ports returns details of all assigned and available ports for the specified user credentials. The metro attribute in the response shows the origin of the proposed connection.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 ports get-ports [flags]
```

### Options

```
  -h, --help             help for get-ports
      --name string      name field
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

