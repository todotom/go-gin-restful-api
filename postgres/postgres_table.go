package postgres

import "log"

type PostgresTable[T any] struct {
	name           string
	postgresClient PostgresClient
}

func ProvidePostgresTable[T any](tableName string, postgresClient PostgresClient) PostgresTable[T] {
	return PostgresTable[T]{tableName, postgresClient}
}

func (table PostgresTable[T]) GetOne(id string) (*T, error) {
	var result T
	// parameters := map[string]interface{}{"table": table.name, "id": id}
	// rows, err := table.postgresClient.database.NamedQuery("SELECT id, title FROM :table WHERE id=:id;", parameters)
	rows, err := table.postgresClient.database.Queryx("SELECT id, title FROM thread WHERE id='123';")
	if err != nil {

	}
	// defer rows.Close()
	// if !rows.Next() {
	// 	return nil, rows.Err()
	// }
	// rows.Next()
	// rows.StructScan(result)
	// var thread Thread
	for rows.Next() {
		err = rows.StructScan(&result)
	}
	if err != nil {
		log.Fatalln(err)
	}
	return &result, nil
}
