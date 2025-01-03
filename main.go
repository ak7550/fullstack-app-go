package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        int    `json:"id" bson: "_id"`
	Completed bool   `json:"Completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection;

func main() {
	fmt.Println("hello")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err);
	}

	mongoUri := os.Getenv("MONGO_URI");

	clientOptions := options.Client().ApplyURI(mongoUri);
	client, err := mongo.Connect(context.Background(), clientOptions);

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background());

	err = client.Ping(context.Background(), nil);
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected with the DB")

	collection = client.Database("golang_db").Collection("todos")

	app := fiber.New();

	app.Get("api/todos", getAllTodos);
	app.Post("api/todods", createTodo);
	app.Patch("api/todos/:id", updateTodo);
	app.Delete("api/todos/:id", deleteTodo);

	port := os.Getenv("PORT");

	if port == "" {
		port  = "5000"
	}

	log.Fatal(app.Listen(":"+port));
}

func getAllTodos(c *fiber.Ctx) error {
	var todos []Todo;
	cursor, err := collection.Find(context.Background(), bson.M{});

	if err != nil {
		return err;
	}

	defer cursor.Close(context.Background());

	for cursor.Next(context.Background()) {
		var todo Todo;
		if err := cursor.Decode(&todo); err != nil {
			return err;
		}

		todos = append(todos, todo);
	}

	return c.Status(200).JSON(todos);
}
func createTodo(c *fiber.Ctx) error {
	
}
func updateTodo(c *fiber.Ctx) error {
	
}
func deleteTodo(c *fiber.Ctx) error {
	
}
