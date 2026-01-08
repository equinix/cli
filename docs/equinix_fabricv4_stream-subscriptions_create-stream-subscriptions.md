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
      --stream-subscription-post-request-additional-properties string                   stream-subscription-post-request-additional-properties (JSON)
      --stream-subscription-post-request-description string                             Customer-provided stream subscription description
      --stream-subscription-post-request-enabled                                        Stream subscription enabled status
      --stream-subscription-post-request-event-selector-additional-properties string    stream-subscription-post-request-event-selector-additional-properties (JSON)
      --stream-subscription-post-request-event-selector-except string                   stream-subscription-post-request-event-selector-except (JSON array)
      --stream-subscription-post-request-event-selector-include string                  stream-subscription-post-request-event-selector-include (JSON array)
      --stream-subscription-post-request-metric-selector-additional-properties string   stream-subscription-post-request-metric-selector-additional-properties (JSON)
      --stream-subscription-post-request-metric-selector-except string                  stream-subscription-post-request-metric-selector-except (JSON array)
      --stream-subscription-post-request-metric-selector-include string                 stream-subscription-post-request-metric-selector-include (JSON array)
      --stream-subscription-post-request-name string                                    Customer-provided stream subscription name
      --stream-subscription-post-request-sink-additional-properties string              stream-subscription-post-request-sink-additional-properties (JSON)
      --stream-subscription-post-request-sink-batch-enabled                             batch mode on/off
      --stream-subscription-post-request-sink-batch-size-max int                        maximum batch size
      --stream-subscription-post-request-sink-batch-wait-time-max int                   maximum batch waiting time
      --stream-subscription-post-request-sink-credential string                         stream-subscription-post-request-sink-credential (JSON)
      --stream-subscription-post-request-sink-host string                               sink host
      --stream-subscription-post-request-sink-settings string                           stream-subscription-post-request-sink-settings (JSON)
      --stream-subscription-post-request-sink-type string                               stream-subscription-post-request-sink-type
      --stream-subscription-post-request-sink-uri string                                any publicly reachable http endpoint
      --stream-subscription-post-request-type string                                    stream-subscription-post-request-type
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 stream-subscriptions](equinix_fabricv4_stream-subscriptions.md)	 - Manage stream-subscriptions resources

