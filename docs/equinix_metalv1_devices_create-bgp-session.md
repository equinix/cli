## equinix metalv1 devices create-bgp-session

Create a BGP session

### Synopsis

Creates a BGP session.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices create-bgp-session [flags]
```

### Options

```
      --b-g-p-session-input-additional-properties string   b-g-p-session-input-additional-properties (JSON)
      --b-g-p-session-input-address_family string          b-g-p-session-input-address_family
      --b-g-p-session-input-default_route                  Set the default route policy.
  -h, --help                                               help for create-bgp-session
      --id string                                          Device UUID (required)
      --include string                                     include field (JSON or string)
      --request string                                     JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 devices](equinix_metalv1_devices.md)	 - Manage devices resources

