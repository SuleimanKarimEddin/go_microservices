.PHONY: run
run:
	docker-compose -f order/docker-compose.yml  up -d
	docker-compose -f  product/docker-compose.yml up -d
	docker-compose -f product_order/docker-compose.yml up -d

.PHONY: stop
stop:
	docker-compose -f order/docker-compose.yml -f product/docker-compose.yml -f product_order/docker-compose.yml down

.PHONY: destroy
destroy:
	docker-compose -f order/docker-compose.yml -f product/docker-compose.yml down --rmi all

.PHONY: install-migration
install-migration:
		curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
		echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
		apt-get update
		apt-get install -y migrate