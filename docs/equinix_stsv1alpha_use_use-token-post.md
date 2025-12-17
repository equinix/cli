## equinix stsv1alpha use use-token-post

Method for UseTokenPost

### Synopsis

An OAuth 2.0 token endpoint supporting RFC 8693 token exchange, used to exchange an OIDC ID token issued by a trusted OIDC provider to a trusted client for an access token that can be used access other Equinix product APIs.

Use --request flag to provide optional JSON payload fields.

```
equinix stsv1alpha use use-token-post [flags]
```

### Options

```
      --grant-type string           grant-type field
  -h, --help                        help for use-token-post
      --request string              JSON payload for additional optional fields not exposed as flags
      --scope string                scope field
      --subject-token string        subject-token field
      --subject-token-type string   subject-token-type field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix stsv1alpha use](equinix_stsv1alpha_use.md)	 - Manage use resources

