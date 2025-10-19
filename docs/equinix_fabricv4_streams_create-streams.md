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
  -h, --help                                                       help for create-streams
      --request string                                             JSON payload for additional optional fields not exposed as flags
      --stream-post-request-additional-properties string           stream-post-request-additional-properties (JSON)
      --stream-post-request-description string                     Customer-provided stream description
      --stream-post-request-name string                            Customer-provided stream name
      --stream-post-request-project-additional-properties string   stream-post-request-project-additional-properties (JSON)
      --stream-post-request-project-project-id string              Subscriber-assigned project ID
      --stream-post-request-type string                            stream-post-request-type
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 streams](equinix_fabricv4_streams.md)	 - Manage streams resources

