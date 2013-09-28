#!/bin/sh


## list items
echo "List items"
GET  http://localhost:8787/orders-service/items
echo

## list users
echo "List users"
GET  http://localhost:8787/orders-service/users
echo

## list orders
echo "List orders"
GET  http://localhost:8787/orders-service/orders
echo


## add an item
echo "Add item"
#echo '{"Id":4,"AvailableStock":55}' | POST http://localhost:8787/orders-service/items
echo '{"AvailableStock":35}' | POST http://localhost:8787/orders-service/items
echo

echo "list items again, test that Id=4 was just added"
GET  http://localhost:8787/orders-service/items
echo


#POST http://localhost:8787/orders-service/orders/new/3/true
#POST http://localhost:8787/orders-service/orders/new/2/true
#POST http://localhost:8787/orders-service/orders/new/\/true
#POST http://localhost:8787/orders-service/orders/new/\{UserId:2\}/true
#POST http://localhost:8787/orders-service/orders/new/2/true
#POST http://localhost:8787/orders-service/orders/
#GET http://localhost:8787/orders-service/orders/2
#GET http://localhost:8787/orders-service/orders/1
#GET http://localhost:8787/orders-service/orders/0
#GET http://localhost:8787/orders-service/orders
#POST http://localhost:8787/orders-service/orders/new/2/true
#GET http://localhost:8787/orders-service/orders
#GET http://localhost:8787/orders-service/orders/1
#POST http://localhost:8787/orders-service/orders/new/2/true
#GET http://localhost:8787/orders-service/orders/1
#GET http://localhost:8787/orders-service/orders/
#GET http://localhost:8787/orders-service/discover

