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
