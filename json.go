package main



import(

	// "testing"

	"fmt"

)



// func TestAdd(t *testing.T){

// 	feed:=New()

// 	feed.Add(Item{})

// 	if len(feed.Items) != 1 {

// 		t.Errorf("Items was not added")

// 	}

// }



// func TestGetAll(t *testing.T) {

// 	feed:=New()

// 	feed.Add(Item{})

// 	if len(feed.Items) != 1 {

// 		t.Errorf("Items was not added")

// 	}

// }



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



func main(){

	feed:=New()

	

	fmt.Println(feed)



	feed.Add(Item{"Hello","How ya' doing mate?"})



	fmt.Println(feed)