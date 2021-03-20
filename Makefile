migrateup:
	./migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/learning_platform?parseTime=True" -verbose up

migratedown: 
	./migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/learning_platform?parseTime=True" -verbose down

migratenew:
	./migrate create -ext sql -dir db/migrations -seq

.PHONY: migrateup migratedown