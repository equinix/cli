## equinix metalv1 bgp update-bgp-session

Update the BGP session

### Synopsis

Updates the BGP session by either enabling or disabling the default route functionality.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 bgp update-bgp-session [flags]
```

### Options

```
      --body             body field
  -h, --help             help for update-bgp-session
      --id string        BGP session UUID (required)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 bgp](equinix_metalv1_bgp.md)	 - Manage bgp resources

