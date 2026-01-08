## equinix metalv1 users find-current-user

Retrieve the current user

### Synopsis

Returns the user object for the currently logged-in user.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 users find-current-user [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-current-user
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

* [equinix metalv1 users](equinix_metalv1_users.md)	 - Manage users resources

