package postgres

import "todotom/go-gin-restful-api/forum"

type PostgresThreadRepository struct {
	postgresTable PostgresTable[*forum.Thread]
}

func ProvideThreadRepository(threadPostgresTable PostgresTable[*forum.Thread]) forum.ThreadRepository {
	return PostgresThreadRepository{threadPostgresTable}
}

func ProvideThreadPostgresTable(postgresClient PostgresClient) PostgresTable[*forum.Thread] {
	return ProvidePostgresTable[*forum.Thread]("thread", postgresClient)
}

func (repository PostgresThreadRepository) GetThreads() []*forum.Thread {
	thread, err := repository.postgresTable.GetOne("123")
	if err == nil {
		return []*forum.Thread{}
	}
	return []*forum.Thread{
		*thread,
	}
}
