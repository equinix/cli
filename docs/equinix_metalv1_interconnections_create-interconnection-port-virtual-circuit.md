## equinix metalv1 interconnections create-interconnection-port-virtual-circuit

Create a new Virtual Circuit

### Synopsis

Create a new Virtual Circuit on a Dedicated Port. To create a regular Virtual Circuit, specify a Virtual Network record and an NNI VLAN value. To create a VRF-based Virtual Circuit, specify the VRF ID and subnet, along with the NNI VLAN value.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 interconnections create-interconnection-port-virtual-circuit [flags]
```

### Options

```
      --connection-id string                                                                          UUID of the interconnection (required)
  -h, --help                                                                                          help for create-interconnection-port-virtual-circuit
      --port-id string                                                                                UUID of the interconnection port (required)
      --request string                                                                                JSON payload for additional optional fields not exposed as flags
      --virtual-circuit-create-input-vlan-virtual-circuit-create-input-additional-properties string   virtual-circuit-create-input-vlan-virtual-circuit-create-input-additional-properties (JSON)
      --virtual-circuit-create-input-vlan-virtual-circuit-create-input-description string             virtual-circuit-create-input-vlan-virtual-circuit-create-input-description
      --virtual-circuit-create-input-vlan-virtual-circuit-create-input-name string                    virtual-circuit-create-input-vlan-virtual-circuit-create-input-name
      --virtual-circuit-create-input-vlan-virtual-circuit-create-input-nni_vlan int                   virtual-circuit-create-input-vlan-virtual-circuit-create-input-nni_vlan
      --virtual-circuit-create-input-vlan-virtual-circuit-create-input-project_id string              virtual-circuit-create-input-vlan-virtual-circuit-create-input-project_id
      --virtual-circuit-create-input-vlan-virtual-circuit-create-input-speed string                   speed can be passed as integer number representing bps speed or string (e.g. '52m' or '100g' or '4 gbps')
      --virtual-circuit-create-input-vlan-virtual-circuit-create-input-tags string                    virtual-circuit-create-input-vlan-virtual-circuit-create-input-tags (JSON array)
      --virtual-circuit-create-input-vlan-virtual-circuit-create-input-vnid string                    A Virtual Network record UUID or the VNID of a Metro Virtual Network in your project (sent as integer).
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-additional-properties string    virtual-circuit-create-input-vrf-virtual-circuit-create-input-additional-properties (JSON)
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-customer_ip string              An IPv4 address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used.
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-customer_ipv6 string            An IPv6 address from the subnet IPv6 that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet IPv6 as the Metal IPv6. By default, the last usable IP address in the subnet IPv6 will be used.
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-description string              virtual-circuit-create-input-vrf-virtual-circuit-create-input-description
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-md5 string                      virtual-circuit-create-input-vrf-virtual-circuit-create-input-md5 (JSON)
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-metal_ip string                 An IPv4 address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used.
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-metal_ipv6 string               An IPv6 address from the subnet IPv6 that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IPv6 address in the subnet IPv6 as the Customer IP. By default, the first usable IPv6 address in the subnet IPv6 will be used.
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-name string                     virtual-circuit-create-input-vrf-virtual-circuit-create-input-name
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-nni_vlan int                    virtual-circuit-create-input-vrf-virtual-circuit-create-input-nni_vlan
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-peer_asn int                    The peer ASN that will be used with the VRF on the Virtual Circuit.
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-project_id string               virtual-circuit-create-input-vrf-virtual-circuit-create-input-project_id
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-speed string                    speed can be passed as integer number representing bps speed or string (e.g. '52m' or '100g' or '4 gbps')
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-subnet string                   The /30 or /31 IPv4 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP. The subnet specified must be contained within an already-defined IP Range for the VRF.
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-subnet_ipv6 string              The /126 or /127 IPv6 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IPv6 and Customer IPv6 must be IPs from this subnet. For /126 subnets, the network and broadcast IPs cannot be used as the Metal IPv6 or Customer IPv6. The subnet specified must be contained within an already-defined IP Range for the VRF.
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-tags string                     virtual-circuit-create-input-vrf-virtual-circuit-create-input-tags (JSON array)
      --virtual-circuit-create-input-vrf-virtual-circuit-create-input-vrf string                      The UUID of the VRF that will be associated with the Virtual Circuit.
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 interconnections](equinix_metalv1_interconnections.md)	 - Manage interconnections resources

