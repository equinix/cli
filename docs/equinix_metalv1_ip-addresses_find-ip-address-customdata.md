## equinix metalv1 ip-addresses find-ip-address-customdata

Retrieve the custom metadata of an IP Reservation or IP Assignment

### Synopsis

Provides the custom metadata stored for this IP Reservation or IP Assignment in json format

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ip-addresses find-ip-address-customdata [flags]
```

### Options

```
  -h, --help             help for find-ip-address-customdata
      --id string        Ip Reservation UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ip-addresses](equinix_metalv1_ip-addresses.md)	 - Manage ip-addresses resources

