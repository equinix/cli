## equinix metalv1 invitations accept-invitation

Accept an invitation

### Synopsis

Accept an invitation.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 invitations accept-invitation [flags]
```

### Options

```
  -h, --help             help for accept-invitation
      --id string        Invitation UUID (required)
      --include string   include field (JSON or string)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 invitations](equinix_metalv1_invitations.md)	 - Manage invitations resources

