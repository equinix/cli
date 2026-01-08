## equinix metalv1 projects find-project-customdata

Retrieve the custom metadata of a project

### Synopsis

Provides the custom metadata stored for this project in json format

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 projects find-project-customdata [flags]
```

### Options

```
  -h, --help             help for find-project-customdata
      --id string        Project UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 projects](equinix_metalv1_projects.md)	 - Manage projects resources

