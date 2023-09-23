package repository

import (
	"context"
	"github.com/uptrace/bun"
)

type EnglishAnswerImpl struct {
	db *bun.DB
}

func NewEnglishAnswerRepository(db *bun.DB) *EnglishAnswerImpl {
	return &EnglishAnswerImpl{
		db: db,
	}
}

func (r *EnglishAnswerImpl) FindAnswerByQuestion(ctx context.Context, question string) (string, error) {
	answer := new(string)
	err := r.db.NewSelect().Model(answer).Where("question = ?", question).Scan(ctx)
	if err != nil {
		return "", err
	}
	return *answer, nil
}

// レコードを作成する
func (r *EnglishAnswerImpl) CreateAnswer(ctx context.Context, question string, answer string) error {
	// user がいるかどうかを確認する
	cnt, err := r.db.NewSelect().Model((*string)(nil)).Where("question = ?", question).Count(ctx)
	if err != nil {
		return err
	}
	if cnt > 0 {
		return nil
	}

	// レコードを作成する
	_, err = r.db.NewInsert().Model(&struct {
		Question string `bun:"question"`
		Answer   string `bun:"answer"`
	}{
		Question: question,
		Answer:   answer,
	}).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
