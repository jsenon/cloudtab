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
docker run --name cloudtab-mongo -v /my/own/datadir:/data/db -d mongo:latest
./cloudtab 
```

### Access

Access through favorite web browser on http://YOURIP:YOURPORT ie http://127.0.0.1:8090

### ToDo





