package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	//return c.Renderx()
	greeting := "Hola World!"
	return c.Render(greeting)
}
