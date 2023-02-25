module github.com/kuuyee/gin-learn

go 1.13

require (
	cloud.google.com/go v0.47.0 // indirect
	github.com/astaxie/beego v1.12.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.49.0
	github.com/jinzhu/gorm v1.9.11
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.1.0 // indirect
	gopkg.in/ini.v1 v1.49.0 // indirect
)

replace (
	github.com/kuuyee/gin-learn/conf => /Users/yeekuu/dev/go/src/go-web-learn/pkg/conf
	github.com/kuuyee/gin-learn/middleware => /Users/yeekuu/dev/go/src/go-web-learn/middleware
	github.com/kuuyee/gin-learn/models => /Users/yeekuu/dev/go/src/go-web-learn/models
	github.com/kuuyee/gin-learn/pkg/setting => /Users/yeekuu/dev/go/src/go-web-learn/pkg/setting
	github.com/kuuyee/gin-learn/routers => /Users/yeekuu/dev/go/src/go-web-learn/routers
)
