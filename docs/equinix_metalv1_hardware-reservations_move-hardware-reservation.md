## equinix metalv1 hardware-reservations move-hardware-reservation

Move a hardware reservation

### Synopsis

Move a hardware reservation to another project

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 hardware-reservations move-hardware-reservation [flags]
```

### Options

```
      --exclude string                                                   exclude field (JSON or string)
  -h, --help                                                             help for move-hardware-reservation
      --id string                                                        Hardware Reservation UUID (required)
      --include string                                                   include field (JSON or string)
      --move-hardware-reservation-request-additional-properties string   move-hardware-reservation-request-additional-properties (JSON)
      --move-hardware-reservation-request-project_id string              move-hardware-reservation-request-project_id
      --request string                                                   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 hardware-reservations](equinix_metalv1_hardware-reservations.md)	 - Manage hardware-reservations resources

