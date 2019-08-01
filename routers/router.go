package routers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"goCart/app/admin"
	"goCart/app/api"
	"goCart/app/api/v1"
	"goCart/app/web"
	admin2 "goCart/middleware/admin"
	"goCart/middleware/cors"
	"goCart/middleware/jwt"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("resource/view/*/**/*.html")
	r.StaticFS("/static", http.Dir("resource/static"))

	store := sessions.NewCookieStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: int(30 * time.Minute), //30min
		Path:   "/",
	})
	//后台管理
	{
		adminGroup := r.Group("/admin")
		adminGroup.Use(sessions.Sessions("goCartAdmin", store))

		//admin未登录
		{
			adminGroup.GET("/login", admin.Login)
			adminGroup.GET("/", admin.Index)
		}
		//admin已经登录
		{
			adminGroup.Use(admin2.Admin())
			adminGroup.GET("/user", admin.User)
		}
	}
	//web前端
	{
		webGroup := r.Group("/")
		webGroup.Use(sessions.Sessions("goCart", store))
		//未登录
		{
			webGroup.GET("/", web.Index)
			webGroup.GET("/login", web.Login)
		}
		//已经登录
		{
			webGroup.GET("/user", web.User)
		}
	}
	//api接口
	{
		apiV1 := r.Group("/api/v1")
		//不需要认证
		{
			apiV1.POST("/api/auth", api.GetAuth)
		}
		//需要认证
		{
			apiV1.Use(jwt.JWT())
			apiV1.Use(cors.Cors())
			apiV1.GET("/", v1.Index)
		}
	}
	return r
}
