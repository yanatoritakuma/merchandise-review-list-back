package router

import (
	"merchandise-review-list-backend/controller"
	"net/http"

	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	uc controller.IUserController,
	pc controller.IProductController,
	rc controller.IReviewPostController,
	lc controller.ILikeController,
	cc controller.ICommentController,
	mc controller.IMoneyManagementController,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode, //PostMan使用する時に使用
		// CookieMaxAge: 60,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	u := e.Group("/user")
	u.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	// JWTが必須なエンドポイント
	u.GET("", uc.GetLoggedInUser)
	u.PUT("", uc.UpdateUser)
	u.DELETE("/:userId", uc.DeleteUser)

	p := e.Group("/product")

	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	// JWTが必須なエンドポイント
	p.POST("", pc.CreateProduct)
	p.PUT("/:productId", pc.UpdateTimeLimit)
	p.GET("/userProducts", pc.GetMyProducts)
	p.GET("/timeLimitAll", pc.GetMyProductsTimeLimitAll)
	p.DELETE("/:productId", pc.DeleteProduct)
	p.GET("/timeLimitYearMonth", pc.GetMyProductsTimeLimitYearMonth)
	p.GET("/timeLimitDate", pc.GetMyProductsTimeLimitDate)

	r := e.Group("/reviewPosts")
	// JWTが必須なエンドポイント
	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	r.POST("", rc.CreateReviewPost)
	r.PUT("/:postId", rc.UpdateReviewPost)
	r.GET("/userReviewPosts", rc.GetMyReviewPosts)
	r.DELETE("/:postId", rc.DeleteReviewPost)
	r.GET("/likes", rc.GetMyLikes)
	// JWTが必須でないエンドポイント
	e.GET("/reviewPosts/postId/:postId", rc.GetReviewPostById)
	e.GET("/reviewPosts/lists/:category", rc.GetReviewPostLists)

	l := e.Group("/like")
	// JWTが必須なエンドポイント
	l.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	l.POST("", lc.CreateLike)
	l.DELETE("/:postUserId", lc.DeleteLike)

	c := e.Group("/comment")
	// JWTが必須なエンドポイント
	c.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	c.POST("", cc.CreateComment)
	c.DELETE("/:id", cc.DeleteComment)

	// JWTが必須でないエンドポイント
	e.GET("/comment", cc.GetCommentsByPostId)

	m := e.Group("/moneyManagement")
	// JWTが必須なエンドポイント
	m.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	m.POST("", mc.CreateMoneyManagement)

	return e
}
