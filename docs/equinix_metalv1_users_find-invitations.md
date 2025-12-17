## equinix metalv1 users find-invitations

Retrieve current user invitations

### Synopsis

Returns all invitations in current user.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 users find-invitations [flags]
```

### Options

```
  -h, --help             help for find-invitations
      --include string   include field (JSON or string)
      --page int         page field
      --per-page int     per-page field
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 users](equinix_metalv1_users.md)	 - Manage users resources

