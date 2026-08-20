## equinix metalv1 ip-addresses update-ip-address

Update an ip address

### Synopsis

Update details about an ip address

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ip-addresses update-ip-address [flags]
```

### Options

```
      --exclude string                                            exclude field (JSON or string)
  -h, --help                                                      help for update-ip-address
      --id string                                                 IP Address UUID (required)
      --include string                                            include field (JSON or string)
      --ip-assignment-update-input-additional-properties string   ip-assignment-update-input-additional-properties (JSON)
      --ip-assignment-update-input-customdata string              ip-assignment-update-input-customdata (JSON)
      --ip-assignment-update-input-details string                 ip-assignment-update-input-details
      --ip-assignment-update-input-tags string                    ip-assignment-update-input-tags (JSON array)
      --request string                                            JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ip-addresses](equinix_metalv1_ip-addresses.md)	 - Manage ip-addresses resources

