build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lukasjarosch/educonn/srv/course  --micro_out=. proto/course/*.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o course .
	docker build -t educonn/course-dev -f Dockerfile .

travis:
	$(HOME)/protobuf/bin/protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lukasjarosch/educonn/srv/course  --micro_out=. proto/course/*.proto
	dep ensure -v
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o course .

travis-deploy:
	../../bin/deploy-travis.sh "educonn-course"

run:
	go build -o course && ./course --debug=1 --db_host=localhost --db_port=3307 --db_user=course --db_pass=secret --db_name=courses