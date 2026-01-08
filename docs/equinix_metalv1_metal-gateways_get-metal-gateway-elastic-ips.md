## equinix metalv1 metal-gateways get-metal-gateway-elastic-ips

List Metal Gateway Elastic IPs

### Synopsis

Returns the list of Elastic IPs assigned to this Metal Gateway

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 metal-gateways get-metal-gateway-elastic-ips [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for get-metal-gateway-elastic-ips
      --id string        Metal Gateway UUID (required)
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

* [equinix metalv1 metal-gateways](equinix_metalv1_metal-gateways.md)	 - Manage metal-gateways resources

