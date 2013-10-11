package main 

import (
	"fmt"
	"github.com/restly/db"
	"crypto/sha256"
	"io"
	"os"
	"labix.org/v2/mgo/bson"
//	"encoding/json"
)


func createUser(){
	u := db.User{}
	u.Email = "kranthi.chalasani@gmail.com"
	u.FirstName = "Kranthi"
	u.LastName = "Chalasani"
	h := sha256.New()
	io.WriteString(h, "venkat")
	u.Password = fmt.Sprintf("%x", h.Sum(nil))
	u.Phone = "248-470-8466"
	
	fmt.Println(u.Create())
}

func createRestaurant(){
	
	u, err := db.UserByEmail("kranthi.chalasani@gmail.com")
	if err != nil{
		panic(err)
	}
	r := db.Restaurant{}
	r.Name = "Peacock India Restaurants"
	
	b := make([]byte, 16)
	f,err := os.Open("/dev/urandom")
	if err != nil{
		panic(err)
	}
	f.Read(b)
    r.Key = fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
    
    r.About = "Authentic South Indian Restaurant"
    r.AdminIds = []bson.ObjectId{u.Id}
    
    branch := db.RestaurantBranch{}
    branch.Name = "Peacock Fremont"
    branch.Address = "2342 Walnut Ave"
    branch.City = "Fremont"
    branch.State = "CA"
    branch.Zipcode = "94582"
    
    r.Branches = []db.RestaurantBranch{branch}
    
    fmt.Println("Restaurant created", r.Create() == nil)
    
}

func main() {
//	u := db.User{}
//	u.Email = "kranthi.chalasani@gmail.com"
//	u.FirstName = "Kranthi"
//	u.LastName = "Chalasani"
//	h := sha256.New()
//	io.WriteString(h, "venkat")
//	u.Password = fmt.Sprintf("%x", h.Sum(nil))
//	u.Phone = "248-470-8466"
//	
//	fmt.Println(u.Create())
	
//	usr, err := db.UserByEmail("prathima82@gmail.com")
//	if err != nil{
//		panic(err)
//	}
//	usrJson, err := json.Marshal(usr)
//	if err != nil{
//		panic(err)
//	}
//	fmt.Println(string(usrJson))

//	usr.Email= "kranthi.chalasani@gmail.com"
//	fmt.Println(usr.Update())
	createUser()
	createRestaurant()

}

