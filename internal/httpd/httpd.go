package httpd

import (
	"embed"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"time"

	"tronopay/internal/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	docs "tronopay/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed templates/*
var htmlFiles embed.FS

var pages = make(map[string]*template.Template)

func RunServer(port string, services *service.Services) *http.Server {
	// Init
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard
	router := gin.New()
	router.Use(gin.Recovery())

	// Static files: bootstrap, styles
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		services.Log.Fatal("error parsing FS - ", err)
	}
	router.StaticFS("/static", http.FS(staticFS))

	// Dinamycal templates
	tmplFileNames := []string{
		"index.html",
		// add new template here
	}
	for _, tmplFileName := range tmplFileNames {
		_template, err := template.New("").Funcs(regFuncs()).ParseFS(
			htmlFiles,
			"templates/base.html",
			"templates/"+tmplFileName,
		)
		if err != nil {
			services.Log.Fatal("error parsing templates - ", err)
		}
		pages[tmplFileName] = _template
	}

	// Routes
	controller := &Controller{
		pages:    pages,
		router:   router,
		services: services,
	}
	router.GET("/", controller.MainPage)
	router.GET("/qr/:value", controller.GenerateQRCode)

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1", RateLimiterMiddleware())
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("", controller.ShowTaskList)
			tasks.GET(":guid", controller.ShowTask)
			tasks.POST("", controller.CreateTask)
			// tasks.PATCH(":guid", controller.UpdateTask)
			// tasks.PUT(":guid", controller.UpdateFullTask)
			tasks.DELETE(":guid", controller.DeleteTask)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Run web server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("app - httpd - listen error: %s\n", err)
		}
	}()

	return srv
}

func RateLimiterMiddleware() gin.HandlerFunc {
	// 1 request per second
	limiter := rate.NewLimiter(rate.Every(time.Second), 1)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests on this route. Limit is 1 request per second.",
			})
			return
		}
		c.Next()
	}
}

func regFuncs() template.FuncMap {
	return template.FuncMap{
		// "name": func(arg string) string { return "" },
	}
}
