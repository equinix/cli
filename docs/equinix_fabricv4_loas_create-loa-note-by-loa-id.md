## equinix fabricv4 loas create-loa-note-by-loa-id

Create Loa Note

### Synopsis

The API provides capability to create Loa note by Loa ID

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 loas create-loa-note-by-loa-id [flags]
```

### Options

```
      --create-loa-note-additional-properties string   create-loa-note-additional-properties (JSON)
      --create-loa-note-comments string                The note content to add to this LOA. Notes are visible to both the issuer and requestor organizations. Use notes to communicate updates, clarifications, or additional context about the LOA.
  -h, --help                                           help for create-loa-note-by-loa-id
      --loa-id string                                  Loa UUID (required)
      --request string                                 JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 loas](equinix_fabricv4_loas.md)	 - Manage loas resources

