build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lukasjarosch/educonn/srv/user  --micro_out=. proto/user/*.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o user .
	docker build -t educonn/user-dev -f Dockerfile .

travis:
	$(HOME)/protobuf/bin/protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lukasjarosch/educonn/srv/user --micro_out=.  proto/user/*.proto
	dep ensure -v
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o user .

travis-deploy:
	../../bin/deploy-travis.sh "educonn-user"

clean:
	rm user


