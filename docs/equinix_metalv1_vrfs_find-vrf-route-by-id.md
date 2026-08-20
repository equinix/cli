## equinix metalv1 vrfs find-vrf-route-by-id

Retrieve a VRF Route

### Synopsis

Returns a single VRF Route resource

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 vrfs find-vrf-route-by-id [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-vrf-route-by-id
      --id string        VRF Route UUID (required)
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

* [equinix metalv1 vrfs](equinix_metalv1_vrfs.md)	 - Manage vrfs resources

