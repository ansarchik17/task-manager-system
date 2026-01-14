package main

import (
	"context"
	"task-manager/config"
	handler "task-manager/handlers"
	"task-manager/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"*"},
	}
	r.Use(cors.New(corsConfig))
	err := loadConfig()
	if err != nil {
		panic(err)
	}
	connection, err := connectToDb()
	if err != nil {
		panic(err)
	}
	tasksRepository := repositories.NewTaskRepository(connection)
	taskHandler := handler.NewTaskHandler(tasksRepository)
	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks/get", taskHandler.GetAllTasks)
	r.GET("/task/get_task/:id", taskHandler.GetTaskById)
	r.PUT("/task/update/:id", taskHandler.UpdateTask)
	r.DELETE("/task/delete/:id", taskHandler.DeleteTask)
	r.Run(":8070")
}

func connectToDb() (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), config.Config.DbConnectionString)

	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func loadConfig() error {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	if err := viper.BindEnv("APP_HOST"); err != nil {
		viper.SetDefault("APP_HOST", ":8070")
	}
	if err := viper.BindEnv("DB_CONNECTION_STRING"); err != nil {
		viper.SetDefault("DB_CONNECTION_STRING", "postgres://postgres:ansar2007+A@localhost:5430/task-manager-system?sslmode=disable")
	}
	if err := viper.BindEnv("JWT_SECRET_KEY"); err != nil {
		viper.SetDefault("JWT_SECRET_KEY", "supersecretkey")
	}
	if err := viper.BindEnv("JWT_EXPIRES_IN"); err != nil {
		viper.SetDefault("JWT_EXPIRES_IN", "24h")
	}
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	var mapConfig config.MapConfig
	err = viper.Unmarshal(&mapConfig)
	if err != nil {
		return err
	}

	config.Config = &mapConfig

	return nil
}
