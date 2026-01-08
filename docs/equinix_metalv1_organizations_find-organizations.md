## equinix metalv1 organizations find-organizations

Retrieve all organizations

### Synopsis

Returns a list of organizations that are accessible to the current user.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 organizations find-organizations [flags]
```

### Options

```
      --exclude string            exclude field (JSON or string)
  -h, --help                      help for find-organizations
      --include string            include field (JSON or string)
      --page int                  page field
      --per-page int              per-page field
      --personal string           personal field
      --request string            JSON payload for additional optional fields not exposed as flags
      --without-projects string   without-projects field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 organizations](equinix_metalv1_organizations.md)	 - Manage organizations resources

