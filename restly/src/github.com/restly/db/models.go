package db

import (
	"errors"
//	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

var missing_req = errors.New("Required properties are missing")
var call_update = errors.New("Connot re-create object. Invoke Update() method to make updates")

type mongo struct{
	Session *mgo.Session
}

func (m *mongo) connect(){
	session, err := mgo.Dial("localhost")
	if err != nil{
		panic(err)
	}
	m.Session = session
}

func (m *mongo) session() *mgo.Session{
	if m.Session == nil{
		m.connect()
	}
	return m.Session.New()
}

var m = mongo{}
var db = "justily"

type User struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Email string `bson:"e" json:"e"`
	Password string `bson:"p" json:"-"`
	FirstName string `bson:"fn" json:"fn"`
	LastName string `bson:"ln" json:"ln"`
	Phone	string	`bson:"phn,omitempty" json:"phn"`
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
}


type Restaurant struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `bson:"n" json:"name"`
	About string `bson:"d,omitempty" json:"about"`
	Branches []RestaurantBranch `bson:"b" json:"branches"`
	AdminIds []bson.ObjectId `bson:"a" json:"-"`
	Admins []User	`bson:"-" json:"admins"`
	Key string `bson:"key" json:"key"` //Need a unique key that can be used to retrieve this restaurant. Name will not be unique
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
}

type RestaurantBranch struct{
	Name string `bson:"n" json:"name"`
	Address string `bson:"a" json:"address"`
	City string `bson:"c" json:"city"`
	State string `bson:"s" json:"state"`
	Zipcode string `bson:"z" json:"zipcode"`
}

type Item struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string	`bson:"n" json:"name"`
	Desc string `bson:"d" json:"desc"`
	Skus []Sku `bson:"s" json:"skus"`
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
	
	RestaurantKey string `bson:"-" json:"-"`
}

type Sku struct{
	Price float64	`bson:"p" json:"price"`
}

type Category struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string	`bson:"n" json:"name"`
	Desc string `bson:"d" json:"desc"`
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
	
	RestaurantKey string `bson:"-" json:"-"`
}

type Menu struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `bson:"n" json:"name"` //lunch, dinner
	Categories []MenuCategory `bson:"c" json:"categories"`
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
	
	RestaurantKey string `bson:"-" json:"-"`
}

type MenuCategory struct{
	Category bson.ObjectId `bson:"id" json:"id"`
	Items []CategoryItem `bson:"i" json:"items"`
}

type CategoryItem struct{
	ItemId bson.ObjectId `bson:"id" json:"itemId"`
	SkuIds []int `bson:"id" json:"skus"`
}

func (u *User) name() string{
	return "user"
}

func (u *User) valid() bool{
	return u.Email != "" && u.FirstName != "" && u.LastName != "" && u.Password != ""
}

func (u *User) Create() error{
	if !u.valid(){
		return missing_req
	}
	
	if bson.ObjectId.Valid(u.Id){
		return call_update
	}
	
	session := m.session()
	defer session.Close()
	
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()
	c := session.DB(db).C(u.name())
	c.EnsureIndex(mgo.Index{Key: []string{"e"},Unique: true})
	return c.Insert(&u)
}

func (u *User) Update() error{
	if !u.valid(){
		return missing_req
	}
	
	session := m.session()
	defer session.Close()
	
	u.UpdateTime = time.Now()
	c := session.DB(db).C(u.name())
	return c.Update(bson.M{"_id":u.Id}, u)
}

func (r *Restaurant) name() string{
	return "restaurant"
}

func (r *Restaurant) valid() bool{
	if r.Name == "" || r.AdminIds == nil || r.Branches == nil || len(r.Branches) == 0{
		return false
	} 
	
	for _,b := range(r.Branches){
		if b.Name == "" || b.Address == "" || b.City == "" || b.State == "" || b.Zipcode == ""{
			return false
		}	
	}
	return true 
}

func (r *Restaurant) Create() error{
	if !r.valid(){
		return missing_req
	}
	
	if bson.ObjectId.Valid(r.Id){
		return call_update
	}
	
	session := m.session()
	defer session.Close()
	
	r.CreateTime = time.Now()
	r.UpdateTime = time.Now()
	c := session.DB(db).C(r.name())
	c.EnsureIndex(mgo.Index{Key: []string{"key"}, Unique: true})
	return c.Insert(&r)
	
}

func (r *Restaurant) Update() error{
	if !r.valid(){
		return missing_req
	}
	
	session := m.session()
	defer session.Close()
	
	r.UpdateTime = time.Now()
	c := session.DB(db).C(r.name())
	return c.Update(bson.M{"_id":r.Id}, r)
}

func (i *Item) name() string{
	return "item"
}

func (i *Item) valid() bool{
	if i.Name == "" || i.RestaurantKey == ""{
		return false
	}
	
	return true
}

