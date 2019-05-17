package main

import (
	_goodsController "github.com/alfssobsd/minishop/entrypoints/http"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
)

func GenrateClientId() echo.MiddlewareFunc {
	//log.Info("TEST")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, _ := session.Get("session", c)
			sess.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   86400 * 7,
				HttpOnly: true,
			}
			log.Info("TEST1")
			if len(sess.Values) == 0 {
				uuidString := uuid.NewV4()
				sess.Values["clientId"] = uuidString
				log.Info("new clientId = ", sess.Values["clientId"])
			}
			_ = sess.Save(c.Request(), c.Response())

			return next(c)
		}
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

	// Database connection
	db, err := mgo.Dial("mongodb://root:example@localhost")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Create indices
	if err = db.Copy().DB("minishop").C("accounts").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}
	//e.Use(session.Middleware(sessions.NewCookieStore([]byte("4h^FG:H+@t4D.3GgLt*9:9Hw3n9m*p?Cu#@Pf9v>8yRG?Par"))))
	//e.Use(GenrateClientId())
	_goodsController.GoodsRoutes(e, db)
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	e.Logger.Fatal(e.Start(":1323"))
}
