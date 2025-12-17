## equinix fabricv4 cloud-events get-cloud-event

Get Cloud Event

### Synopsis

This API provides capability to retrieve a cloud event by uuid

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-events get-cloud-event [flags]
```

### Options

```
      --cloud-event-id string   Cloud Event UUID (required)
  -h, --help                    help for get-cloud-event
      --request string          JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cloud-events](equinix_fabricv4_cloud-events.md)	 - Manage cloud-events resources

