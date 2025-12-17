## equinix metalv1 projects create-project-invitation

Create an invitation for a project

### Synopsis

In order to add a user to a project, they must first be invited.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 projects create-project-invitation [flags]
```

### Options

```
  -h, --help                                            help for create-project-invitation
      --include string                                  include field (JSON or string)
      --invitation-input-additional-properties string   invitation-input-additional-properties (JSON)
      --invitation-input-invitee string                 invitation-input-invitee
      --invitation-input-message string                 invitation-input-message
      --invitation-input-organization_id string         invitation-input-organization_id
      --invitation-input-projects_ids string            invitation-input-projects_ids (JSON array)
      --invitation-input-roles string                   invitation-input-roles (JSON array)
      --project-id string                               Project UUID (required)
      --request string                                  JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 projects](equinix_metalv1_projects.md)	 - Manage projects resources

