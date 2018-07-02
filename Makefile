travis:
	cd srv/mail/ && $(MAKE) deploy && cd ../..
	cd srv/user/ && $(MAKE) deploy && cd ../..

