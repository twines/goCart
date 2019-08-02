##关于goCart

##后台管理采用了<a href="https://github.com/ColorlibHQ/AdminLTE">adminlte</a>

##go依赖

```
gorm
gin
jwt
```
##目录接口

```

├─app
│  ├─admin【后台管理】
│  │      login.go
│  │      user.go
│  │      
│  ├─api【API接口】
│  │  │  auth.go
│  │  │  
│  │  └─v1【API接口的版本】
│  │          index.go
│  │          
│  └─web【前台应用】
│          index.go
│          login.go
│          user.go
│          
├─conf【配置文件】
│      app.ini
│      
├─middleware【中间件】
│  ├─admin【后台认证的中间件】
│  │      admin.go
│  │      
│  ├─cors【允许跨域的中间件】
│  │      cors.go
│  │      
│  └─jwt【API认证中间件】
│          jwt.go
│          
├─models【model层】
│      admin.go
│      auth.go
│      models.go
│      user.go
│      
├─pkg【自定义的包】
│  ├─auth
│  │      auth.go
│  │      
│  ├─e【http返回信息】
│  │      code.go
│  │      msg.go
│  │      
│  ├─logging【日志管理】
│  │      file.go
│  │      log.go
│  │      
│  ├─setting【配置管理】
│  │      setting.go
│  │      
│  └─util【工具类】
│          jwt.go
│          md5.go
│          pagination.go
│          
├─resource【前端静载资源】
│  ├─static
│  │
│  └─view
│      └─admin
│          └─login
│                  index.html
│                  
├─routers【路由】
│      router.go
│      
├─runtime
│  └─logs
└─task


```
##路由说明

```
    //后台管理
	{
		adminGroup := r.Group("/admin")
		adminGroup.Use(sessions.Sessions("goCartAdmin", store))

		//admin未登录
		{
			adminGroup.GET("/login", admin.Login)
			adminGroup.POST("/login", admin.DoLogin)
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

```
##PS:还有很多不完善的地方，希望越来越好，心怀向往，砥砺前行ヾ(◍°∇°◍)ﾉﾞ。