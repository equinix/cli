## equinix fabricv4 stream-subscriptions update-stream-subscription-by-uuid

Update Subscription

### Synopsis

This API provides capability to update user's Stream Subscriptions

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 stream-subscriptions update-stream-subscription-by-uuid [flags]
```

### Options

```
  -h, --help                     help for update-stream-subscription-by-uuid
      --request string           JSON payload for request body. Available fields: stream-subscription-put-request (StreamSubscriptionPutRequest)
      --stream-id string         Stream UUID (required)
      --subscription-id string   Stream Subscription UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 stream-subscriptions](equinix_fabricv4_stream-subscriptions.md)	 - Manage stream-subscriptions resources

