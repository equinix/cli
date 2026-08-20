## equinix metalv1 ip-addresses find-ip-availabilities

Retrieve all available subnets of a particular reservation

### Synopsis

Provides a list of IP resevations for a single project.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ip-addresses find-ip-availabilities [flags]
```

### Options

```
      --cidr string      cidr field
  -h, --help             help for find-ip-availabilities
      --id string        IP Reservation UUID (required)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ip-addresses](equinix_metalv1_ip-addresses.md)	 - Manage ip-addresses resources

