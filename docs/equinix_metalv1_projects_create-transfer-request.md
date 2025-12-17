## equinix metalv1 projects create-transfer-request

Create a transfer request

### Synopsis

Organization owners can transfer their projects to other organizations. Deprecated

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 projects create-transfer-request [flags]
```

### Options

```
  -h, --help                                                   help for create-transfer-request
      --id string                                              UUID of the project to be transferred (required)
      --include string                                         include field (JSON or string)
      --request string                                         JSON payload for additional optional fields not exposed as flags
      --transfer-request-input-additional-properties string    transfer-request-input-additional-properties (JSON)
      --transfer-request-input-target_organization_id string   transfer-request-input-target_organization_id
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 projects](equinix_metalv1_projects.md)	 - Manage projects resources

