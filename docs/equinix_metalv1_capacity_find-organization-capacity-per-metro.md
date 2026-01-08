## equinix metalv1 capacity find-organization-capacity-per-metro

View available hardware plans per Metro for given organization

### Synopsis

Returns a list of metros and plans with their current capacity.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 capacity find-organization-capacity-per-metro [flags]
```

### Options

```
  -h, --help             help for find-organization-capacity-per-metro
      --id string        Organization UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 capacity](equinix_metalv1_capacity.md)	 - Manage capacity resources

