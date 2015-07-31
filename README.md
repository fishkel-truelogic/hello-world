Mongo install
-------------
* http://docs.mongodb.org/manual/tutorial/install-mongodb-on-ubuntu/

Go install
----------
* sudo apt-get install golang
* mkdir ~/go
* export GOPATH=~/go
* mkdir ~/go/src


Project setup
-------------
* cd ~/go/src
* git clone git@github.com:fishkel-truelogic/hello-world
* export MONGOHQ_URL=localhost:27017
* go get github.com/gorilla/mux
* go get gopkg.in/mgo.v2

Project Build and Run
---------------------
* cd ~/go/src/hello-world
* go run Main.go

Try
---
* http://localhost:8080/