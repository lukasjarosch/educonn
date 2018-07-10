travis:
	cd srv/mail/ && $(MAKE) travis && cd ../..
	cd srv/user/ && $(MAKE) travis && cd ../..
	cd srv/course/ && $(MAKE) travis && cd ../..

travis-deploy:
	cd srv/mail/ && $(MAKE) travis-deploy && cd ../..
	cd srv/user/ && $(MAKE) travis-deploy && cd ../..
	cd srv/course/ && $(MAKE) travis-deploy && cd ../..
