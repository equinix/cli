## equinix accesstokenv1 o-auth2-token get-o-auth2-access-token

Generate New Access Token

### Synopsis

This API handles authentication and authorization of the API developer. It returns an access_token which would be valid for 60 minutes. The user account will get locked after 5 successive invalid attempts.

Use --request flag to provide optional JSON payload fields.

```
equinix accesstokenv1 o-auth2-token get-o-auth2-access-token [flags]
```

### Options

```
  -h, --help                                   help for get-o-auth2-access-token
      --payload-additional-properties string   payload-additional-properties (JSON)
      --payload-client_id string               API Consumer Key available under \"My Apps\" in developer portal
      --payload-client_secret string           API Consumer secret available under \"My Apps\" in developer portal
      --payload-grant_type string              The OAuth2 grant type used for authorization. Supported values are \"password\" & \"client_credentials\". user_name and password is not considered in case this value is \"client_credentials\". If the grant_type is not passed, by default it would consider \"password\" type in which user_name and password is required. Note that the password grant type is deprecated. Recommended to use grant_type of 'client_credentials' instead.
      --payload-password_encoding string       For enhanced security, you may encrypt the password value while requesting for an access_token. Currently only \"md5-b64\" hashing is supported. Any other value would treat password value as raw string
      --payload-user_name string               Deprecated - The Equinix username used to access portals
      --payload-user_password string           Deprecated - The Equinix user password used to access portals
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

