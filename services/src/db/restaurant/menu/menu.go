package menu

import (
	"db"
	"time"
	dbi "db/restaurant/item"
	"labix.org/v2/mgo/bson"
)

type Menu struct{
	Id bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"n" json:"name"`
	Categories []Category `bson:"c" json:"categories"`
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
}

type Category struct{
	Name string `bson:"n" json:"name"`
	Items []CatItem `bson:"ci" json:"items"`
}

type CatItem struct{
	ItemId bson.ObjectId `bson:"id"`
	Sku []int `bson:"s"`
	Item *dbi.Item `bson:"-"`
}

const coll = "menu"

func NewMenu() *Menu{
	return &Menu{}
}

func (this *Menu) NewCategory() *Category{
	return &Category{}
}

func (this *Category) NewCatItem() *CatItem{
	return &CatItem{}
}

func (this *Category) AddItem(ci *CatItem){
	this.Items = append(this.Items, *ci)
}

func (this *Menu) AddCategory(c *Category){
	this.Categories = append(this.Categories, *c)
}

func GetMenuById(id string, restaurantKey string) (*Menu, error){
	s := db.NewMongoSession()
	defer s.Close()
	
	var m Menu
	c := s.DB(restaurantKey).C(coll)
	err := c.FindId(bson.ObjectIdHex(id)).One(&m)
	
	return &m, err
}

func (this *Menu) Save(restaurantKey string) error{
	s := db.NewMongoSession()
	defer s.Close()
	
	c := s.DB(restaurantKey).C(coll)
	var ins bool
	now := time.Now()
	if this.Id == ""{
		this.Id = bson.NewObjectId()
		this.CreateTime = now
		ins = true
	}
	this.UpdateTime = now
	
	if ins{ return c.Insert(&this)} else{ return c.UpdateId(this.Id, &this)}
}
