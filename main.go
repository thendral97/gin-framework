package main



import(

	"github.com/gin-gonic/gin"

)



type Item struct{

	Title string `json:"title"`

	Post  string `json:"post"`

}



type Repo struct{

	Items []Item

}



func New() *Repo{

	return &Repo{}

}



func (r *Repo) Add(item Item){

	r.Items=append(r.Items,item)

}



func (r *Repo) GetAll() []Item {

	return r.Items

}



func PingGet(c *gin.Context){

	c.JSON(200,gin.H{

		"Hello":"Found Me",

	})

}



func NewsFeedGet(feed *Repo) gin.HandlerFunc{

	return func(c *gin.Context){

		results:=feed.GetAll()

		c.JSON(200, results)

	}

}



type newsfeedPostRequest struct{

	Title string `json:"title"`

	Post string `json:"post"`

}



func NewsFeedPost(feed *Repo) gin.HandlerFunc{

	return func(c *gin.Context){

		requestBody:=newsfeedPostRequest{}

		c.Bind(&requestBody)

		

		item:= Item{

			Title: requestBody.Title,

			Post: requestBody.Post,

		}



		feed.Add(item)



		c.JSON()

	}

}



func main() {



	feed:=New()

	

	r:=gin.Default()



	r.GET("/ping",PingGet)

	r.GET("newsfeed",NewsFeedGet(feed))

	r.POST("newsfeed",NewsFeedPost(feed))

	r.Run()

}