## equinix metalv1 i-p-addresses request-i-p-reservation

Execute request-i-p-reservation operation

### Synopsis

Execute the request-i-p-reservation operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 i-p-addresses request-i-p-reservation [flags]
```

### Options

```
      --exclude string                                                                                 exclude field (JSON or string)
  -h, --help                                                                                           help for request-i-p-reservation
      --include string                                                                                 include field (JSON or string)
      --param-1 string                                                                                 param-1 (required)
      --request string                                                                                 JSON payload for additional optional fields not exposed as flags
      --request-i-p-reservation-request-i-p-reservation-request-input-additional-properties string     request-i-p-reservation-request-i-p-reservation-request-input-additional-properties (JSON)
      --request-i-p-reservation-request-i-p-reservation-request-input-comments string                  request-i-p-reservation-request-i-p-reservation-request-input-comments
      --request-i-p-reservation-request-i-p-reservation-request-input-customdata string                request-i-p-reservation-request-i-p-reservation-request-input-customdata (JSON)
      --request-i-p-reservation-request-i-p-reservation-request-input-details string                   request-i-p-reservation-request-i-p-reservation-request-input-details
      --request-i-p-reservation-request-i-p-reservation-request-input-facility string                  request-i-p-reservation-request-i-p-reservation-request-input-facility
      --request-i-p-reservation-request-i-p-reservation-request-input-fail_on_approval_required        request-i-p-reservation-request-i-p-reservation-request-input-fail_on_approval_required
      --request-i-p-reservation-request-i-p-reservation-request-input-metro string                     The code of the metro you are requesting the IP reservation in.
      --request-i-p-reservation-request-i-p-reservation-request-input-quantity int                     request-i-p-reservation-request-i-p-reservation-request-input-quantity
      --request-i-p-reservation-request-i-p-reservation-request-input-tags string                      request-i-p-reservation-request-i-p-reservation-request-input-tags (JSON array)
      --request-i-p-reservation-request-i-p-reservation-request-input-type string                      request-i-p-reservation-request-i-p-reservation-request-input-type
      --request-i-p-reservation-request-vrf-ip-reservation-create-input-additional-properties string   request-i-p-reservation-request-vrf-ip-reservation-create-input-additional-properties (JSON)
      --request-i-p-reservation-request-vrf-ip-reservation-create-input-cidr int                       The size of the VRF IP Reservation's subnet. The following subnet sizes are supported: - IPv4: between 22 - 29 inclusive - IPv6: exactly 64
      --request-i-p-reservation-request-vrf-ip-reservation-create-input-customdata string              request-i-p-reservation-request-vrf-ip-reservation-create-input-customdata (JSON)
      --request-i-p-reservation-request-vrf-ip-reservation-create-input-details string                 request-i-p-reservation-request-vrf-ip-reservation-create-input-details
      --request-i-p-reservation-request-vrf-ip-reservation-create-input-network string                 The starting address for this VRF IP Reservation's subnet. Both IPv4 and IPv6 are supported.
      --request-i-p-reservation-request-vrf-ip-reservation-create-input-tags string                    request-i-p-reservation-request-vrf-ip-reservation-create-input-tags (JSON array)
      --request-i-p-reservation-request-vrf-ip-reservation-create-input-type string                    Must be set to 'vrf'
      --request-i-p-reservation-request-vrf-ip-reservation-create-input-vrf_id string                  The ID of the VRF in which this VRF IP Reservation is created. The VRF must have an existing IP Range that contains the requested subnet. This field may be aliased as just 'vrf'.
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 i-p-addresses](equinix_metalv1_i-p-addresses.md)	 - Manage i-p-addresses resources

