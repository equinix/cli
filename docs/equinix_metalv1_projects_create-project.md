## equinix metalv1 projects create-project

Create a project

### Synopsis

Creates a new project for the user's default organization. If the user does not have a default organization, the API will look for a personal organization belonging to the user with the name "{User's Full Name} Projects" to associate the project with. If that organization does not exist a new organization named "{User's Full Name} Projects" will be created and the new project will be tied to that organization.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 projects create-project [flags]
```

### Options

```
      --exclude string                                                exclude field (JSON or string)
  -h, --help                                                          help for create-project
      --include string                                                include field (JSON or string)
      --project-create-from-root-input-additional-properties string   project-create-from-root-input-additional-properties (JSON)
      --project-create-from-root-input-customdata string              project-create-from-root-input-customdata (JSON)
      --project-create-from-root-input-name string                    The name of the project. Cannot contain characters encoded in greater than 3 bytes such as emojis.
      --project-create-from-root-input-organization_id string         project-create-from-root-input-organization_id
      --project-create-from-root-input-payment_method_id string       project-create-from-root-input-payment_method_id
      --project-create-from-root-input-tags string                    project-create-from-root-input-tags (JSON array)
      --project-create-from-root-input-type string                    project-create-from-root-input-type
      --request string                                                JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 projects](equinix_metalv1_projects.md)	 - Manage projects resources

