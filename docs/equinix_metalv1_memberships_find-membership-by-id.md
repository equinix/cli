## equinix metalv1 memberships find-membership-by-id

Retrieve a membership

### Synopsis

Returns a single membership.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 memberships find-membership-by-id [flags]
```

### Options

```
  -h, --help             help for find-membership-by-id
      --id string        Membership UUID (required)
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

* [equinix metalv1 memberships](equinix_metalv1_memberships.md)	 - Manage memberships resources

