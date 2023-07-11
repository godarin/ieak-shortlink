package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Engine struct {
	Api *echo.Echo
	DB  *Table
}

func Create(db *Table) *Engine {
	var engine = Engine{
		Api: echo.New(),
		DB:  db,
	}
	return &engine
}

func (api *Engine) InitRoute() {

	api.Api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://link.ieak.fun", "https://link.ieak.fun"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	api.Api.File("/", "static/index.html")
	api.Api.GET("/:string", func(c echo.Context) error {
		value, _, err := api.DB.GetValue(c.Param("string"))
		if err != nil {
			return echo.ErrNotFound
		} else {
			return c.Redirect(http.StatusPermanentRedirect, value)
		}
	})
	type Data struct {
		Link string `json:"link"`
	}

	api.Api.POST("/create", func(c echo.Context) error {

		data := Data{}
		err := json.NewDecoder(c.Request().Body).Decode(&data)
		if err != nil {
			log.Fatalf("Failed reading the request body %s", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
		}
		compares := CompressString(data.Link)
		err = api.DB.AddKey(compares, data.Link)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
		}

		return c.String(http.StatusOK, "http://link.ieak.fun/"+compares)
	})
}
func (api *Engine) Run(port string) {
	log.Fatal(api.Api.Start(port))
}

func (api *Engine) AddRoute(method string, path string, handler echo.HandlerFunc) {
	api.Api.Add(
		method,
		path,
		handler,
	)
}
