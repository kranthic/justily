package menu

import (
	"db"
	"time"
	"labix.org/v2/mgo/bson"
)

type menu struct{
	Id bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"n" json:"name"`
	Categories []category `bson:"c" json:"categories"`
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
}

type category struct{
	Name string `bson:"n" json:"name"`
	Items []catItem `bson:"ci" json:"-"`
}

type catItem struct{
	ItemId bson.ObjectId `bson:"id"`
	Sku []int `bson:s`
}

const coll = "menu"

func NewMenu() *menu{
	return &menu{}
}

func (this *menu) NewCategory() *category{
	return &category{}
}

func (this *category) NewCatItem() *catItem{
	return &catItem{}
}

func (this *category) AddItem(ci *catItem){
	this.Items = append(this.Items, *ci)
}

func (this *menu) AddCategory(c *category){
	this.Categories = append(this.Categories, *c)
}

func GetMenuById(id string, restaurantKey string) (*menu, error){
	s := db.NewMongoSession()
	defer s.Close()
	
	var m menu
	c := s.DB(restaurantKey).C(coll)
	err := c.FindId(bson.ObjectIdHex(id)).One(&m)
	
	return &m, err
}

func (this *menu) Save(restaurantKey string) error{
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
