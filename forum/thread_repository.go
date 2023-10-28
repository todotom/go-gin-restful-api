package forum

type ThreadRepository interface {
	GetThreads() []Thread
}
