package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Product মডেল
type Product struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Price     float64            `bson:"price" json:"price"`
	Quantity  int                `bson:"quantity" json:"quantity"`
	Category  string             `bson:"category" json:"category"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

var productCollection *mongo.Collection

// MongoDB initialization
func initDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	productCollection = client.Database("shop").Collection("products")
	fmt.Println("✅ MongoDB connected")
}

func main() {
	initDB()

	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

	// ১. সব প্রোডাক্ট পাওয়া
	r.GET("/api/products", func(c *gin.Context) {
		cursor, err := productCollection.Find(context.Background(), bson.M{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(context.Background())

		var products []Product
		if err = cursor.All(context.Background(), &products); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"products": products,
			"count":    len(products),
		})
	})

	// ২. একটি প্রোডাক্ট পাওয়া (ID দিয়ে)
	r.GET("/api/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "ভুল ID"})
			return
		}

		var product Product
		err = productCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&product)
		if err != nil {
			c.JSON(404, gin.H{"error": "প্রোডাক্ট পাওয়া যায়নি"})
			return
		}

		c.JSON(200, product)
	})

	// ৩. নতুন প্রোডাক্ট তৈরি
	r.POST("/api/products", func(c *gin.Context) {
		var product Product
		if err := c.BindJSON(&product); err != nil {
			c.JSON(400, gin.H{"error": "ভুল তথ্য"})
			return
		}

		product.ID = primitive.NewObjectID()
		product.CreatedAt = time.Now()

		_, err := productCollection.InsertOne(context.Background(), product)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{
			"message": "প্রোডাক্ট তৈরি সফল",
			"product": product,
		})
	})

	// ৪. প্রোডাক্ট আপডেট
	r.PUT("/api/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "ভুল ID"})
			return
		}

		var updateData map[string]interface{}
		if err := c.BindJSON(&updateData); err != nil {
			c.JSON(400, gin.H{"error": "ভুল তথ্য"})
			return
		}

		delete(updateData, "id")
		delete(updateData, "_id")

		_, err = productCollection.UpdateOne(
			context.Background(),
			bson.M{"_id": objectID},
			bson.M{"$set": updateData},
		)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "প্রোডাক্ট আপডেট সফল"})
	})

	// ৫. প্রোডাক্ট ডিলিট
	r.DELETE("/api/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "ভুল ID"})
			return
		}

		_, err = productCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "প্রোডাক্ট ডিলিট সফল"})
	})

	fmt.Println("🛒 ই-কমার্স API সার্ভার শুরু: http://localhost:8080")
	r.Run(":8080")
}
