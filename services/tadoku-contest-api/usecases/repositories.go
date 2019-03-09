//go:generate gex mockgen -source=repositories.go -package usecases -destination=repositories_mock.go

package usecases

import (
	"github.com/tadoku/api/domain"
)

// UserRepository handles User related database interactions
type UserRepository interface {
	Store(user domain.User) error
	FindByID(id uint64) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
}

// ContestRepository handles Contest related database interactions
type ContestRepository interface {
	Store(contest domain.Contest) error
	GetOpenContests() ([]uint64, error)
}

// RankingRepository handles Ranking related database interactions
type RankingRepository interface {
	Store(contest domain.Ranking) error
}
