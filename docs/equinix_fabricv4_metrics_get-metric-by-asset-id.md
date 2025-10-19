## equinix fabricv4 metrics get-metric-by-asset-id

Get Metrics by Asset Id

### Synopsis

This API provides capability to retrieve Metrics of an asset id

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 metrics get-metric-by-asset-id [flags]
```

### Options

```
      --asset string      asset (required)
      --asset-id string   asset UUID (required)
  -h, --help              help for get-metric-by-asset-id
      --limit int         limit field
      --name string       name field
      --offset int        offset field
      --request string    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 metrics](equinix_fabricv4_metrics.md)	 - Manage metrics resources

