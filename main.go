package main

import (
	"fmt"
"github.com/kuuyee/gin-learn/pkg/setting"
"github.com/kuuyee/gin-learn/routers"
"net/http"
)

func main()  {
	router := routers.InitRouter()

	//router.Run()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
