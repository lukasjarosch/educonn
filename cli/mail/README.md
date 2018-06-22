# mail-cli

gRPC client for the mail service

### build

```
make build
```

### run

Edit parameters in Makefile and then just run

```
make run
```

Or run it directly with:
```
	docker run -it -e MICRO_REGISTRY="consul" educonn/mail-cli --to="hans@peter.com" --from="dein@portal.com" --subject="Irgendwas" --body="Noch mehr Nonsens"
```
