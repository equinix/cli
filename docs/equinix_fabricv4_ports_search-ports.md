## equinix fabricv4 ports search-ports

Search ports

### Synopsis

The API provides capability to get list of user's virtual ports using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 ports search-ports [flags]
```

### Options

```
  -h, --help             help for search-ports
      --request string   JSON payload for request body. Available fields: port-v4-search-request (PortV4SearchRequest)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

