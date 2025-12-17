## equinix metalv1 organizations create-organization-project

Create a project for the organization

### Synopsis

Creates a new project for the organization

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 organizations create-organization-project [flags]
```

### Options

```
      --exclude string                                      exclude field (JSON or string)
  -h, --help                                                help for create-organization-project
      --id string                                           Organization UUID (required)
      --include string                                      include field (JSON or string)
      --project-create-input-additional-properties string   project-create-input-additional-properties (JSON)
      --project-create-input-customdata string              project-create-input-customdata (JSON)
      --project-create-input-name string                    The name of the project. Cannot contain characters encoded in greater than 3 bytes such as emojis.
      --project-create-input-payment_method_id string       project-create-input-payment_method_id
      --project-create-input-tags string                    project-create-input-tags (JSON array)
      --project-create-input-type string                    project-create-input-type
      --request string                                      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 organizations](equinix_metalv1_organizations.md)	 - Manage organizations resources

