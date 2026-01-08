## equinix metalv1 users find-user-customdata

Retrieve the custom metadata of a user

### Synopsis

Provides the custom metadata stored for this user in json format

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 users find-user-customdata [flags]
```

### Options

```
  -h, --help             help for find-user-customdata
      --id string        User UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 users](equinix_metalv1_users.md)	 - Manage users resources

