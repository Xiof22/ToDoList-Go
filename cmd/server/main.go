package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Xiof22/ToDoList/config"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/middleware"
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/Xiof22/ToDoList/internal/repository"
	"github.com/Xiof22/ToDoList/internal/repository/mysql"
	"github.com/Xiof22/ToDoList/internal/router"
	"github.com/Xiof22/ToDoList/internal/service"
	_ "github.com/Xiof22/ToDoList/internal/validator"
	"github.com/gorilla/sessions"
	"net/http"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	loc, err := time.LoadLocation(cfg.TimezoneLocation)
	if err != nil {
		fmt.Printf("Failed to load location %s: %v\nFalling back to UTC...\n", cfg.TimezoneLocation, err)
		time.Local = time.UTC
	} else {
		time.Local = loc
	}

	cs := sessions.NewCookieStore([]byte(cfg.CookieStoreKey))
	repo, err := mysql.New(cfg.DBDSN)
	if err != nil {
		panic(err)
	}

	svc := service.New(repo)
	h := handlers.New(svc, cs, cfg)
	mw := middleware.New(cs, cfg)
	r := router.New(h, mw)

	if err := seedAdmin(repo, cfg.AdminEmail, cfg.AdminPassword); err != nil {
		fmt.Println("Failed to seed admin:", err)
	}

	fmt.Printf("Serving start on %d...\n", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r); err != nil {
		panic(err)
	}
}

func seedAdmin(repo repository.Repository, email, password string) error {
	ctx := context.Background()

	_, err := repo.GetUserByEmail(ctx, email)
	if err == nil {
		return nil
	}

	if !errors.Is(err, errorsx.ErrUserNotFound) {
		return err
	}

	admin, err := models.NewUser(email, password)
	if err != nil {
		return err
	}
	admin.Role = models.Admin

	repo.CreateUser(ctx, admin)
	return nil
}
