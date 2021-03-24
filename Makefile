tests:
	go test -v ./...

cover:
	go test -short -coverprofile=./coverage.out ./...

sec:
	gosec -fmt=sonarqube -out report.json ./...

sonarqube-scan:
	./sonar-scan.sh