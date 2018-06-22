# mail
Very basic mail microservice using gomail.v2.
Once an EmailRequest hits the service, the request is stuffed into a buffered channel (max. 1000).
The queue is continuously polled from the Service to send out the emails.

There is NO GUARANTEE for the email to be sent. This is v1 and will surely change in the future. For the first iteration
that's totally fine.

### build
To compile the protobufs, link the binary and push it into an alpine container, simply use the provided Makefile and run:

```
make build
```

### run
The service uses Consul for discovery. The *docker-compose.local.yml* provides a local dev agent which is totally fine
for development.

Simply fire up the docker-compose: ```docker-compose -f docker-compose.local.yml up consul mail```