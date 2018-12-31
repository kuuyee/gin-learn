package main

import (
	"fmt"
	"net/http"

	"./pkg/setting"
	"./routers"
)

func main()  {
	router:= routers.InitRouter()

	fmt.Printf("HTTPPort : %d \n",setting.HTTPPort)

	s:= &http.Server{
		Addr: fmt.Sprintf(":%d",setting.HTTPPort),
		Handler : router,
		ReadTimeout:setting.ReadTimeout,
		WriteTimeout:setting.WriteTimeout,
		MaxHeaderBytes:1 << 20,
	}

	s.ListenAndServe()
}
