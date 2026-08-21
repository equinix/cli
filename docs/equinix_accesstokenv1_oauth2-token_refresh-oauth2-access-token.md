## equinix accesstokenv1 oauth2-token refresh-oauth2-access-token

Execute refresh-oauth2-access-token operation

### Synopsis

Execute the refresh-oauth2-access-token operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix accesstokenv1 oauth2-token refresh-oauth2-access-token [flags]
```

### Options

```
  -h, --help                                   help for refresh-oauth2-access-token
      --payload-additional-properties string   payload-additional-properties (JSON)
      --payload-client_id string               API Consumer Key available under \"My Apps\" in developer portal
      --payload-client_secret string           API Consumer secret available under \"My Apps\" in developer portal
      --payload-refresh_token string           The OAuth2 refresh_token retrieved from the previous successful Access Token API call
      --request string                         JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix accesstokenv1 oauth2-token](equinix_accesstokenv1_oauth2-token.md)	 - Manage oauth2-token resources

