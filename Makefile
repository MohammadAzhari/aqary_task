rundb:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=0000 -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root aqary_task

migrateup:
	migrate -path db/migration -database "postgresql://root:0000@localhost:5432/aqary_task?sslmode=disable" up

migratedown:
	migrate -path db/migration -database "postgresql://root:0000@localhost:5432/aqary_task?sslmode=disable" down