## equinix fabricv4 stream-alert-rules create-stream-alert-rules

Create Stream Alert Rules

### Synopsis

This API provides capability to create user's Stream Alert Rules

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 stream-alert-rules create-stream-alert-rules [flags]
```

### Options

```
      --alert-rule-post-request-additional-properties string                     alert-rule-post-request-additional-properties (required) (JSON)
      --alert-rule-post-request-critical-threshold string                        Stream alert rule metric critical threshold
      --alert-rule-post-request-description string                               Customer-provided stream description
      --alert-rule-post-request-enabled                                          Stream alert rule enabled status
      --alert-rule-post-request-metric-name string                               Stream alert rule metric name
      --alert-rule-post-request-name string                                      Customer-provided stream name
      --alert-rule-post-request-operand string                                   alert-rule-post-request-operand
      --alert-rule-post-request-resource-selector-additional-properties string   alert-rule-post-request-resource-selector-additional-properties (JSON)
      --alert-rule-post-request-resource-selector-include *                      ### Supported metric names to use on filters with property /subject:   * * - all events or metrics   * `*_/ports/<uuid>` - port metrics   * `*_/connections/<uuid>` - connection metrics   * `*_/metros/<metroCode>` - metro latency metrics (JSON array)
      --alert-rule-post-request-type string                                      alert-rule-post-request-type
      --alert-rule-post-request-warning-threshold string                         Stream alert rule metric warning threshold
      --alert-rule-post-request-window-size string                               Stream alert rule metric window size
  -h, --help                                                                     help for create-stream-alert-rules
      --request string                                                           JSON payload for additional optional fields not exposed as flags
      --stream-id string                                                         Stream UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 stream-alert-rules](equinix_fabricv4_stream-alert-rules.md)	 - Manage stream-alert-rules resources

