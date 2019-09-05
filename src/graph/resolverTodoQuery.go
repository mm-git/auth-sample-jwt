package graph

import (
	"context"
	"time"
)

func (r *queryResolver) TodoList(ctx context.Context) ([]*Todo, error) {
	// 本来はデータベースからtodoの一覧を取得する。
	// Query実行の際に、データベースインスタンスを渡していれば、
	// r.DB.TodoList() のようにしてDBにアクセスできる

	dummyTime, _ := time.Parse("2006-01-02", "2019-09-18")
	todos := []*Todo{
		{
			ID:       "00001",
			Title:    "大根を買う",
			Status:   "notYes",
			CreateAt: dummyTime,
		},
		{
			ID:       "00002",
			Title:    "郵便局に行く",
			Status:   "notYes",
			CreateAt: dummyTime,
		},
		{
			ID:       "00003",
			Title:    "机の整理をする",
			Status:   "done",
			CreateAt: dummyTime,
		},
	}

	return todos, nil
}
