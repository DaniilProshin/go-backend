package n4j

import "github.com/go-backend/internal/storage"

type Storage struct {
	//neo4j driver
	userRepository  *UserRepository
	offerRepository *OfferRepository
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) User() storage.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			storage: s,
		}
	}
	return s.userRepository
}

func (s *Storage) Offer() storage.OfferRepository {
	if s.offerRepository == nil {
		s.offerRepository = &OfferRepository{
			storage: s,
		}
	}
	return s.offerRepository
}
