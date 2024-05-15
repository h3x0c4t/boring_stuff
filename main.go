package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gofiber/fiber/v2"
)

var DB *sql.DB

type ProjectPutJSON struct {
	Name string `json:"name"`
}

type ProjectDeleteJSON struct {
	ID int `json:"id"`
}

type ProjectGetJSON struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	TimeCreated string `json:"time_created"`
}

func CreateProject(name string) error {
	_, err := DB.Exec("INSERT INTO projects (name) VALUES (?)", name)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProject(id int) error {
	_, err := DB.Exec("DELETE FROM projects WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func GetProjects() ([]ProjectGetJSON, error) {
	rows, err := DB.Query("SELECT id, name, datetime(time_created) FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := []ProjectGetJSON{}
	for rows.Next() {
		project := ProjectGetJSON{}
		err = rows.Scan(&project.ID, &project.Name, &project.TimeCreated)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func initializeDB(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		time_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Подключение к БД
	var err error
	DB, err = sql.Open("sqlite3", "./evilmsg_db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	initializeDB(DB)

	// Веб-сервер
	app := fiber.New()

	// Basic авторизация
	// app.Use(basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"admin": "1234",
	// 	},
	// }))

	// Фронтенд
	app.Static("/", "./frontend/evilmsg/dist")

	// Получение списка проектов
	// curl 127.0.0.1:3000/api/projects
	app.Get("/api/projects", func(c *fiber.Ctx) error {
		projects, err := GetProjects()
		if err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(500)
		}

		return c.JSON(projects)
	})

	// Новый проект
	// curl -X PUT -H "Content-Type: application/json" --data '{"name":"test"}' 127.0.0.1:3000/api/projects
	app.Put("/api/projects", func(c *fiber.Ctx) error {
		project := new(ProjectPutJSON)

		err := c.BodyParser(project)
		if err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(400)
		}

		if project.Name == "" {
			log.Println("WARNING: Name is empty")
			return c.SendStatus(400)
		}

		if err := CreateProject(project.Name); err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(500)
		}

		return c.SendStatus(200)
	})

	// Удаление проекта
	// curl -X DELETE -H "Content-Type: application/json" --data '{"id":1}' 127.0.0.1:3000/api/projects
	app.Delete("/api/projects", func(c *fiber.Ctx) error {
		project := new(ProjectDeleteJSON)

		err := c.BodyParser(project)
		if err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(400)
		}

		if project.ID == 0 {
			log.Println("WARNING: ID is empty")
			return c.SendStatus(400)
		}

		if err := DeleteProject(project.ID); err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(500)
		}

		return c.SendStatus(200)
	})

	app.Listen(":3000")
}
