# go-analyse-sonarqube
Simple example static code analyse with sonar qube


### Check endpoint with console tool curl
```
curl -X GET 'http://localhost:7777/sum/5/3'
```

### Tests
```
go test -v ./...
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
./sonar-scan.sh
```

### Use make file
```
make cover
make sec
make sonarqube-scan

or

make all
```

### Stop sonarqube server
```
make sonarqube-server-down
```

### Screenshots
![sonarqube in work #1](https://github.com/yvv4git/go-analyse-sonarqube/blob/master/sonarqube_scan.png)  
![sonarqube in work #2](https://github.com/yvv4git/go-analyse-sonarqube/blob/master/sonarqube_scan2.png)