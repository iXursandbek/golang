# migration fayl yaratish
migrate create -ext sql -dir ./migrations/postgres -seq -digits 2 create_article_table

# bazaga qo'shish va ochirish
migrate -path ./migrations/postgres -database 
		'postgres://postgres:test1234@localhost:5432/new_service_db?sslmode=disable' up/down