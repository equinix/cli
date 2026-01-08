## equinix metalv1 licenses create-license

Create a License

### Synopsis

Creates a new license for the given project

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 licenses create-license [flags]
```

### Options

```
      --exclude string                                      exclude field (JSON or string)
  -h, --help                                                help for create-license
      --id string                                           Project UUID (required)
      --include string                                      include field (JSON or string)
      --license-create-input-additional-properties string   license-create-input-additional-properties (JSON)
      --license-create-input-description string             license-create-input-description
      --license-create-input-licensee_product_id string     license-create-input-licensee_product_id
      --license-create-input-size float                     license-create-input-size
      --request string                                      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 licenses](equinix_metalv1_licenses.md)	 - Manage licenses resources

