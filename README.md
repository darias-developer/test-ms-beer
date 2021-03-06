# Test MS Beer v100


### **use mod**
```
go mod init
go mod tidy
```

### **config**
```
actualizar datos en el archivo .env
```

### **instalar**
```
go get github.com/joho/godotenv
go get go.mongodb.org/mongo-driver/bson
go get go.mongodb.org/mongo-driver/bson/primitive
go get github.com/gorilla/mux
go get github.com/rs/cors
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```

### **para ejecutar**
```
go clean & go build main.go
go clean & go run main.go
```

### **para ejecutar test**
```
go clean -testcache & go test ./... -cover
go test ./... -coverprofile coverage_out
go tool cover -func coverage_out
go tool cover -html=coverage_out
```

### **para probar api**
```
postman/test-ms-beer.postman_collection.json
```

### **para probar servicio api.currencylayer.com**
```
postman/api.currencylayer.com.postman_collection.json
```

### **docker images**
```
docker pull golang
docker pull mongo
```

### **docker logs config**
```
docker volume create logs
docker volume create mongo_data
```

### **construir imagen*
```
docker build --tag test-ms-beer:v1.0 .

## cambiar tag
docker image tag test-ms-beer:latest test-ms-beer:v1.0
docker image rm test-ms-beer:latest

docker image ls
```

### **correr imagen*
```
export DB_SOURCE=mongodb://localhost:27017
export PORT=8080
export LOG_PATH=C:\workspace\test-ms-beer.log
export ACCESS_KEY=e59979e596dc86b3aaea9f1727e41416
docker run --env DB_SOURCE --publish 8080:8080 -v logs:/app/logs/test-ms-beer.log test-ms-beer:v1.0 &
```

### **revisar imagen*
```
docker ps -a
docker rm <CONTAINER ID>

#quitar todos los containers
docker rm $(docker ps -a -q)

docker exec -it <container_name> bash
```

### **docker-compose*
```
docker-compose up -d

logs: 
docker exec -it go-app bash

tail -n 500 -f logs/

```