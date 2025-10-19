## equinix fabricv4 stream-alert-rules update-stream-alert-rule-by-uuid

Update Stream Alert Rules

### Synopsis

This API provides capability to update a user's stream alert rule

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 stream-alert-rules update-stream-alert-rule-by-uuid [flags]
```

### Options

```
      --alert-rule-id string                                                    alert rule UUID (required)
      --alert-rule-put-request-additional-properties string                     alert-rule-put-request-additional-properties (JSON)
      --alert-rule-put-request-critical-threshold string                        Stream alert rule metric critical threshold
      --alert-rule-put-request-description string                               Customer-provided stream description
      --alert-rule-put-request-enabled                                          Stream alert rule enabled status
      --alert-rule-put-request-metric-name string                               Stream alert rule metric name
      --alert-rule-put-request-name string                                      Customer-provided stream name
      --alert-rule-put-request-operand string                                   alert-rule-put-request-operand
      --alert-rule-put-request-resource-selector-additional-properties string   alert-rule-put-request-resource-selector-additional-properties (JSON)
      --alert-rule-put-request-resource-selector-include *                      ### Supported metric names to use on filters with property /subject:   * * - all events or metrics   * `*_/ports/<uuid>` - port metrics   * `*_/connections/<uuid>` - connection metrics   * `*_/metros/<metroCode>` - metro latency metrics (JSON array)
      --alert-rule-put-request-type string                                      alert-rule-put-request-type
      --alert-rule-put-request-warning-threshold string                         Stream alert rule metric warning threshold
      --alert-rule-put-request-window-size string                               Stream alert rule metric window size
  -h, --help                                                                    help for update-stream-alert-rule-by-uuid
      --request string                                                          JSON payload for additional optional fields not exposed as flags
      --stream-id string                                                        Stream UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 stream-alert-rules](equinix_fabricv4_stream-alert-rules.md)	 - Manage stream-alert-rules resources

