# cloudtab

Simple CMDB Web Server in GO with mongodb backend

### Prerequisite

You need to have:

* Go 1.8
* Go Environment properly set
* Mongodb

### Compilation

```sh
git clone https://github.com/jsenon/websocket.git
go build -o cloudtab
```

### Start

```sh
docker run --name cloudtab-mongo -v /my/own/datadir:/data/db -d -p 27017:27017 mongo:latest
./cloudtab 
```

### Access

Access through favorite web browser on http://YOURIP:YOURPORT ie http://127.0.0.1:9010

Access docker
```sh
docker exec -it cloudtab-mongo mongo admin
```

### API

POST API One Server
```sh
curl -X POST  -H "Content-Type: application/json"  -d @body_example.json http://localhost:9010/api/servers
```
POST API Multiple Servers
```sh
curl -X POST  -H "Content-Type: application/json"  -d @body_example.json http://localhost:9010/api/servers/import
```
GET API ALL
```sh
curl  -H "Content-Type: application/json"  http://localhost:9010/api/servers
```

GET API Specific ID
```sh
curl  -H "Content-Type: application/json"  http://localhost:9010/api/servers/YOURID
```

DELETE API 
```sh
curl -X DELETE -H "Content-Type: application/json"  http://localhost:9010/api/servers/YOURID
```
UPDATE API 
```sh
curl -X UPDATE -H  "Content-Type: application/json" -d @body_example_update.json http://localhost:9010/api/servers/YOURID
```

### ScreenShot

Main
![Alt text](/img/main.png?raw=true "Main")

Full Table View
![Alt text](/img/fullview.png?raw=true "FullView")

API Swagger Doc
![Alt text](/img/api_swagger.png?raw=true "API Swagger")



### ToDo

- [x] API
- [x] Package for struct JSON Server
- [x] Web part to view details
- [x] Web part Update
- [x] Web part add network func
- [x] Web part delete network func
- [x] API part Update
- [x] Import 
- [x] Detail Select Column to show
- [ ] Unique fields in db
- [x] Store multiple networks per asset
- [ ] Authentication
- [ ] Jwt implementation
- [x] Add date insertion asset
- [x] API Doc
- [ ] Build Docker App
- [ ] Move to microservice (dockerize cloudtab and add redis)





