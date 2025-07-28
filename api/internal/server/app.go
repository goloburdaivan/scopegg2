package server

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func NewApp(router *gin.Engine) *App {
	return &App{router: router}
}

func (a *App) Run() error {
	return a.router.Run(":8080")
}
