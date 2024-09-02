postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

gentool_specify_table:
    @echo "Generating Model with table: $(table_name)"
    gentool -db=postgres -dsn "host=127.0.0.1 user=root password=secret dbname=root port=5432" -tables "$(table_name)" -onlyModel

gentool_all_table:
	@echo "Generating Model with all table"
	gentool -db=postgres -dsn "host=127.0.0.1 user=root password=secret dbname=root port=5432" -onlyModel

