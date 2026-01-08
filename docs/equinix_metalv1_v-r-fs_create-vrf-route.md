## equinix metalv1 v-r-fs create-vrf-route

Execute create-vrf-route operation

### Synopsis

Execute the create-vrf-route operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 v-r-fs create-vrf-route [flags]
```

### Options

```
      --exclude string                                        exclude field (JSON or string)
  -h, --help                                                  help for create-vrf-route
      --include string                                        include field (JSON or string)
      --param-1 string                                        param-1 (required)
      --request string                                        JSON payload for additional optional fields not exposed as flags
      --vrf-route-create-input-additional-properties string   vrf-route-create-input-additional-properties (JSON)
      --vrf-route-create-input-next_hop string                The IPv4 address within the VRF of the host that will handle this route
      --vrf-route-create-input-prefix string                  The IPv4 prefix for the route, in CIDR-style notation. For a static default route, this will always be \"0.0.0.0/0\"
      --vrf-route-create-input-tags string                    vrf-route-create-input-tags (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 v-r-fs](equinix_metalv1_v-r-fs.md)	 - Manage v-r-fs resources

