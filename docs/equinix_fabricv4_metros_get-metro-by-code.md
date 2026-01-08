## equinix fabricv4 metros get-metro-by-code

Get Metro by Code

### Synopsis

GET Metros retrieves all Equinix Fabric metros, as well as latency data between each metro location.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 metros get-metro-by-code [flags]
```

### Options

```
  -h, --help                help for get-metro-by-code
      --metro-code string   Metro Code (required)
      --request string      JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 metros](equinix_fabricv4_metros.md)	 - Manage metros resources

