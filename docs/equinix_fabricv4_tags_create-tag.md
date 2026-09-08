## equinix fabricv4 tags create-tag

Create Tag

### Synopsis

Create Tag for Equinix Fabric™.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 tags create-tag [flags]
```

### Options

```
  -h, --help                                       help for create-tag
      --request string                             JSON payload for additional optional fields not exposed as flags
      --tag-request-additional-properties string   tag-request-additional-properties (JSON)
      --tag-request-display-name string            Display name of the Tag
      --tag-request-notifications string           tag-request-notifications (JSON array)
      --tag-request-type string                    Type of tag
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 tags](equinix_fabricv4_tags.md)	 - Manage tags resources

