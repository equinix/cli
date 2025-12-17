## equinix metalv1 batches find-batches-by-project

Retrieve all batches by project

### Synopsis

Returns all batches for the given project

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 batches find-batches-by-project [flags]
```

### Options

```
  -h, --help             help for find-batches-by-project
      --id string        Project UUID (required)
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

* [equinix metalv1 batches](equinix_metalv1_batches.md)	 - Manage batches resources

