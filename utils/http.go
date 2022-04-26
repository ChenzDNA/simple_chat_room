package utils

import "fmt"

func CorsHeader(origin string) string {
	return fmt.Sprintf(fmt.Sprintf("access-control-allow-credentials: true\naccess-control-allow-headers: Origin,Content-Type,Accept,User-Agent,Cookie,Authorization,X-Auth-Token,X-Requested-With\naccess-control-allow-methods: GET,PUT,POST,DELETE,PATCH,HEAD,CONNECT,OPTIONS,TRACE\naccess-control-allow-origin: %s\naccess-control-max-age: 3628800\n", origin))
}
