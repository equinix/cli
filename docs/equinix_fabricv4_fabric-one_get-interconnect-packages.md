## equinix fabricv4 fabric-one get-interconnect-packages

Get All Interconnect Packages

### Synopsis

Get All Interconnect Packages returns details of all available interconnect packages for the specified user credentials. <font color="red"> <sup color='red'>Beta</sup></font>

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 fabric-one get-interconnect-packages [flags]
```

### Options

```
  -h, --help             help for get-interconnect-packages
      --limit int        limit field
      --offset int       offset field
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 fabric-one](equinix_fabricv4_fabric-one.md)	 - Manage fabric-one resources

