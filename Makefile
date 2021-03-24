tests:
	go test -v ./...

cover:
	go test -short -coverprofile=./coverage.out ./...

sec:
	gosec -fmt=sonarqube -out report.json ./...

sonarqube-scan:
	./sonar-scan.sh

sonarqube-server-up:
	docker run -d --name sonarqube -p 9777:9000 sonarqube

all: tests cover sec sonarqube-scan

sonarqube-server-down:
	docker stop sonarqube