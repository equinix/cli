## equinix metalv1 ip-addresses find-ip-address-by-id

Retrieve an ip address

### Synopsis

Returns a single ip address if the user has access.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ip-addresses find-ip-address-by-id [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-ip-address-by-id
      --id string        IP Address UUID (required)
      --include string   include field (JSON or string)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ip-addresses](equinix_metalv1_ip-addresses.md)	 - Manage ip-addresses resources

