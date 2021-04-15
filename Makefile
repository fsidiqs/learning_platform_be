migrateup:
	./migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/learning_platform?parseTime=True" -verbose up

migratedown: 
	./migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/learning_platform?parseTime=True" -verbose down

migratenew:
	./migrate create -ext sql -dir db/migrations -seq


# If I try to migrate again, migrate will STILL refuse:

# error: Dirty database version 16. Fix and force version.
# So I have to go the last successfull version, which is 15.

migrateforce:
	./migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/learning_platform?parseTime=True" force 15


.PHONY: migrateup migratedown