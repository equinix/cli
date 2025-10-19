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
  -h, --help                                                                           help for update-stream-subscription-by-uuid
      --request string                                                                 JSON payload for additional optional fields not exposed as flags
      --stream-id string                                                               Stream UUID (required)
      --stream-subscription-put-request-additional-properties string                   stream-subscription-put-request-additional-properties (required) (JSON)
      --stream-subscription-put-request-description string                             stream-subscription-put-request-description
      --stream-subscription-put-request-enabled                                        stream-subscription-put-request-enabled
      --stream-subscription-put-request-event-selector-additional-properties string    stream-subscription-put-request-event-selector-additional-properties (JSON)
      --stream-subscription-put-request-event-selector-except string                   stream-subscription-put-request-event-selector-except (JSON array)
      --stream-subscription-put-request-event-selector-include string                  stream-subscription-put-request-event-selector-include (JSON array)
      --stream-subscription-put-request-filters-additional-properties string           stream-subscription-put-request-filters-additional-properties (JSON)
      --stream-subscription-put-request-filters-and string                             stream-subscription-put-request-filters-and (JSON array)
      --stream-subscription-put-request-metric-selector-additional-properties string   stream-subscription-put-request-metric-selector-additional-properties (JSON)
      --stream-subscription-put-request-metric-selector-except string                  stream-subscription-put-request-metric-selector-except (JSON array)
      --stream-subscription-put-request-metric-selector-include string                 stream-subscription-put-request-metric-selector-include (JSON array)
      --stream-subscription-put-request-name string                                    stream-subscription-put-request-name
      --stream-subscription-put-request-sink-additional-properties string              stream-subscription-put-request-sink-additional-properties (JSON)
      --stream-subscription-put-request-sink-batch-enabled                             stream-subscription-put-request-sink-batch-enabled
      --stream-subscription-put-request-sink-batch-size-max int                        stream-subscription-put-request-sink-batch-size-max
      --stream-subscription-put-request-sink-batch-wait-time-max int                   stream-subscription-put-request-sink-batch-wait-time-max
      --stream-subscription-put-request-sink-credential string                         stream-subscription-put-request-sink-credential (JSON)
      --stream-subscription-put-request-sink-host string                               stream-subscription-put-request-sink-host
      --stream-subscription-put-request-sink-settings string                           stream-subscription-put-request-sink-settings (JSON)
      --stream-subscription-put-request-sink-type string                               stream-subscription-put-request-sink-type
      --stream-subscription-put-request-sink-uri string                                stream-subscription-put-request-sink-uri
      --subscription-id string                                                         Stream Subscription UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 stream-subscriptions](equinix_fabricv4_stream-subscriptions.md)	 - Manage stream-subscriptions resources

