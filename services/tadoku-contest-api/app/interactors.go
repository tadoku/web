package app

import (
	"time"

	"github.com/tadoku/api/infra"
	"github.com/tadoku/api/usecases"
)

// Interactors is a collection of all repositories
type Interactors struct {
	Session    usecases.SessionInteractor
	Contest    usecases.ContestInteractor
	ContestLog usecases.ContestLogInteractor
	Ranking    usecases.RankingInteractor
}

// NewInteractors initializes all repositories
func NewInteractors(
	r *Repositories,
	jwtGenerator usecases.JWTGenerator,
	sessionLength time.Duration,
) *Interactors {
	return &Interactors{
		Session: usecases.NewSessionInteractor(
			r.User,
			infra.NewPasswordHasher(),
			jwtGenerator,
			sessionLength,
		),
		Contest:    usecases.NewContestInteractor(r.Contest, infra.NewValidator()),
		ContestLog: usecases.NewContestLogInteractor(r.ContestLog, r.Contest, r.Ranking, infra.NewValidator()),
		Ranking:    usecases.NewRankingInteractor(r.Ranking, r.Contest, r.ContestLog, r.User, infra.NewValidator()),
	}
}
