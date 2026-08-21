## equinix metalv1 ssh-keys find-ssh-keys

Retrieve all ssh keys

### Synopsis

Returns a collection of the user’s ssh keys.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ssh-keys find-ssh-keys [flags]
```

### Options

```
  -h, --help             help for find-ssh-keys
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

* [equinix metalv1 ssh-keys](equinix_metalv1_ssh-keys.md)	 - Manage ssh-keys resources

