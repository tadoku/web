package repositories

import (
	"github.com/srvc/fail"
	"github.com/tadoku/api/domain"
	"github.com/tadoku/api/interfaces/rdb"
	"github.com/tadoku/api/usecases"
)

// NewRankingRepository instantiates a new ranking repository
func NewRankingRepository(sqlHandler rdb.SQLHandler) usecases.RankingRepository {
	return &rankingRepository{sqlHandler: sqlHandler}
}

type rankingRepository struct {
	sqlHandler rdb.SQLHandler
}

func (r *rankingRepository) Store(ranking domain.Ranking) error {
	if ranking.ID == 0 {
		return r.create(ranking)
	}

	return r.update(ranking)
}

func (r *rankingRepository) create(ranking domain.Ranking) error {
	query := `
		insert into rankings
		(contest_id, user_id, language_code, amount, created_at, updated_at)
		values (:contest_id, :user_id, :language_code, :amount, :created_at, :updated_at)
	`

	_, err := r.sqlHandler.NamedExecute(query, ranking)
	return fail.Wrap(err)
}

func (r *rankingRepository) update(ranking domain.Ranking) error {
	query := `
		update rankings
		set amount = :amount, updated_at = NOW()
		where id = :id
	`

	_, err := r.sqlHandler.NamedExecute(query, ranking)
	return fail.Wrap(err)
}
