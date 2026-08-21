## equinix metalv1 bgp request-bgp-config

Requesting bgp config

### Synopsis

Requests to enable bgp configuration for a project.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 bgp request-bgp-config [flags]
```

### Options

```
      --bgp-config-request-input-additional-properties string   bgp-config-request-input-additional-properties (JSON)
      --bgp-config-request-input-asn int                        Autonomous System Number for local BGP deployment.
      --bgp-config-request-input-deployment_type string         bgp-config-request-input-deployment_type
      --bgp-config-request-input-md5 string                     The plaintext password to share between BGP neighbors as an MD5 checksum: * must be 10-20 characters long * may not include punctuation * must be a combination of numbers and letters * must contain at least one lowercase, uppercase, and digit character
      --bgp-config-request-input-use_case string                A use case explanation (necessary for global BGP request review).
  -h, --help                                                    help for request-bgp-config
      --id string                                               Project UUID (required)
      --include string                                          include field (JSON or string)
      --request string                                          JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 bgp](equinix_metalv1_bgp.md)	 - Manage bgp resources

