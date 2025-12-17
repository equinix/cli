## equinix metalv1 hardware-reservations activate-hardware-reservation

Activate a spare hardware reservation

### Synopsis

Activate a spare hardware reservation

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 hardware-reservations activate-hardware-reservation [flags]
```

### Options

```
      --activate-hardware-reservation-request-additional-properties string   activate-hardware-reservation-request-additional-properties (JSON)
      --activate-hardware-reservation-request-description string             activate-hardware-reservation-request-description
      --exclude string                                                       exclude field (JSON or string)
  -h, --help                                                                 help for activate-hardware-reservation
      --id string                                                            Hardware Reservation UUID (required)
      --include string                                                       include field (JSON or string)
      --request string                                                       JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 hardware-reservations](equinix_metalv1_hardware-reservations.md)	 - Manage hardware-reservations resources

