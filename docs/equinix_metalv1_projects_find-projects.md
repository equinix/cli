## equinix metalv1 projects find-projects

Retrieve all projects

### Synopsis

Returns a collection of projects that the current user is a member of.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 projects find-projects [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-projects
      --include string   include field (JSON or string)
      --name string      name field
      --page int         page field
      --per-page int     per-page field
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 projects](equinix_metalv1_projects.md)	 - Manage projects resources

