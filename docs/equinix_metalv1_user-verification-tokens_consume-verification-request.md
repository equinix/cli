## equinix metalv1 user-verification-tokens consume-verification-request

Verify a user using an email verification token

### Synopsis

Consumes an email verification token and verifies the user associated with it.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 user-verification-tokens consume-verification-request [flags]
```

### Options

```
  -h, --help                                        help for consume-verification-request
      --include string                              include field (JSON or string)
      --request string                              JSON payload for additional optional fields not exposed as flags
      --verify-email-additional-properties string   verify-email-additional-properties (JSON)
      --verify-email-user_token string              User verification token
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 user-verification-tokens](equinix_metalv1_user-verification-tokens.md)	 - Manage user-verification-tokens resources

