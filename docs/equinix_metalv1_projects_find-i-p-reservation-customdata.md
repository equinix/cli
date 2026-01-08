## equinix metalv1 projects find-i-p-reservation-customdata

Retrieve the custom metadata of an IP Reservation

### Synopsis

Provides the custom metadata stored for this IP Reservation in json format

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 projects find-i-p-reservation-customdata [flags]
```

### Options

```
  -h, --help                help for find-i-p-reservation-customdata
      --id string           Ip Reservation UUID (required)
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

* [equinix metalv1 projects](equinix_metalv1_projects.md)	 - Manage projects resources

