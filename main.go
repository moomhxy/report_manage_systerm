package main

import (
	"fmt"

	"github.com/moomhxy/report_manage_systerm/config"
)

func main() {
	if err := config.InitConfig(); err != nil {
		fmt.Printf("init config error: %+v", err)
		return
	}

	r := RegisterRouter()
	if err := r.Run(":" + config.Conf.Port); err != nil {
		fmt.Printf("server startup failed, err:%+v\n", err)
	}
}
