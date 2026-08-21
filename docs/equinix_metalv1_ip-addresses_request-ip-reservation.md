## equinix metalv1 ip-addresses request-ip-reservation

Requesting IP reservations

### Synopsis

Request more IP space for a project in order to have additional IP addresses to assign to devices.  If the request is within the max quota, an IP reservation will be created. If the project will exceed its IP quota, a request will be submitted for review, and will return an IP Reservation with a `state` of `pending`. You can automatically have the request fail with HTTP status 422 instead of triggering the review process by providing the `fail_on_approval_required` parameter set to `true` in the request.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ip-addresses request-ip-reservation [flags]
```

### Options

```
      --exclude string                                                                                exclude field (JSON or string)
  -h, --help                                                                                          help for request-ip-reservation
      --id string                                                                                     Project UUID (required)
      --include string                                                                                include field (JSON or string)
      --request string                                                                                JSON payload for additional optional fields not exposed as flags
      --request-ip-reservation-request-ip-reservation-request-input-additional-properties string      request-ip-reservation-request-ip-reservation-request-input-additional-properties (JSON)
      --request-ip-reservation-request-ip-reservation-request-input-comments string                   request-ip-reservation-request-ip-reservation-request-input-comments
      --request-ip-reservation-request-ip-reservation-request-input-customdata string                 request-ip-reservation-request-ip-reservation-request-input-customdata (JSON)
      --request-ip-reservation-request-ip-reservation-request-input-details string                    request-ip-reservation-request-ip-reservation-request-input-details
      --request-ip-reservation-request-ip-reservation-request-input-facility string                   request-ip-reservation-request-ip-reservation-request-input-facility
      --request-ip-reservation-request-ip-reservation-request-input-fail_on_approval_required         request-ip-reservation-request-ip-reservation-request-input-fail_on_approval_required
      --request-ip-reservation-request-ip-reservation-request-input-metro string                      The code of the metro you are requesting the IP reservation in.
      --request-ip-reservation-request-ip-reservation-request-input-quantity int                      request-ip-reservation-request-ip-reservation-request-input-quantity
      --request-ip-reservation-request-ip-reservation-request-input-tags string                       request-ip-reservation-request-ip-reservation-request-input-tags (JSON array)
      --request-ip-reservation-request-ip-reservation-request-input-type string                       request-ip-reservation-request-ip-reservation-request-input-type
      --request-ip-reservation-request-vrf-ip-reservation-create-input-additional-properties string   request-ip-reservation-request-vrf-ip-reservation-create-input-additional-properties (JSON)
      --request-ip-reservation-request-vrf-ip-reservation-create-input-cidr int                       The size of the VRF IP Reservation's subnet. The following subnet sizes are supported: - IPv4: between 22 - 29 inclusive - IPv6: exactly 64
      --request-ip-reservation-request-vrf-ip-reservation-create-input-customdata string              request-ip-reservation-request-vrf-ip-reservation-create-input-customdata (JSON)
      --request-ip-reservation-request-vrf-ip-reservation-create-input-details string                 request-ip-reservation-request-vrf-ip-reservation-create-input-details
      --request-ip-reservation-request-vrf-ip-reservation-create-input-network string                 The starting address for this VRF IP Reservation's subnet. Both IPv4 and IPv6 are supported.
      --request-ip-reservation-request-vrf-ip-reservation-create-input-tags string                    request-ip-reservation-request-vrf-ip-reservation-create-input-tags (JSON array)
      --request-ip-reservation-request-vrf-ip-reservation-create-input-type string                    Must be set to 'vrf'
      --request-ip-reservation-request-vrf-ip-reservation-create-input-vrf_id string                  The ID of the VRF in which this VRF IP Reservation is created. The VRF must have an existing IP Range that contains the requested subnet. This field may be aliased as just 'vrf'.
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ip-addresses](equinix_metalv1_ip-addresses.md)	 - Manage ip-addresses resources

