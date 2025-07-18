run: 
	go run cmd/http/main.go

# init.up and init.down are used to start postgres and adminer services
init.up:
	docker compose -f docker-compose.yml up -d

init.down:
	docker compose -f docker-compose.yml down

# pg.up and pg.down are used to manage the PostgreSQL service
pg.up:
	docker compose -f docker-compose.yml up -d postgres

pg.down:
	docker compose -f docker-compose.yml down postgres

# redis.up and redis.down are used to manage the Redis service
redis.up:
	docker compose -f docker-compose.yml up -d redis

redis.down:
	docker compose -f docker-compose.yml down redis

# adminer.up and adminer.down are used to manage the Adminer service
adminer.up:
	docker compose -f docker-compose.yml up -d adminer

adminer.down:
	docker compose -f docker-compose.yml down adminer

# migrate.up and migrate.down are used to run database migrations
migrate.up:
	docker compose -f docker-compose.yml --profile migration run --rm migrate-up

migrate.down:
	docker compose -f docker-compose.yml --profile migration run --rm migrate-down