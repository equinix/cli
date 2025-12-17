## equinix metalv1 batches create-device-batch

Create a devices batch

### Synopsis

Creates new devices in batch and provisions them in our datacenter.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 batches create-device-batch [flags]
```

### Options

```
  -h, --help                                                        help for create-device-batch
      --id string                                                   Project UUID (required)
      --instances-batch-create-input-additional-properties string   instances-batch-create-input-additional-properties (JSON)
      --instances-batch-create-input-batches string                 instances-batch-create-input-batches (JSON array)
      --request string                                              JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 batches](equinix_metalv1_batches.md)	 - Manage batches resources

