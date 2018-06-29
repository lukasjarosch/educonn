build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o user .
	docker build -t educonn/user-cli -f Dockerfile .

clean:
	rm user


run:
	docker run -it \
	    -e MICRO_REGISTRY="consul" \
	    educonn/mail-cli --to="hans@peter.com" --from="dein@portal.com" --subject="Irgendwas" --body="Noch mehr Nonsens"