func (i *Item) Create() error{
	if !i.valid(){
		return missing_req
	}
	
	if bson.ObjectId.Valid(i.Id){
		return call_update
	}
	
	session := m.session()
	defer session.Close()
	
	i.CreateTime = time.Now()
	i.UpdateTime = time.Now()
	c := session.DB(i.RestaurantKey).C(i.name())
	return c.Insert(&i)
}

func (i *Item) Update() error{
	if !i.valid(){
		return missing_req
	}
	
	session := m.session()
	defer session.Close()
	
	i.UpdateTime = time.Now()
	c := session.DB(i.RestaurantKey).C(i.name())
	return c.Update(bson.M{"_id": i.Id}, i)
}


func (c *Category) name() string{
	return "category"
}

func (c *Category) valid() bool{

	if c.Name == "" || c.RestaurantKey == ""{
		return false
	}
	
	return true
}

func (c *Category) Create() error{
	if !c.valid(){
		return missing_req
	}
	
	if bson.ObjectId.Valid(c.Id){
		return call_update
	}
	
	session := m.session()
	defer session.Close()
	
	c.CreateTime = time.Now()
	c.UpdateTime = time.Now()
	collection := session.DB(c.RestaurantKey).C(c.name())
	return collection.Insert(&c)
}

func (c *Category) Update() error{
	if !c.valid(){
		return missing_req
	}
	
	session := m.session()
	defer session.Close()
	
	c.UpdateTime = time.Now()
	collection := session.DB(c.RestaurantKey).C(c.name())
	return collection.Update(bson.M{"_id": c.Id}, c)
}

func (m *Menu) name() string{
	return "menu"
}

func (m *Menu) valid() bool{

	if m.RestaurantKey == "" || m.Categories == nil || len(m.Categories) == 0{
		return false
	}
	
	for _,cat := range(m.Categories){
		if !bson.ObjectId.Valid(cat.Category) || cat.Items == nil || len(cat.Items) == 0{
			return false
		}
		
		for _,itm := range(cat.Items){
			if !bson.ObjectId.Valid(itm.ItemId) || itm.SkuIds == nil || len(itm.SkuIds) == 0{
				return false
			}
		}
	}
	
	return true
}

func (menu *Menu) Create() error{
	if !menu.valid(){
		return missing_req
	}
	
	if bson.ObjectId.Valid(menu.Id){
		return call_update
	}
	
	session := m.session()
	defer session.Close()
	
	menu.CreateTime = time.Now()
	menu.UpdateTime = time.Now()
	collection := session.DB(menu.RestaurantKey).C(menu.name())
	return collection.Insert(&menu)
}

func (menu *Menu) Update() error{
	if !menu.valid(){
		return missing_req
	}
	
	session := m.session()
	defer session.Close()
	
	menu.UpdateTime = time.Now()
	collection := session.DB(menu.RestaurantKey).C(menu.name())
	return collection.Update(bson.M{"_id": menu.Id}, menu)
}



func UserByEmail(email string) (user *User, err error){
	session := m.session()
	defer session.Close()
	c := session.DB(db).C("user")
	user = &User{}
	err = c.Find(bson.M{"e": email}).One(user)
	
	return user, err
}

func RestaurantById(id string) (*Restaurant, error){
	session := m.session()
	defer session.Close()
	
	r := &Restaurant{}
	c := session.DB(db).C(r.name())
	err := c.Find(bson.M{"_id": id}).One(r)
	
	return r, err
}

func RestaurantByKey(key string) (*Restaurant, error){
	session := m.session()
	defer session.Close()
	
	r := &Restaurant{}
	c := session.DB(db).C(r.name())
	err := c.Find(bson.M{"key": key}).One(r)
	
	return r, err
}

func ItemById(restaurantKey string, id bson.ObjectId) (*Item, error){
	session := m.session()
	defer session.Close()
	
	i := &Item{}
	c := session.DB(restaurantKey).C(i.name())
	err := c.FindId(id).One(i)
	
	return i, err
}

func CategoryById(restaurantKey string, id bson.ObjectId)(*Category, error){
	session := m.session()
	defer session.Close()
	
	ctg := &Category{}
	c := session.DB(restaurantKey).C(ctg.name())
	err := c.FindId(id).One(ctg)
	
	return ctg, err
}

func MenuByName(restaurantKey string, name string)(*Menu, error){
	session := m.session()
	defer session.Close()
	
	menu := &Menu{}
	c := session.DB(restaurantKey).C(menu.name())
	err := c.Find(bson.M{"n" : name}).One(menu)
	
	return menu, err
}

func Menus(restaurantKey string)(*[]Menu, error){
	session := m.session()
	defer session.Close()
	
	var menus []Menu
	menu := Menu{}
	c := session.DB(restaurantKey).C(menu.name())
	err := c.Find(nil).All(&menus)
	
	return &menus, err
}







