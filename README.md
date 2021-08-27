# GO-PROJECT

This is a example of use of go in a backend project. I will use gin like backend framework and it will be generated a connection with a MongoDB

		To obtain gin framework
`go get -u github.com/gin-gonic/gin`

		To obtain mongodb driver in golang
`go get go.mongodb.org/mongo-driver/mongo`

		I will use Golang to my reading to my configuration file
`go get github.com/tkanos/gonfig`

		I will use gjson to parse json data of the response of market provider
`go get -u github.com/tidwall/gjson`

		Migrations will assume with golang-migrate tool
`brew install golang-migrate` To install in mac

## Description of the app

This app consist in query market information of 'alpaca' and save the data in a mongo database.

### v0.0.1

With one route the user say the backend: 'Hey! query some data of some market', and with other route the user can get the data in the database

# How to Init the project

First, change to executable the script of init_db.sh in the scripts folder:

`chmod +x scripts/init_db.sh`

Then execute it to initialize the database of the project and the migrations

`./scripts/init_db.sh`
