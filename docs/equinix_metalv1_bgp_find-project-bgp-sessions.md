## equinix metalv1 bgp find-project-bgp-sessions

Retrieve all BGP sessions for project

### Synopsis

Provides a listing of available BGP sessions for the project.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 bgp find-project-bgp-sessions [flags]
```

### Options

```
  -h, --help             help for find-project-bgp-sessions
      --id string        Project UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 bgp](equinix_metalv1_bgp.md)	 - Manage bgp resources

