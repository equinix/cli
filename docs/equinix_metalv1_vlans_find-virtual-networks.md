## equinix metalv1 vlans find-virtual-networks

Retrieve all virtual networks

### Synopsis

Provides a list of virtual networks for a single project.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 vlans find-virtual-networks [flags]
```

### Options

```
      --exclude string    exclude field (JSON or string)
      --facility string   facility field
  -h, --help              help for find-virtual-networks
      --id string         Project UUID (required)
      --include string    include field (JSON or string)
      --metro string      metro field
      --request string    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 vlans](equinix_metalv1_vlans.md)	 - Manage vlans resources

