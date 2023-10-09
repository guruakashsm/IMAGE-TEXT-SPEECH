package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors" // Import the cors package
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	Text  string `json:"text"`
	Email string `json:"email"`
}
type Signin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SignUpData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var mongoClient *mongo.Client

func main() {
	// Initialize MongoDB client
	mongoURI := "mongodb+srv://guru:guru@banking.sy1piq8.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	mongoClient = client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Create a Gin router
	router := gin.Default()

	// Use CORS middleware to allow requests from your Python application
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5000"} // Replace with the actual port of your Python application
	router.Use(cors.New(config))

	// Define a route to handle OPTIONS requests at "/result" (for pre-flight requests)
	router.OPTIONS("/result", handleOptions)
	router.OPTIONS("/signin", SigninOptions)
	router.OPTIONS("/signup", SignupOptions)
	router.OPTIONS("/getdata", GetOptions)

	// Define a route to handle POST requests at "/result"
	router.POST("/result", handlePostResult)
	router.POST("/signin", Signincon)
	router.POST("/signup", Signup)
	router.GET("/getdata", Getalldata)
	router.GET("/delete", Deleteall)

	// Start the HTTP server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func handleOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Status(http.StatusOK)
}
func GetOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "GET")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Status(http.StatusOK)
}
func SignupOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Status(http.StatusOK)
}

func SigninOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Status(http.StatusOK)
}

func handlePostResult(c *gin.Context) {
	var data Data
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(data)
	data.Email = Email
	// Insert data into MongoDB
	collection := mongoClient.Database("Image-Speech").Collection("TEXT")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into MongoDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data inserted into MongoDB successfully"})
}

var Email string

func Signincon(c *gin.Context) {
	var data Signin
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(data)
	collection := mongoClient.Database("Image-Speech").Collection("USER")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user SignUpData
	var result int64
	filter := bson.M{"email": data.Email}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		result = 0
		c.JSON(http.StatusOK, result)
	}
	fmt.Println(user)
	if user.Password == data.Password {
		Email = user.Email
		result = 1
		fmt.Println(result)
		c.JSON(http.StatusOK, result)
	} else {
		result = 0
		c.JSON(http.StatusOK, result)
	}
}

func Signup(c *gin.Context) {
	var data SignUpData
	var result int64
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(data)
	collection := mongoClient.Database("Image-Speech").Collection("USER")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		result = 0
		c.JSON(http.StatusOK, result)

	}else{
	result = 1
	c.JSON(http.StatusOK, result)}
}


 func Getalldata(c *gin.Context) {
	collection := mongoClient.Database("Image-Speech").Collection("TEXT")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var results []Data
	cur, err := collection.Find(ctx, bson.M{"email":Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data from MongoDB"})
		return
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var data Data
		err := cur.Decode(&data)
		fmt.Println(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding data"})
			return
		}
		results = append(results, data)
	}
    fmt.Println(results)
	c.JSON(http.StatusOK, results)
}

func Deleteall(c *gin.Context) {

	collection := mongoClient.Database("Image-Speech").Collection("TEXT")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	 collection.DeleteMany(ctx, bson.M{})
	
}