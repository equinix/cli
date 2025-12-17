## equinix metalv1 v-r-fs update-vrf

Execute update-vrf operation

### Synopsis

Execute the update-vrf operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 v-r-fs update-vrf [flags]
```

### Options

```
      --exclude string                                            exclude field (JSON or string)
  -h, --help                                                      help for update-vrf
      --include string                                            include field (JSON or string)
      --param-1 string                                            param-1 (required)
      --request string                                            JSON payload for additional optional fields not exposed as flags
      --vrf-update-input-additional-properties string             vrf-update-input-additional-properties (JSON)
      --vrf-update-input-bgp_dynamic_neighbors_bfd_enabled        Toggle BFD on dynamic bgp neighbors sessions
      --vrf-update-input-bgp_dynamic_neighbors_enabled            Toggle to enable the dynamic bgp neighbors feature on the VRF
      --vrf-update-input-bgp_dynamic_neighbors_export_route_map   Toggle to export the VRF route-map to the dynamic bgp neighbors
      --vrf-update-input-description string                       vrf-update-input-description
      --vrf-update-input-ip_ranges ip_ranges                      A list of CIDR network addresses. Like [\"10.0.0.0/16\", \"2001:d78::/59\"]. IPv4 blocks must be between /8 and /29 in size. IPv6 blocks must be between /59 and /64. A VRF\\'s IP ranges must be defined in order to create VRF IP Reservations, which can then be used for Metal Gateways or Virtual Circuits. Adding a new CIDR address to the list will result in the creation of a new IP Range for this VRF. Removal of an existing CIDR address from the list will result in the deletion of an existing IP Range for this VRF. Deleting an IP Range will result in the deletion of any VRF IP Reservations contained within the IP Range, as well as the VRF IP Reservation\\'s associated Metal Gateways or Virtual Circuits. If you do not wish to add or remove IP Ranges, either include the full existing list of IP Ranges in the update request, or do not specify the ip_ranges field in the update request. Specifying a value of `[]` will remove all existing IP Ranges from the VRF. (JSON array)
      --vrf-update-input-local_asn local_asn                      The new local_asn value for the VRF. This field cannot be updated when there are active Interconnection Virtual Circuits associated to the VRF, or if any of the VLANs of the VRF's metal gateway has been assigned on an instance.
      --vrf-update-input-name string                              vrf-update-input-name
      --vrf-update-input-tags string                              vrf-update-input-tags (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 v-r-fs](equinix_metalv1_v-r-fs.md)	 - Manage v-r-fs resources

