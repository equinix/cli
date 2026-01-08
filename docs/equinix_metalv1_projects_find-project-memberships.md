## equinix metalv1 projects find-project-memberships

Retrieve project memberships

### Synopsis

Returns all memberships in a project.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 projects find-project-memberships [flags]
```

### Options

```
  -h, --help                help for find-project-memberships
      --include string      include field (JSON or string)
      --page int            page field
      --per-page int        per-page field
      --project-id string   Project UUID (required)
      --request string      JSON payload for additional optional fields not exposed as flags
      --search string       search field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 projects](equinix_metalv1_projects.md)	 - Manage projects resources

