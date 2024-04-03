package processing

import "github.com/dico87/stori-challenge/internal/file_processing/domain"

type UseCase struct {
}

func NewProcessingUseCase() UseCase {
	return UseCase{}
}

func (u UseCase) Processing(transactions []domain.Transaction) {

}
