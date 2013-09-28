package main

import (
	"fmt"
	"net/http"
	"code.google.com/p/gorest"

	"strings"
)


var userStore map[int]User
var itemStore []Item
var orderStore []Order

func main(){

	gorest.RegisterService(new(OrderService))
	http.Handle("/",gorest.Handle())    
	if err := http.ListenAndServe(":8787",nil); err != nil {
		fmt.Printf("Error %v\n" , err )
	}
}

type User struct {
	Id int
	OrderIds []int
}

type Item struct {
	Id int
	AvailableStock int
}
type Order struct {
	Id int
	UserId int
	ItemId int
	Amount float64
	Discount float64
	Cancelled bool
}

type Discover struct {
	User User
	Item Item
	Order Order
}


//************************Define Service***************************

type OrderService struct{
    //Service level config
    gorest.RestService    `root:"/orders-service/" consumes:"application/json" produces:"application/json"`

    //End-Point level configs: Field names must be the same as the corresponding method names,
    // but not-exported (starts with lowercase)

    discover 	 gorest.EndPoint `method:"GET"  path:"/discover/"      output:"Discover"`
    userDetails gorest.EndPoint `method:"GET"  path:"/users/{Id:int}" output:"User"`
    listUsers   gorest.EndPoint `method:"GET"  path:"/users/"         output:"[]User"`
    listItems   gorest.EndPoint `method:"GET"  path:"/items/"         output:"[]Item"`
    addItem     gorest.EndPoint `method:"POST" path:"/item/"         postdata:"Item"`
    addItems    gorest.EndPoint `method:"POST" path:"/items/"         postdata:"[]Item"`

    //On a real app for placeOrder below, the POST URL would probably be just /orders/, this is just to
    // demo the ability of mixing post-data parameters with URL mapped parameters.
    placeOrder  gorest.EndPoint `method:"POST"   path:"/orders/new/{UserId:int}/{RequestDiscount:bool}" postdata:"Order"`
    viewOrder   gorest.EndPoint `method:"GET"    path:"/orders/{OrderId:int}"                           output:"Order"`
    deleteOrder gorest.EndPoint `method:"DELETE" path:"/orders/{OrderId:int}"`
    listOrders  gorest.EndPoint `method:"GET"  path:"/orders/"         output:"[]Order"`


}

//Handler Methods: Method names must be the same as in config, but exported (starts with uppercase)


func(serv OrderService) Discover() Discover {
	return Discover{User{},Item{},Order{}}
}

func(serv OrderService) UserDetails(Id int) (u User){
    if user,found:=userStore[Id];found{
        u =user
        return
    }
    serv.ResponseBuilder().SetResponseCode(404).Overide(true)  //Overide causes the entity returned by the method to be ignored. Other wise it would send back zeroed object
    return
}

func(serv OrderService) ListUsers()[]User{
    serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
	retval := make([]User,0)
	for _,v := range userStore {
		retval = append(retval,v)
	}
    return retval
}

func(serv OrderService) ListItems()[]Item{
    serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
    return itemStore
}

func(serv OrderService) ListOrders()[]Order{
    serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
    return orderStore
}

func(serv OrderService) _AddItem(i Item) Item {

    for _,item:=range itemStore{
        if item.Id == i.Id{
            item=i
            //serv.ResponseBuilder().SetResponseCode(200) //Updated http 200, or you could just return without setting this. 200 is the default for POST
            return item
        }
    }

    //Item Id not in database, so create new
    i.Id = len(itemStore)
    itemStore=append(itemStore,i)
    return i
}

func(serv OrderService) AddItem(i Item){
    itemAdded := serv._AddItem(i)
    serv.ResponseBuilder().Created("http://localhost:8787/orders-service/items/"+string(itemAdded.Id)) //Created, http 201
}

func(serv OrderService) AddItems(items []Item){
	idsAdded := make([]string,0)
	for _,item := range items {
		idsAdded = append( idsAdded, string( serv._AddItem( item ).Id) )
	}
	
   serv.ResponseBuilder().Created("http://localhost:8787/orders-service/items/"+strings.Join( idsAdded, ",") ) //Created, http 201
}

func findItem( itemId int ) (Item,bool) {
	return itemStore[itemId],true
}

//On the method parameters, the posted data(http-entity) is always first, followed by the URL mapped parameters
func(serv OrderService) PlaceOrder(order Order,UserId int,AskForDiscount bool){
    order.Id = len(orderStore)

    if user,found:= userStore[UserId];found{
          if item,exists:=findItem(order.ItemId);exists{
                itemStore[item.Id].AvailableStock--

                if AskForDiscount && order.Amount >5{
                    order.Discount = 2.5
                }
                order.Id=len(orderStore)
                order.UserId=UserId
                order.Cancelled=false
                orderStore=append(orderStore,order)
                user.OrderIds=append(user.OrderIds,order.Id)

                userStore[user.Id]=user

                serv.ResponseBuilder().SetResponseCode(201).Location("http://localhost:8787/orders-service/orders/"+string(order.Id))//Created
                return

          } else{
              serv.ResponseBuilder().SetResponseCode(404).WriteAndOveride([]byte("Item not found"))//You can still manually place an entity on the response, even on a POST
              return
          }
    }

    serv.ResponseBuilder().SetResponseCode(404).WriteAndOveride([]byte("User not found"))
    return
}
func(serv OrderService) ViewOrder(id int) (retOrder Order){
     for _,order:=range orderStore{
        if order.Id == id{
            retOrder = order
            return
        }
     }
     serv.ResponseBuilder().SetResponseCode(404).Overide(true)
     return
}
func(serv OrderService) DeleteOrder(id int) {
     for pos,order:=range orderStore{
        if order.Id == id{
            order.Cancelled =true
            orderStore[pos]=order
            return               //Default http code for DELETE is 200
        }
     }
     serv.ResponseBuilder().SetResponseCode(404).Overide(true)
     return
}

func init () {
	//itemStore = append(itemStore,Item{0 , 10})
	//itemStore = append(itemStore,Item{1 , 22})
	//itemStore = append(itemStore,Item{2 , 33})
	//itemStore = append(itemStore,Item{3 , 40})

	userStore = make(map[int]User)
	//userStore[0] = User{0,[]int{1}}
	//userStore[1] = User{1,[]int{2}}
	//userStore[2] = User{2,[]int{3}}
	//orderStore = append(orderStore,Order{1,0,2,5.00,1.00,false})

}
