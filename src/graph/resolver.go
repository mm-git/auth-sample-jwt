// 以下を実行すると、 graph/schema/graphqlに追加した query/mutationが generated.go/modes_gen.goに追加される。
// generated.go の QueryResolver interface 及び MutationResolver interface に追加された関数を実装すると、
// APIでgraphQLの応答を返すことができる。
//go:generate go run ../../../scripts/gqlgen.go

package graph

type Resolver struct {
	//　ここに GraphQLの処理に必要なデータ(アカウント名、データベースインスタンス)を定義する
	// UserId string
	// DB     *database.DB
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }
