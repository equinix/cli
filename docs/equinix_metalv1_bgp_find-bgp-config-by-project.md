## equinix metalv1 bgp find-bgp-config-by-project

Retrieve a bgp config

### Synopsis

Returns a bgp config

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 bgp find-bgp-config-by-project [flags]
```

### Options

```
  -h, --help             help for find-bgp-config-by-project
      --id string        Project UUID (required)
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

* [equinix metalv1 bgp](equinix_metalv1_bgp.md)	 - Manage bgp resources

