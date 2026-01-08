## equinix metalv1 metal-gateways create-metal-gateway

Create a metal gateway

### Synopsis

Create a metal gateway in a project

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 metal-gateways create-metal-gateway [flags]
```

### Options

```
      --create-metal-gateway-request-metal-gateway-create-input-additional-properties string       create-metal-gateway-request-metal-gateway-create-input-additional-properties (JSON)
      --create-metal-gateway-request-metal-gateway-create-input-ip_reservation_id string           The UUID of an IP reservation that belongs to the same project as where the metal gateway will be created in. This field is required unless the private IPv4 subnet size is specified.
      --create-metal-gateway-request-metal-gateway-create-input-private_ipv4_subnet_size int       The subnet size (8, 16, 32, 64, or 128) of the private IPv4 reservation that will be created for the metal gateway. This field is required unless a project IP reservation was specified.           Please keep in mind that the number of private metal gateway ranges are limited per project. If you would like to increase the limit per project, please contact support for assistance.
      --create-metal-gateway-request-metal-gateway-create-input-virtual_network_id string          The UUID of a metro virtual network that belongs to the same project as where the metal gateway will be created in.
      --create-metal-gateway-request-vrf-metal-gateway-create-input-additional-properties string   create-metal-gateway-request-vrf-metal-gateway-create-input-additional-properties (JSON)
      --create-metal-gateway-request-vrf-metal-gateway-create-input-ip_reservation_id string       The UUID an a VRF IP Reservation that belongs to the same project as the one in which the Metal Gateway is to be created. Additionally, the VRF IP Reservation and the Virtual Network must reside in the same Metro.
      --create-metal-gateway-request-vrf-metal-gateway-create-input-virtual_network_id string      The UUID of a Metro Virtual Network that belongs to the same project as the one in which the Metal Gateway is to be created. Additionally, the Virtual Network and the VRF IP Reservation must reside in the same metro. In the case of the IP reservation being an IPv6 based VRF IP Reservation, the Virtual Network must not already have an associated IPv6 based VRF IP Reservation. There can be exactly one IPv6 based VRF IP Reservation associated to a Virtual Network.
      --exclude string                                                                             exclude field (JSON or string)
  -h, --help                                                                                       help for create-metal-gateway
      --include string                                                                             include field (JSON or string)
      --page int                                                                                   page field
      --per-page int                                                                               per-page field
      --project-id string                                                                          Project UUID (required)
      --request string                                                                             JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 metal-gateways](equinix_metalv1_metal-gateways.md)	 - Manage metal-gateways resources

