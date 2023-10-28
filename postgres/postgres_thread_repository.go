package postgres

import "todotom/go-gin-restful-api/forum"

type PostgresThreadRepository struct {
	threads []forum.Thread
}

func ProvideThreadRepository() forum.ThreadRepository {
	return PostgresThreadRepository{
		threads: []forum.Thread{
			{
				Id:    "123",
				Title: "lol",
			},
		},
	}
}

func (repository PostgresThreadRepository) GetThreads() []forum.Thread {
	return repository.threads
}
