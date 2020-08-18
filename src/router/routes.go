package router

import (
	"booleanservice/src/controller"
	"booleanservice/src/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

//SetupRouter is
func SetupRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware, err := middleware.JWTMiddleware()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	r.POST("/login", authMiddleware.LoginHandler)
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		// claims, exists := c.Get("JWT_PAYLOAD")
		// fmt.Println(reflect.TypeOf(claims))
		// for key, value := range claims.(jwt.MapClaims) {
		// 	claims[key] = value
		// }
		// if !exists {
		// 	fmt.Println("not exist : ", make(jwt.MapClaims))
		// } else {
		// 	x, ok := claims.(jwt.MapClaims)
		// 	if !ok {
		// 		fmt.Println("Error PArsing")
		// 	} else {
		// 		fmt.Println("X : ", x)
		// 	}
		// }

		// log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	// r.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth := r
	// Refresh time can be longer than token timeout
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/", controller.CreateValue)
		auth.GET("/", controller.ListAll)
		auth.DELETE("/:id", controller.DeleteValue)
		auth.PATCH("/:id", controller.UpdateValue)
		auth.GET("/:id", controller.GetValue)
	}
	// if err := http.ListenAndServe(":"+"8080", r); err != nil {
	// 	log.Fatal(err)
	// }
	return auth
}
func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"userID":   "sads",
		"userName": "sas",
		"text":     "Hello World.",
	})
}
