## equinix metalv1 metal-gateways create-metal-gateway-elastic-ip

Create a Metal Gateway Elastic IP

### Synopsis

Create a new Elastic IP on this Metal Gateway. Assign an IPv4 range as an elastic IP to the Metal Gateway, with a specified next-hop address contained within the Metal Gateway. Notice: Elastic IPs on Metal Gateways are a test feature currently under active development, and only available to certain users. Please contact Customer Success for more information.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 metal-gateways create-metal-gateway-elastic-ip [flags]
```

### Options

```
      --exclude string                                                       exclude field (JSON or string)
  -h, --help                                                                 help for create-metal-gateway-elastic-ip
      --id string                                                            Metal Gateway UUID (required)
      --include string                                                       include field (JSON or string)
      --metal-gateway-elastic-ip-create-input-additional-properties string   metal-gateway-elastic-ip-create-input-additional-properties (JSON)
      --metal-gateway-elastic-ip-create-input-address string                 An IP address (or IP Address range) contained within one of the project's IP Reservations
      --metal-gateway-elastic-ip-create-input-customdata string              Optional User-defined JSON object value. (JSON)
      --metal-gateway-elastic-ip-create-input-next_hop string                An IP address contained within the Metal Gateways' IP Reservation range.
      --metal-gateway-elastic-ip-create-input-tags string                    Optional list of User-defined tags. Can be used by users to provide additional details or context regarding the purpose or usage of this resource. (JSON array)
      --request string                                                       JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 metal-gateways](equinix_metalv1_metal-gateways.md)	 - Manage metal-gateways resources

