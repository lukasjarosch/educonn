build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o mail .
	docker build -t educonn/mail-cli -f Dockerfile .


run:
	docker run -it \
	    -e MICRO_REGISTRY="mdns" \
	    educonn/mail-cli --to="hans@peter.com" --from="depp@depp.com" --subject="asdf" --body="wooooo"