## equinix metalv1 two-factor-auth enable-tfa-app

Enable two factor auth using app

### Synopsis

Enables two factor authentication using authenticator app.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 two-factor-auth enable-tfa-app [flags]
```

### Options

```
  -h, --help                 help for enable-tfa-app
      --request string       JSON payload for additional optional fields not exposed as flags
      --x-otp-token string   x-otp-token field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 two-factor-auth](equinix_metalv1_two-factor-auth.md)	 - Manage two-factor-auth resources

