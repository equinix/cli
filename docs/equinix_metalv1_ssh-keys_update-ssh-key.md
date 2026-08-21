## equinix metalv1 ssh-keys update-ssh-key

Update the ssh key

### Synopsis

Updates the ssh key.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ssh-keys update-ssh-key [flags]
```

### Options

```
  -h, --help                                         help for update-ssh-key
      --id string                                    SSH Key UUID (required)
      --include string                               include field (JSON or string)
      --request string                               JSON payload for additional optional fields not exposed as flags
      --ssh-key-input-additional-properties string   ssh-key-input-additional-properties (JSON)
      --ssh-key-input-key string                     ssh-key-input-key
      --ssh-key-input-label string                   ssh-key-input-label
      --ssh-key-input-tags string                    ssh-key-input-tags (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ssh-keys](equinix_metalv1_ssh-keys.md)	 - Manage ssh-keys resources

