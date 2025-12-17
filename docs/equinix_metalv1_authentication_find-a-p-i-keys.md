## equinix metalv1 authentication find-a-p-i-keys

Retrieve all user API keys

### Synopsis

Returns all API keys for the current user.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 authentication find-a-p-i-keys [flags]
```

### Options

```
  -h, --help             help for find-a-p-i-keys
      --include string   include field (JSON or string)
      --request string   JSON payload for additional optional fields not exposed as flags
      --search string    search field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 authentication](equinix_metalv1_authentication.md)	 - Manage authentication resources

