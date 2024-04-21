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
