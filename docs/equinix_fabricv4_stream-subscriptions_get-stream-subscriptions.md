## equinix fabricv4 stream-subscriptions get-stream-subscriptions

Get Subscriptions

### Synopsis

This API provides capability to retrieve stream subscriptions

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 stream-subscriptions get-stream-subscriptions [flags]
```

### Options

```
  -h, --help               help for get-stream-subscriptions
      --request string     JSON payload for request body. Available fields: limit (int32), offset (int32)
      --stream-id string   Stream UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 stream-subscriptions](equinix_fabricv4_stream-subscriptions.md)	 - Manage stream-subscriptions resources

