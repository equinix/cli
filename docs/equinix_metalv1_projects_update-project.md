## equinix metalv1 projects update-project

Update the project

### Synopsis

Updates the project.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 projects update-project [flags]
```

### Options

```
      --exclude string                                      exclude field (JSON or string)
  -h, --help                                                help for update-project
      --id string                                           Project UUID (required)
      --include string                                      include field (JSON or string)
      --project-update-input-additional-properties string   project-update-input-additional-properties (JSON)
      --project-update-input-backend_transfer_enabled       project-update-input-backend_transfer_enabled
      --project-update-input-customdata string              project-update-input-customdata (JSON)
      --project-update-input-name string                    The name of the project. Cannot contain characters encoded in greater than 3 bytes such as emojis.
      --project-update-input-payment_method_id string       project-update-input-payment_method_id
      --project-update-input-tags string                    project-update-input-tags (JSON array)
      --request string                                      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 projects](equinix_metalv1_projects.md)	 - Manage projects resources

