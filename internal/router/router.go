package router

import (
	"net/http"

	"are.moe/testtask/internal/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func New(productHandler handler.ProductHandler) *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	// router.Use(csrf.Middleware(csrf.Options{
	// 	Secret: "secret",
	// 	ErrorFunc: func(ctx *gin.Context) {
	// 		ctx.String(400, "CSRF token mismatch")
	// 		ctx.Abort()
	// 	},
	// }))

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials:       true,
		AllowHeaders:           []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With"},
		AllowBrowserExtensions: true,
		AllowPrivateNetwork:    true,
	}))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.Status(http.StatusNotFound)
	})
	router.NoMethod(func(ctx *gin.Context) {
		ctx.Status(http.StatusMethodNotAllowed)
	})

	productRouter := router.Group("/product")
	{
		productRouter.POST("/", productHandler.AddProduct)
		productRouter.GET("/:productID", productHandler.GetProductByID)
		productRouter.PUT("/:productID", productHandler.UpdateProduct)
		productRouter.DELETE("/:productID", productHandler.RemoveProduct)
		productRouter.GET("/", productHandler.ListProducts)
		productRouter.GET("/find", productHandler.FindProductsByName)
	}

	return router
}
