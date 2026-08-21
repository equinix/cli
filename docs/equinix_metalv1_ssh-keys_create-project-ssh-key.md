## equinix metalv1 ssh-keys create-project-ssh-key

Create a ssh key for the given project

### Synopsis

Creates a ssh key.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ssh-keys create-project-ssh-key [flags]
```

### Options

```
  -h, --help                                                help for create-project-ssh-key
      --id string                                           Project UUID (required)
      --include string                                      include field (JSON or string)
      --request string                                      JSON payload for additional optional fields not exposed as flags
      --ssh-key-create-input-additional-properties string   ssh-key-create-input-additional-properties (JSON)
      --ssh-key-create-input-instances_ids string           List of instance UUIDs to associate SSH key with, when empty array is sent all instances belonging       to entity will be included (JSON array)
      --ssh-key-create-input-key string                     ssh-key-create-input-key
      --ssh-key-create-input-label string                   ssh-key-create-input-label
      --ssh-key-create-input-tags string                    ssh-key-create-input-tags (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ssh-keys](equinix_metalv1_ssh-keys.md)	 - Manage ssh-keys resources

