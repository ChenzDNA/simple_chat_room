package configs

import (
	"context"
	"flag"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
)

func init() {
	glog.Info(context.TODO(), "init config server")
	cfg := g.Cfg().GetAdapter().(*gcfg.AdapterFile)
	s := flag.String("env", "", "environment")
	flag.Parse()
	if *s != "" {
		glog.Info(context.TODO(), "environment: "+*s)
		cfg.SetFileName("config-" + *s)
	} else {
		cfg.SetFileName("config")
	}
}
