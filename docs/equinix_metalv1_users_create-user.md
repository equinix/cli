## equinix metalv1 users create-user

Create a user

### Synopsis

Creates a user.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 users create-user [flags]
```

### Options

```
      --exclude string                                   exclude field (JSON or string)
  -h, --help                                             help for create-user
      --include string                                   include field (JSON or string)
      --request string                                   JSON payload for additional optional fields not exposed as flags
      --user-create-input-additional-properties string   user-create-input-additional-properties (JSON)
      --user-create-input-company_name string            user-create-input-company_name
      --user-create-input-company_url string             user-create-input-company_url
      --user-create-input-customdata string              user-create-input-customdata (JSON)
      --user-create-input-emails string                  user-create-input-emails (JSON array)
      --user-create-input-first_name string              user-create-input-first_name
      --user-create-input-invitation_id string           user-create-input-invitation_id
      --user-create-input-last_name string               user-create-input-last_name
      --user-create-input-level string                   user-create-input-level
      --user-create-input-nonce string                   user-create-input-nonce
      --user-create-input-password string                user-create-input-password
      --user-create-input-phone_number string            user-create-input-phone_number
      --user-create-input-social_accounts string         user-create-input-social_accounts (JSON)
      --user-create-input-timezone string                user-create-input-timezone
      --user-create-input-title string                   user-create-input-title
      --user-create-input-two_factor_auth string         user-create-input-two_factor_auth
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 users](equinix_metalv1_users.md)	 - Manage users resources

