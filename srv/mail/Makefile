build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lukasjarosch/educonn/srv/mail --micro_out=.  proto/mail.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o mail .
	docker build -t educonn/mail -f Dockerfile .


run:
	docker run -it \
	    -p 50051:50051 \
	    -e MICRO_ADDRESS=":50051" \
	    -e MICRO_REGISTRY="mdns" \
	    -e SMTP_HOST=smtp.mailtrap.io \
	    -e SMTP_PORT=2525 \
	    -e SMTP_USERNAME=0f92ba981e4a2f \
	    -e SMTP_PASSWORD=09f0dcb79f8f24 \
	    educonn/mail