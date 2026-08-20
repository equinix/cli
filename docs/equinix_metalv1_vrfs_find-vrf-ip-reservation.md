## equinix metalv1 vrfs find-vrf-ip-reservation

Retrieve the Specified VRF IP Reservation

### Synopsis

Returns the specified IP Reservation for the VRF.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 vrfs find-vrf-ip-reservation [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-vrf-ip-reservation
      --id string        IP UUID (required)
      --include string   include field (JSON or string)
      --request string   JSON payload for additional optional fields not exposed as flags
      --vrf-id string    VRF UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 vrfs](equinix_metalv1_vrfs.md)	 - Manage vrfs resources

