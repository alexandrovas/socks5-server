# go-socks5-proxy

## List of supported config parameters

|ENV variable|Type|Default|Description|
|------------|----|-------|-----------|
|PROXY_PORT|String|1080|Set listen port for application inside docker container|
|ALLOWED_DEST_FQDN|String|EMPTY|Allowed destination address regular expression pattern. Default allows all.|

## Credentials

By default, proxy starts without authentication.

To enable authentication yout must add list of variables with usernames and passwords:

```sh
PROXY_CREDS_0_USERNAME='user0'
PROXY_CREDS_0_PASSWORD='password0'
PROXY_CREDS_1_USERNAME='user1'
PROXY_CREDS_1_PASSWORD='password1'
```
