## equinix metalv1 self-service-reservations create-self-service-reservation

Create a reservation

### Synopsis

Creates a reservation.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 self-service-reservations create-self-service-reservation [flags]
```

### Options

```
      --create-self-service-reservation-request-additional-properties string          create-self-service-reservation-request-additional-properties (JSON)
      --create-self-service-reservation-request-item string                           create-self-service-reservation-request-item (JSON array)
      --create-self-service-reservation-request-notes string                          create-self-service-reservation-request-notes
      --create-self-service-reservation-request-period-additional-properties string   create-self-service-reservation-request-period-additional-properties (JSON)
      --create-self-service-reservation-request-period-count int                      create-self-service-reservation-request-period-count
      --create-self-service-reservation-request-period-unit string                    create-self-service-reservation-request-period-unit
  -h, --help                                                                          help for create-self-service-reservation
      --project-id string                                                             Project UUID (required)
      --request string                                                                JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 self-service-reservations](equinix_metalv1_self-service-reservations.md)	 - Manage self-service-reservations resources

