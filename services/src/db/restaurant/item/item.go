package item

import (
	"db"
//	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Item struct{
	Id bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"n" json:"name"`
	Desc string `bson:"d" json:"desc"`
	Skus []Sku `bson:"s" json:"skus"`
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
}

type Sku struct{
	Price float32 `bson:"p" json:"price"`
}

const coll = "item"

func NewItem() *Item{
	return &Item{}
}

func GetByItemId(id string, restaurantKey string) (*Item, error){
	s := db.NewMongoSession()
	defer s.Close()
	
	c := s.DB(restaurantKey).C(coll)
	var i Item 
	err := c.FindId(bson.ObjectIdHex(id)).One(&i)
	return &i, err
}

func (this *Item) NewSku() *Sku{
	return &Sku{}
}

func (this *Item) AddSku(s *Sku){
	this.Skus = append(this.Skus, *s)
}

func (this *Item) Save(restaurantKey string) error{
	s := db.NewMongoSession()
	defer s.Close()
	
	var ins bool
	now := time.Now()
	if this.Id == ""{
		this.Id = bson.NewObjectId()
		this.CreateTime = now
		ins = true
	}
	this.UpdateTime = now
	c := s.DB(restaurantKey).C(coll)
	if ins{ return c.Insert(&this)} else{ return c.UpdateId(this.Id, &this)}
}



