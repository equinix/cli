## equinix metalv1 users find-users

Retrieve all users

### Synopsis

Returns a list of users that the are accessible to the current user (all users in the current user’s projects, essentially).

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 users find-users [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-users
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

