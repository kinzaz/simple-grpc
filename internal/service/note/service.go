package note

import (
	"simple-grpc/internal/client/db"
	"simple-grpc/internal/repository"
	"simple-grpc/internal/service"
)

type serv struct {
	noteRepository repository.NoteRepository
	txManager      db.TxManager
}

func NewService(
	noteRepository repository.NoteRepository,
	txManager db.TxManager,
) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
		txManager:      txManager,
	}
}
