package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/situs-forum/internal/configs"
	membershipsHandler "github.com/ilhamrdh/situs-forum/internal/handlers/memberships"
	postsHandler "github.com/ilhamrdh/situs-forum/internal/handlers/posts"
	membershipsRepository "github.com/ilhamrdh/situs-forum/internal/repositories/memberships"
	postsRepo "github.com/ilhamrdh/situs-forum/internal/repositories/posts"
	membershipsService "github.com/ilhamrdh/situs-forum/internal/services/memberships"
	postsSvc "github.com/ilhamrdh/situs-forum/internal/services/posts"
	"github.com/ilhamrdh/situs-forum/pkg/internalsql"
)

func main() {
	r := gin.Default()
	var cfg *configs.Config

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal membaca config")
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DatabaseSourceName)
	if err != nil {
		log.Fatal("Gagal inisial database", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipsRepository.NewRepository(db)
	postRepo := postsRepo.NewRepository(db)

	membershipSvc := membershipsService.NewService(cfg, membershipRepo)
	postSvc := postsSvc.NewService(cfg, postRepo)

	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	postHandler := postsHandler.NewHandler(r, postSvc)

	membershipHandler.RegisterRoute()
	postHandler.PostRoute()

	r.Run(cfg.Service.Port)
}
