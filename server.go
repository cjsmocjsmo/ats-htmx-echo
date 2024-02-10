package main

import (
	// "fmt"
	"html/template"
	"io"
	"net/http"
	// "net/mail"
	"os"
	// "regexp"
	// "strconv"
	// "unicode"

	// "github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func ats_port1(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port1", "WORKED")
}

func ats_port2(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port2", "WORKED")
}

func ats_port3(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port3", "WORKED")
}

func ats_port4(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port4", "WORKED")
}

func ats_port5(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port5", "WORKED")
}

func ats_port6(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port6", "WORKED")
}

func ats_port7(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port7", "WORKED")
}

func ats_port8(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port8", "WORKED")
}

func ats_port9(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port9", "WORKED")
}

func ats_port10(c echo.Context) error {
	return c.Render(http.StatusOK, "ats_port10", "WORKED")
}

func com_upload(c echo.Context) string {
	name := c.FormValue("name")
	email := c.FormValue("email")
	rating := c.FormValue("rating")
	comment := c.FormValue("comment")
	areInputsValid := checkInputs(name, email, rating, comment)
	if !areInputsValid {
		return "Invalid input"
	}
	

	file, err := c.FormFile("filepicker")
	if err != nil {
		println("filepicker error: ")
	}

	comid := atsUUID()

	outpath, err := save_file(comid, file)
	if err != nil {
		println("save_file error: ")
	}

	return outpath
}

func init() {
	godotenv.Load()
	os.Setenv("ATS_UPLOAD_DIR", "/usr/share/ats-htmx-echo/uploads")
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	t := &Template{
		templates: template.Must(template.ParseGlob("AtsTemplates/*")),
	}
	e.Renderer = t

	e.GET("/", ats_index)
	e.GET("/about", ats_about)
	e.GET("/comments", ats_comments)
	e.GET("/estimates", ats_estimates)
	e.GET("/images", ats_images)
	e.GET("/services", ats_services)
	e.GET("/videos", ats_videos)
	e.GET("/port1", ats_port1)
	e.GET("/port2", ats_port2)
	e.GET("/port3", ats_port3)
	e.GET("/port4", ats_port4)
	e.GET("/port5", ats_port5)
	e.GET("/port6", ats_port6)
	e.GET("/port7", ats_port7)
	e.GET("/port8", ats_port8)
	e.GET("/port9", ats_port9)
	e.GET("/port10", ats_port10)
	e.POST("/upload", com_upload)
	e.Static("/assets", "assets")
	e.Logger.Fatal(e.Start(":8181"))
}
