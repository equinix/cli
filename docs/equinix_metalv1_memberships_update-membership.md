## equinix metalv1 memberships update-membership

Update the membership

### Synopsis

Updates the membership.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 memberships update-membership [flags]
```

### Options

```
  -h, --help                                            help for update-membership
      --id string                                       Membership UUID (required)
      --include string                                  include field (JSON or string)
      --membership-input-additional-properties string   membership-input-additional-properties (JSON)
      --membership-input-role string                    membership-input-role (JSON array)
      --request string                                  JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 memberships](equinix_metalv1_memberships.md)	 - Manage memberships resources

