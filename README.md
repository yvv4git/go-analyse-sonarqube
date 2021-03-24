# go-analyse-sonarqube
Simple example static code analyse with sonar qube


### Check endpoint with console tool curl
```
curl -X GET 'http://localhost:7777/sum/5/3'
```

### Tests
```
go test -v
```

### Code coverage
```
go test -short -coverprofile=./coverage.out ./...
```

### Use gosec
```
go get github.com/securego/gosec/cmd/gosec
gosec -fmt=sonarqube -out report.json ./...
```

### Use sonarqube
```
docker run -d --name sonarqube -p 9777:9000 sonarqube
```