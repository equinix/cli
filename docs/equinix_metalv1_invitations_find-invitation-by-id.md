## equinix metalv1 invitations find-invitation-by-id

View an invitation

### Synopsis

Returns a single invitation. (It include the `invitable` to maintain backward compatibility but will be removed soon)

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 invitations find-invitation-by-id [flags]
```

### Options

```
  -h, --help             help for find-invitation-by-id
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

