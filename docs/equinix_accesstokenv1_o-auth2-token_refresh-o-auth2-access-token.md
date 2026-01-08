## equinix accesstokenv1 o-auth2-token refresh-o-auth2-access-token

Renew Access Tokens

### Synopsis

Use this API to refresh an access token using its refresh token. A valid refresh token is needed to retrieve a new access_token.

Use --request flag to provide optional JSON payload fields.

```
equinix accesstokenv1 o-auth2-token refresh-o-auth2-access-token [flags]
```

### Options

```
  -h, --help                                   help for refresh-o-auth2-access-token
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

* [equinix accesstokenv1 o-auth2-token](equinix_accesstokenv1_o-auth2-token.md)	 - Manage o-auth2-token resources

