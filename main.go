package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID int `json:"id"`
	Completed bool `json:"Completed"`
	Body string `json:"body"`
}

func main() {
	fmt.Println("hello world")
	app := fiber.New()
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file");
	}

	PORT := os.Getenv("PORT");
	todos:= []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	app.Post("/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} // assigned with default value

		if err :=  c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "todo body is required"})
		}

		todo.ID = len(todos) + 1;
		todos = append(todos, *todo);

		return c.Status(201).JSON(todo);
	})

	app.Patch("/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id");

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i]);
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "todo not found!"});
	})

	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id");
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(todos[i]);
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "todo not found!"});
	})

	log.Fatal(app.Listen(":" + PORT))
}