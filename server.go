package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"html/template"
	"io"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func ats_index(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_index", "WORKED")
}

func ats_about(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_about", "WORKED")
}

func ats_comments(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_comments", "WORKED")
}

func ats_estimates(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_estimates", "WORKED")
}

func ats_images(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_images", "WORKED")
}

func ats_services(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_services", "WORKED")
}

func ats_videos(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_videos", "WORKED")
}


func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	t := &Template{
		templates: template.Must(template.ParseGlob("AtsTemplates/*.html")),
	}
	e.Renderer = t
	
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.GET("/", ats_index)
	e.GET("/about", ats_about)
	e.GET("/comments", ats_comments)
	e.GET("/estimates", ats_estimates)
	e.GET("/images", ats_images)
	e.GET("/services", ats_services)
	e.GET("/videos", ats_videos)
	e.Logger.Fatal(e.Start(":8181"))
}
