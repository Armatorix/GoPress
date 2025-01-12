 PHONE: all
all:
	docker compose --profile local up --watch --build --force-recreate --renew-anon-volumes --remove-orphans

.PHONY: build
build:
	docker compose --profile local build --no-cache

.PHONY: api
api:
	cd api && make
	cd be && make gen
	cd web && bun run openapi

.PHONY: fresh
fresh:
	-docker compose -p gopress stop postgres
	docker compose -p gopress up -d --build --force-recreate --renew-anon-volumes --remove-orphans postgres
	docker compose -p gopress exec -it backend sh -c 'touch __tmp_reload_db.go && rm __tmp_reload_db.go'

