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
      --request string   Raw JSON payload for optional request fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

