## equinix metalv1 authentication find-project-a-p-i-keys

Retrieve all API keys for the project.

### Synopsis

Returns all API keys for a specific project.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 authentication find-project-a-p-i-keys [flags]
```

### Options

```
  -h, --help             help for find-project-a-p-i-keys
      --id string        Project UUID (required)
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

* [equinix metalv1 authentication](equinix_metalv1_authentication.md)	 - Manage authentication resources

