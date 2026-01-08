## equinix metalv1 v-r-fs create-vrf

Execute create-vrf operation

### Synopsis

Execute the create-vrf operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 v-r-fs create-vrf [flags]
```

### Options

```
      --exclude string                                            exclude field (JSON or string)
  -h, --help                                                      help for create-vrf
      --include string                                            include field (JSON or string)
      --param-1 string                                            param-1 (required)
      --request string                                            JSON payload for additional optional fields not exposed as flags
      --vrf-create-input-additional-properties string             vrf-create-input-additional-properties (JSON)
      --vrf-create-input-bgp_dynamic_neighbors_bfd_enabled        Toggle BFD on dynamic bgp neighbors sessions
      --vrf-create-input-bgp_dynamic_neighbors_enabled            Toggle to enable the dynamic bgp neighbors feature on the VRF
      --vrf-create-input-bgp_dynamic_neighbors_export_route_map   Toggle to export the VRF route-map to the dynamic bgp neighbors
      --vrf-create-input-description string                       vrf-create-input-description
      --vrf-create-input-ip_ranges string                         A list of CIDR network addresses. Like [\"10.0.0.0/16\", \"2001:d78::/59\"]. IPv4 blocks must be between /8 and /29 in size. IPv6 blocks must be between /59 and /64. A VRF\\'s IP ranges must be defined in order to create VRF IP Reservations, which can then be used for Metal Gateways or Virtual Circuits. (JSON array)
      --vrf-create-input-local_asn int                            vrf-create-input-local_asn
      --vrf-create-input-metro string                             The UUID (or metro code) for the Metro in which to create this VRF.
      --vrf-create-input-name string                              vrf-create-input-name
      --vrf-create-input-tags string                              vrf-create-input-tags (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 v-r-fs](equinix_metalv1_v-r-fs.md)	 - Manage v-r-fs resources

