## equinix metalv1 devices find-traffic

Retrieve device traffic

### Synopsis

Returns the total amount of inbound or outbound traffic for a specific device. The default time period is 1 hour. Please note the results capture all network traffic for the server, but not all traffic may come from or be destined to the Internet and may be non-billable. Only Internet bound traffic is charged.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices find-traffic [flags]
```

### Options

```
      --bucket string                            bucket field
      --direction string                         direction field
  -h, --help                                     help for find-traffic
      --id string                                Device UUID (required)
      --interval string                          interval field
      --request string                           JSON payload for additional optional fields not exposed as flags
      --timeframe-additional-properties string   timeframe-additional-properties (JSON)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 devices](equinix_metalv1_devices.md)	 - Manage devices resources

