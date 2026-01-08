## equinix metalv1 interconnections update-interconnection

Update interconnection

### Synopsis

Update the details of a interconnection

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 interconnections update-interconnection [flags]
```

### Options

```
      --connection-id string                                        Interconnection UUID (required)
      --exclude string                                              exclude field (JSON or string)
  -h, --help                                                        help for update-interconnection
      --include string                                              include field (JSON or string)
      --interconnection-update-input-additional-properties string   interconnection-update-input-additional-properties (JSON)
      --interconnection-update-input-contact_email string           interconnection-update-input-contact_email
      --interconnection-update-input-description string             interconnection-update-input-description
      --interconnection-update-input-mode string                    interconnection-update-input-mode
      --interconnection-update-input-name string                    interconnection-update-input-name
      --interconnection-update-input-tags string                    interconnection-update-input-tags (JSON array)
      --request string                                              JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 interconnections](equinix_metalv1_interconnections.md)	 - Manage interconnections resources

