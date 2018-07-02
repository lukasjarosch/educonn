travis:
	cd srv/mail/ && $(MAKE) travis && cd ../..
	cd srv/user/ && $(MAKE) travis && cd ../..

