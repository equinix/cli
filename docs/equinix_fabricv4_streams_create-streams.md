## equinix fabricv4 streams create-streams

Create Stream

### Synopsis

This API provides capability to create user's stream

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 streams create-streams [flags]
```

### Options

```
  -h, --help[=true|false]                                             help for create-streams
      --request string                                                JSON payload for additional optional fields not exposed as flags
      --stream-post-request-account-account-name string               Account name
      --stream-post-request-account-account-number int                Account number
      --stream-post-request-account-additional-properties string      stream-post-request-account-additional-properties (JSON)
      --stream-post-request-account-global-cust-id string             Account name
      --stream-post-request-account-global-org-id string              Global organization identifier
      --stream-post-request-account-global-organization-name string   Global organization name
      --stream-post-request-account-org-id int                        Customer organization identifier
      --stream-post-request-account-organization-name string          Customer organization name
      --stream-post-request-account-reseller-account-name string      Reseller account name
      --stream-post-request-account-reseller-account-number int       Reseller account number
      --stream-post-request-account-reseller-org-id int               Reseller customer organization identifier
      --stream-post-request-account-reseller-ucm-id string            Reseller account ucmId
      --stream-post-request-account-ucm-id string                     Account ucmId
      --stream-post-request-additional-properties string              stream-post-request-additional-properties (JSON)
      --stream-post-request-description string                        Customer-provided stream description
      --stream-post-request-name string                               Customer-provided stream name
      --stream-post-request-notifications string                      stream-post-request-notifications (JSON array)
      --stream-post-request-project-additional-properties string      stream-post-request-project-additional-properties (JSON)
      --stream-post-request-project-project-id string                 Subscriber-assigned project ID
      --stream-post-request-type string                               stream-post-request-type
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 streams](equinix_fabricv4_streams.md)	 - Manage streams resources

