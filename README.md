# cloudtab

Build Web Server with mongodb backend

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

Access through favorite web browser on http://YOURIP:YOURPORT ie http://127.0.0.1:8090

### ToDo

[x] API
[] Package for struct JSON Server
[x] Web part to view details
[] Web part Update
[] Import CSV
[] Detail Select Column to show





