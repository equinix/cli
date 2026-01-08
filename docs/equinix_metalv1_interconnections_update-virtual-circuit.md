## equinix metalv1 interconnections update-virtual-circuit

Update a virtual circuit

### Synopsis

Update the details of a virtual circuit.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 interconnections update-virtual-circuit [flags]
```

### Options

```
      --exclude string                                                                                exclude field (JSON or string)
  -h, --help                                                                                          help for update-virtual-circuit
      --id string                                                                                     Virtual Circuit UUID (required)
      --include string                                                                                include field (JSON or string)
      --request string                                                                                JSON payload for additional optional fields not exposed as flags
      --virtual-circuit-update-input-vlan-virtual-circuit-update-input-additional-properties string   virtual-circuit-update-input-vlan-virtual-circuit-update-input-additional-properties (JSON)
      --virtual-circuit-update-input-vlan-virtual-circuit-update-input-description string             virtual-circuit-update-input-vlan-virtual-circuit-update-input-description
      --virtual-circuit-update-input-vlan-virtual-circuit-update-input-name string                    virtual-circuit-update-input-vlan-virtual-circuit-update-input-name
      --virtual-circuit-update-input-vlan-virtual-circuit-update-input-speed string                   Speed can be changed only if it is an interconnection on a Dedicated Port
      --virtual-circuit-update-input-vlan-virtual-circuit-update-input-tags string                    virtual-circuit-update-input-vlan-virtual-circuit-update-input-tags (JSON array)
      --virtual-circuit-update-input-vlan-virtual-circuit-update-input-vnid string                    A Virtual Network record UUID or the VNID of a Metro Virtual Network in your project.
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-additional-properties string    virtual-circuit-update-input-vrf-virtual-circuit-update-input-additional-properties (JSON)
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-customer_ip string              An IPv4 address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used.
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-customer_ipv6 string            An IPv6 address from the subnet IPv6 that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet IPv6 as the Metal IPv6. By default, the last usable IP address in the subnet IPv6 will be used.
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-description string              virtual-circuit-update-input-vrf-virtual-circuit-update-input-description
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-md5 string                      The plaintext BGP peering password shared by neighbors as an MD5 checksum: * must be 10-20 characters long * may not include punctuation * must be a combination of numbers and letters * must contain at least one lowercase, uppercase, and digit character
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-metal_ip string                 An IPv4 address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used.
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-metal_ipv6 string               An IPv6 address from the subnet IPv6 that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IPv6 address in the subnet IPv6 as the Customer IP. By default, the first usable IPv6 address in the subnet IPv6 will be used.
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-name string                     virtual-circuit-update-input-vrf-virtual-circuit-update-input-name
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-peer_asn int                    The peer ASN that will be used with the VRF on the Virtual Circuit.
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-speed string                    Speed can be changed only if it is an interconnection on a Dedicated Port
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-subnet string                   The /30 or /31 IPv4 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP.
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-subnet_ipv6 string              The /126 or /127 IPv6 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IPv6 and Customer IPv6 must be IPs from this subnet. For /126 subnets, the network and broadcast IPs cannot be used as the Metal IPv6 or Customer IPv6. The subnet specified must be contained within an already-defined IP Range for the VRF.
      --virtual-circuit-update-input-vrf-virtual-circuit-update-input-tags string                     virtual-circuit-update-input-vrf-virtual-circuit-update-input-tags (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 interconnections](equinix_metalv1_interconnections.md)	 - Manage interconnections resources

