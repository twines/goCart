package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"goCart/app/admin"
	"goCart/app/api"
	"goCart/app/api/v1"
	"goCart/app/web"
	admin2 "goCart/middleware/admin"
	"goCart/middleware/cors"
	"goCart/middleware/jwt"
	"goCart/pkg/setting"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("resource/view/*/**/*.html")
	r.StaticFS("/static", http.Dir("resource/static"))
	r.StaticFS("/upload/images", http.Dir(setting.AppSetting.RuntimeRootPath+setting.AppSetting.ImageSavePath))
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
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
			adminGroup.POST("/login", admin.DoLogin)
			adminGroup.GET("/", admin.Index)

			//图片上传
			adminGroup.POST("/upload", admin.Upload)
			//多图上传
			adminGroup.POST("/upload/multi", admin.UploadMulti)

		}
		//admin已经登录
		{
			adminGroup.Use(admin2.Admin())
			//用户列表
			adminGroup.GET("/user/list", admin.User)
			adminGroup.GET("/user/add", admin.AddUserPage)

			//商品列表
			adminGroup.GET("/product/list", admin.GetProductList)
			//新增商品页面
			adminGroup.GET("/product/add", admin.AddProductPage)
			//新增商品
			adminGroup.POST("/product/add", admin.DoAddProduct)
			//编辑商品页面
			adminGroup.GET("/product/edit/:id", admin.Edit)
			//更新商品信息
			adminGroup.POST("/product/save/:id", admin.Save)

			adminGroup.POST("/product/off", admin.PostChangeProductStatus)

			//图片上传
			//adminGroup.POST("/upload", admin.Upload)
			////多图上传
			//adminGroup.POST("/upload/multi", admin.UploadMulti)

			//订单列表
			adminGroup.GET("/order/list", admin.OrderList)
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
