## equinix metalv1 licenses update-license

Update the license

### Synopsis

Updates the license.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 licenses update-license [flags]
```

### Options

```
      --exclude string                                      exclude field (JSON or string)
  -h, --help                                                help for update-license
      --id string                                           License UUID (required)
      --include string                                      include field (JSON or string)
      --license-update-input-additional-properties string   license-update-input-additional-properties (JSON)
      --license-update-input-description string             license-update-input-description
      --license-update-input-size float                     license-update-input-size
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

