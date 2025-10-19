## equinix fabricv4 stream-subscriptions create-stream-subscriptions

Create Subscription

### Synopsis

This API provides capability to create user's Stream Subscriptions

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 stream-subscriptions create-stream-subscriptions [flags]
```

### Options

```
  -h, --help                                                                            help for create-stream-subscriptions
      --request string                                                                  JSON payload for additional optional fields not exposed as flags
      --stream-id string                                                                Stream UUID (required)
      --stream-subscription-post-request-additional-properties string                   stream-subscription-post-request-additional-properties (required) (JSON)
      --stream-subscription-post-request-description string                             stream-subscription-post-request-description
      --stream-subscription-post-request-enabled                                        stream-subscription-post-request-enabled
      --stream-subscription-post-request-event-selector-additional-properties string    stream-subscription-post-request-event-selector-additional-properties (JSON)
      --stream-subscription-post-request-event-selector-except string                   stream-subscription-post-request-event-selector-except (JSON array)
      --stream-subscription-post-request-event-selector-include string                  stream-subscription-post-request-event-selector-include (JSON array)
      --stream-subscription-post-request-metric-selector-additional-properties string   stream-subscription-post-request-metric-selector-additional-properties (JSON)
      --stream-subscription-post-request-metric-selector-except string                  stream-subscription-post-request-metric-selector-except (JSON array)
      --stream-subscription-post-request-metric-selector-include string                 stream-subscription-post-request-metric-selector-include (JSON array)
      --stream-subscription-post-request-name string                                    stream-subscription-post-request-name
      --stream-subscription-post-request-sink-additional-properties string              stream-subscription-post-request-sink-additional-properties (JSON)
      --stream-subscription-post-request-sink-batch-enabled                             stream-subscription-post-request-sink-batch-enabled
      --stream-subscription-post-request-sink-batch-size-max int                        stream-subscription-post-request-sink-batch-size-max
      --stream-subscription-post-request-sink-batch-wait-time-max int                   stream-subscription-post-request-sink-batch-wait-time-max
      --stream-subscription-post-request-sink-credential string                         stream-subscription-post-request-sink-credential (JSON)
      --stream-subscription-post-request-sink-host string                               stream-subscription-post-request-sink-host
      --stream-subscription-post-request-sink-settings string                           stream-subscription-post-request-sink-settings (JSON)
      --stream-subscription-post-request-sink-type string                               stream-subscription-post-request-sink-type
      --stream-subscription-post-request-sink-uri string                                stream-subscription-post-request-sink-uri
      --stream-subscription-post-request-type string                                    stream-subscription-post-request-type
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 stream-subscriptions](equinix_fabricv4_stream-subscriptions.md)	 - Manage stream-subscriptions resources

