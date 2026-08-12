package main

import (
	"log"
	"net/http"

	"github.com/criptycpizza7/habit-tracker/common"
	"github.com/criptycpizza7/habit-tracker/common/config"
	gormdb "github.com/criptycpizza7/habit-tracker/common/db/gorm_db"
	gormaccessors "github.com/criptycpizza7/habit-tracker/common/db/gorm_db/gorm_accessors"
	"github.com/criptycpizza7/habit-tracker/http/internal"
	"github.com/criptycpizza7/habit-tracker/http/internal/web"
	httpin_integration "github.com/ggicci/httpin/integration"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func main() {
	http_config := internal.HTTPConfig{} // TODO: перенести в какой-нибудь BuildServer
	config.MustLoadConfig(&http_config)

	db_config := config.DBConfig{}
	config.MustLoadConfig(&db_config)

	gorm_cfg := &gorm.Config{
		TranslateError: true,
	}
	db := gormdb.New(db_config, gorm_cfg)

	accessors := gormaccessors.New(db)
	controllers := common.MakeControllers(accessors)
	handlers := web.MakeHandlers(controllers)

	httpin_integration.UseGochiURLParam("path", chi.URLParam)

	middlewares := make([]func(http.Handler) http.Handler, 0)
	middlewares = append(middlewares, middleware.Recoverer)
	middlewares = append(middlewares, middleware.Logger)

	log.Println("Starting server")

	if err := web.RunServer(http_config, handlers, middlewares); err != nil { // поменять кучу аргументов на структуру
		panic(err)
	}
}
