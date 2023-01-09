package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"example/cmd/handlers"
	"example/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getMaxId() int {
	var maxid int
	for _, p := range services.Products {
		if p.Id > maxid {
			maxid = p.Id
		}
	}
	return maxid
}

func main() {
	// Open Json
	jsonFile, err := os.Open("/Users/mbiagetti/Bootcamp/GoWeb/Practica_Post/products.json")
	// Control Error
	if err != nil {
		log.Panic("Error al recuperar datos")
	}
	fmt.Println("Datos cargados correctamente")
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &services.Products)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(services.Products); i++ {
		fmt.Println("Product id: " + strconv.Itoa(services.Products[i].Id))
		fmt.Println("Product name: " + services.Products[i].Name)
		fmt.Println("Product quantity: " + strconv.Itoa(services.Products[i].Quantity))
		fmt.Println("Product code_value: " + services.Products[i].Code_value)
		fmt.Println("Product is_published: " + strconv.FormatBool(services.Products[i].Is_published))
		fmt.Println("Product price: " + strconv.FormatFloat(services.Products[i].Price, 'E', -1, 32))
	}

	fmt.Println(services.Products)
	services.LastID = getMaxId()
	fmt.Printf("%d", services.LastID)

	// server
	sv := gin.Default()

	// router
	products := sv.Group("/products")
	products.GET("", handlers.Get)
	products.POST("", handlers.Create)
	products.GET("/ping", handlers.Pong)

	// start
	if err := sv.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
