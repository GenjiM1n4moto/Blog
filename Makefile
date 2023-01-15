postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.1-alpine
start:
	docker start postgres15
stop:
	docker stop postgres15
createdb:
	docker exec -it postgres15 createdb --username=root --owner=root blog
dropdb:
	docker exec -it postgres15 dropdb blog
# createmigrate:
migrateup:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5432/blog?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5432/blog?sslmode=disable" -verbose down
sqlc:
	sqlc generate
.PHONY:	postgres start stop createdb dropdb migrateup migratedown sqlc