## equinix metalv1 ssh-keys find-device-ssh-keys

Retrieve a device's ssh keys

### Synopsis

Returns a collection of the device's ssh keys.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ssh-keys find-device-ssh-keys [flags]
```

### Options

```
  -h, --help                   help for find-device-ssh-keys
      --id string              Project UUID (required)
      --include string         include field (JSON or string)
      --request string         JSON payload for additional optional fields not exposed as flags
      --search-string string   search-string field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ssh-keys](equinix_metalv1_ssh-keys.md)	 - Manage ssh-keys resources

