## equinix metalv1 organizations create-organization-invitation

Create an invitation for an organization

### Synopsis

In order to add a user to an organization, they must first be invited. To invite to several projects the parameter `projects_ids:[a,b,c]` can be used

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 organizations create-organization-invitation [flags]
```

### Options

```
  -h, --help                                            help for create-organization-invitation
      --id string                                       Organization UUID (required)
      --include string                                  include field (JSON or string)
      --invitation-input-additional-properties string   invitation-input-additional-properties (JSON)
      --invitation-input-invitee string                 invitation-input-invitee
      --invitation-input-message string                 invitation-input-message
      --invitation-input-organization_id string         invitation-input-organization_id
      --invitation-input-projects_ids string            invitation-input-projects_ids (JSON array)
      --invitation-input-roles string                   invitation-input-roles (JSON array)
      --request string                                  JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 organizations](equinix_metalv1_organizations.md)	 - Manage organizations resources

