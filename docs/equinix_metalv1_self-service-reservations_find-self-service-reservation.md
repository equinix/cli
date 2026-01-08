## equinix metalv1 self-service-reservations find-self-service-reservation

Retrieve a reservation

### Synopsis

Returns a reservation

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 self-service-reservations find-self-service-reservation [flags]
```

### Options

```
  -h, --help                help for find-self-service-reservation
      --id string           Reservation short_id (required)
      --project-id string   Project UUID (required)
      --request string      JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 self-service-reservations](equinix_metalv1_self-service-reservations.md)	 - Manage self-service-reservations resources

