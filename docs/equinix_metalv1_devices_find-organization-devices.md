## equinix metalv1 devices find-organization-devices

Retrieve all devices of an organization

### Synopsis

Provides a collection of devices for a given organization.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices find-organization-devices [flags]
```

### Options

```
      --categories string      categories field (JSON or string)
      --exclude string         exclude field (JSON or string)
      --facility string        facility field
      --has-termination-time   has-termination-time field
  -h, --help                   help for find-organization-devices
      --hostname string        hostname field
      --id string              Organization UUID (required)
      --include string         include field (JSON or string)
      --mac-address string     mac-address field
      --page int               page field
      --per-page int           per-page field
      --request string         JSON payload for additional optional fields not exposed as flags
      --reserved               reserved field
      --search string          search field
      --tag string             tag field
      --type_ string           type_ field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 devices](equinix_metalv1_devices.md)	 - Manage devices resources

