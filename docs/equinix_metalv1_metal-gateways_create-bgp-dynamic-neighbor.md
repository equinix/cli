## equinix metalv1 metal-gateways create-bgp-dynamic-neighbor

Create a VRF BGP Dynamic Neighbor range

### Synopsis

Create a VRF BGP Dynamic Neighbor range. BGP Dynamic Neighbor records are limited to 2 per Virtual Network.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 metal-gateways create-bgp-dynamic-neighbor [flags]
```

### Options

```
      --bgp-dynamic-neighbor-create-input-additional-properties string   bgp-dynamic-neighbor-create-input-additional-properties (JSON)
      --bgp-dynamic-neighbor-create-input-bgp_neighbor_asn int           The ASN of the dynamic BGP neighbor
      --bgp-dynamic-neighbor-create-input-bgp_neighbor_range string      Network range of the dynamic BGP neighbor in CIDR format
      --bgp-dynamic-neighbor-create-input-tags string                    bgp-dynamic-neighbor-create-input-tags (JSON array)
      --exclude string                                                   exclude field (JSON or string)
  -h, --help                                                             help for create-bgp-dynamic-neighbor
      --id string                                                        Metal Gateway UUID (required)
      --include string                                                   include field (JSON or string)
      --request string                                                   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 metal-gateways](equinix_metalv1_metal-gateways.md)	 - Manage metal-gateways resources

