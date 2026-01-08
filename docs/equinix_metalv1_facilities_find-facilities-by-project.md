## equinix metalv1 facilities find-facilities-by-project

Retrieve all facilities visible by the project

### Synopsis

Returns a listing of available datacenters for the given project Deprecated

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 facilities find-facilities-by-project [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-facilities-by-project
      --id string        Project UUID (required)
      --include string   include field (JSON or string)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 facilities](equinix_metalv1_facilities.md)	 - Manage facilities resources

