## equinix fabricv4 application-links attach-app-domain-to-app-link

Attach App Domain to App Link

### Synopsis

This API provides ability to attach the user's App Domain to App Link.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-links attach-app-domain-to-app-link [flags]
```

### Options

```
      --app-domain-id string   App Domain UUID (required)
      --app-link-id string     App Link UUID (required)
  -h, --help                   help for attach-app-domain-to-app-link
      --request string         JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-links](equinix_fabricv4_application-links.md)	 - Manage application-links resources

