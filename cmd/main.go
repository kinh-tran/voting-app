package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

type Candidate struct {
	Name        string
	Image       string
	Preselected bool
	Id          int
}

func DummyVotingData() *Data[Candidate] {
	return &Data[Candidate]{
		Data: []Candidate{
			{
				Name:        "Ice Cream",
				Image:       "https://tb-static.uber.com/prod/image-proc/processed_images/d9782f5be876bced7b8ad068ad0d38f7/16bb0a3ab8ea98cfe8906135767f7bf4.webp",
				Preselected: false,
				Id:          1,
			},
			{
				Name:        "Pizza",
				Image:       "https://imgs.search.brave.com/RA2aE_owg_BIacd3RIATkWqz-R2KH2P0fBD-aciBBAo/rs:fit:860:0:0/g:ce/aHR0cHM6Ly90My5m/dGNkbi5uZXQvanBn/LzAwLzU3Lzg0Lzkw/LzM2MF9GXzU3ODQ5/MDgyX1RaYTdxOGxJ/UktYQ2dKcXNpdTRw/MDlwbU44RmtQMklp/LmpwZw",
				Preselected: false,
				Id:          2,
			},
			{
				Name:        "Hot Dogs",
				Image:       "https://imgs.search.brave.com/mXX8oOIOqHKKJ5C3VCXNJqZazcShGQ-7F7_jLl47j1A/rs:fit:860:0:0/g:ce/aHR0cHM6Ly9tZWRp/YS5pc3RvY2twaG90/by5jb20vaWQvMTg1/MTIzMzc3L3Bob3Rv/L2hvdGRvZy5qcGc_/cz02MTJ4NjEyJnc9/MCZrPTIwJmM9d0N2/eFhkTVh6bWtSM2VE/T0hlaWZuZW5IRFMx/b3dDNWIyTnRpSzdi/TzlVOD0",
				Preselected: false,
				Id:          3,
			},
			{
				Name:        "Salad",
				Image:       "https://imgs.search.brave.com/Y3n3r0lsFhLdzFcj0eTd_YCeq9ojvZWB_QwRWs17EZ4/rs:fit:860:0:0/g:ce/aHR0cHM6Ly93d3cu/aGF1dGVhbmRoZWFs/dGh5bGl2aW5nLmNv/bS93cC1jb250ZW50/L3VwbG9hZHMvMjAy/MS8xMC9MZW50aWwt/VGFiYm91bGVoLVNh/bGFkLTEwLmpwZw",
				Preselected: false,
				Id:          4,
			},
		},
	}
}

func DummyResultsData() *Data[Result] {
	return &Data[Result]{
		Data: []Result{
			{
				Name:  "Bar",
				Image: "/images/bars.svg",
				Id:    1,
			},
			// {
			// 	Name:  "Line",
			// 	Image: "",
			// 	Id:    2,
			// },
		},
	}
}

type Option struct {
	Name string
	Id   int
}

type Result struct {
	Name  string
	Image string
	Id    int
}

type Data[T any] struct {
	Data []T
}

func DummyRegisterData() *Data[Option] {
	return &Data[Option]{
		Data: []Option{
			{
				Name: "Scan QR Code",
				Id:   1,
			},
			{
				Name: "Link Authenticator App",
				Id:   2,
			},
		},
	}
}

func DummySettingsData() *Data[Option] {
	return &Data[Option]{
		Data: []Option{
			{
				Name: "Report compromised blockchain node",
				Id:   1,
			},
		},
	}
}

type FormData struct {
	Errors map[string]string
	Values map[string]string
}

func NewFormData() FormData {
	return FormData{
		Errors: map[string]string{},
		Values: map[string]string{},
	}
}

type PageData[T any] struct {
	Data Data[T]
	Form FormData
}

func VotingData(data Data[Candidate], form FormData) PageData[Candidate] {
	return PageData[Candidate]{
		Data: data,
		Form: form,
	}
}

func RegisterData(data Data[Option], form FormData) PageData[Option] {
	return PageData[Option]{
		Data: data,
		Form: form,
	}
}

func SettingsData(data Data[Option], form FormData) PageData[Option] {
	return PageData[Option]{
		Data: data,
		Form: form,
	}
}

func ResultsData(data Data[Result], form FormData) PageData[Result] {
	return PageData[Result]{
		Data: data,
		Form: form,
	}
}
func main() {

	e := echo.New()
	e.Static("/images", "images")
	e.Renderer = newTemplate()
	e.Use(middleware.Logger())

	e.GET("/", func(context echo.Context) error {
		return context.Render(200, "index.html", NewFormData())
	})

	e.POST("/login", func(context echo.Context) error {
		data := DummyVotingData()
		return context.Render(200, "voting", VotingData(*data, NewFormData()))
	})

	e.GET("/logout", func(context echo.Context) error {
		return context.Render(200, "logout", NewFormData())
	})

	e.GET("/register", func(context echo.Context) error {
		data := DummyRegisterData()
		return context.Render(200, "register", RegisterData(*data, NewFormData()))
	})

	e.GET("/settings", func(context echo.Context) error {
		data := DummySettingsData()
		return context.Render(200, "settings", SettingsData(*data, NewFormData()))
	})

	e.POST("/vote", func(context echo.Context) error {
		return context.Render(200, "voted", NewFormData())
	})

	e.GET("/results", func(context echo.Context) error {
		data := DummyResultsData()
		return context.Render(200, "results", ResultsData(*data, NewFormData()))
	})

	e.Logger.Fatal(e.Start(":4444"))
}
