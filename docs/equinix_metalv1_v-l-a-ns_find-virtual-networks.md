## equinix metalv1 v-l-a-ns find-virtual-networks

Execute find-virtual-networks operation

### Synopsis

Execute the find-virtual-networks operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix metalv1 v-l-a-ns find-virtual-networks [flags]
```

### Options

```
      --exclude string      exclude field (JSON or string)
      --facility string     facility field
  -h, --help                help for find-virtual-networks
      --include string      include field (JSON or string)
      --metro string        metro field
      --network-id string   network-id (required)
      --request string      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 v-l-a-ns](equinix_metalv1_v-l-a-ns.md)	 - Manage v-l-a-ns resources

