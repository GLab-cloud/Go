package main

import (
	"testing"

	"github.com/labstack/echo"

	"github-trend-BE/db"
	"github-trend-BE/handler"
	"github-trend-BE/repository/repo_implement"
	"github-trend-BE/router"
)

func TestMainSetup(t *testing.T) {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5433,
		UserName: "postgres",
		Password: "123",
		DbName:   "golang",
	}

	e := echo.New()

	userHandler := handler.UserHandler{
		UserRepo: repo_implement.NewUserRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()
}
