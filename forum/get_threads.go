package forum

type GetThreadsService struct {
	threadRepository ThreadRepository
}

func ProvideGetThreadsService(threadRepository ThreadRepository) GetThreadsService {
	return GetThreadsService{threadRepository}
}

func (service GetThreadsService) Run() []*Thread {
	return service.threadRepository.GetThreads()
}
