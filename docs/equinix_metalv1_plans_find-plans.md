## equinix metalv1 plans find-plans

Retrieve all plans

### Synopsis

Provides a listing of available plans to provision your device on.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 plans find-plans [flags]
```

### Options

```
      --categories string   categories field (JSON or string)
      --exclude string      exclude field (JSON or string)
  -h, --help                help for find-plans
      --include string      include field (JSON or string)
      --request string      JSON payload for additional optional fields not exposed as flags
      --slug string         slug field
      --type_ string        type_ field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 plans](equinix_metalv1_plans.md)	 - Manage plans resources

