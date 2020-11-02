package main



import (

	"fmt"

	"github.com/gin-gonic/gin"

	"io/ioutil"

)



func HomePage(c *gin.Context) {

	c.JSON(200, gin.H{

		"message": "Hello World",

	})

}

func PostHomePage(c *gin.Context) {

	body:= c.Request.Body

	value,err := ioutil.ReadAll(body)

	if err!=nil{

		fmt.Println(err.Error())

	}

	c.JSON(200, gin.H{

		"message": string(value),

	})

}

func QueryStrings(c *gin.Context) {

	name:=c.Query("name")

	age:=c.Query("age")



	c.JSON(200, gin.H{

		"name":name,

		"age":age,

	})

}

func PathParameters(c *gin.Context){

	name:=c.Param("name")

	age:=c.Param("age")



	c.JSON(200, gin.H{

		"name":name,

		"age":age,

	})

}



func main() {

	fmt.Printf("Hello World")



	r := gin.Default()

	r.GET("/", HomePage)

	r.POST("/", PostHomePage)

	r.GET("/query", QueryStrings) // /query?name=priya&age=22

	r.GET("/path/:name/:age",PathParameters)// /path/testname/100/

	r.Run()

}