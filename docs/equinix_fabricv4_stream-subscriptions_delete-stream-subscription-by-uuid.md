## equinix fabricv4 stream-subscriptions delete-stream-subscription-by-uuid

Delete Subscription

### Synopsis

This API provides capability to delete user's Stream Subscriptions

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 stream-subscriptions delete-stream-subscription-by-uuid [flags]
```

### Options

```
  -h, --help                     help for delete-stream-subscription-by-uuid
      --request string           JSON payload for request body fields
      --stream-id string         Stream UUID (required)
      --subscription-id string   Stream Subscription UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 stream-subscriptions](equinix_fabricv4_stream-subscriptions.md)	 - Manage stream-subscriptions resources

