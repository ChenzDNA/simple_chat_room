package utils

import "fmt"

func CorsHeader(origin string) string {
	return fmt.Sprintf("access-control-allow-credentials: true\r\naccess-control-allow-headers: Origin,Content-Type,Accept,User-Agent,Cookie,Authorization,X-Auth-Token,X-Requested-With\r\naccess-control-allow-methods: GET,PUT,POST,DELETE,PATCH,HEAD,CONNECT,OPTIONS,TRACE\r\naccess-control-allow-origin: %s\r\naccess-control-max-age: 3628800\r\n", origin)
}
