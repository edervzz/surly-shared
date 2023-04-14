package abstractions

import "context"

type ICreateRepository[TEntity any] interface {
	Create(context.Context, *TEntity) error
}

type IUpdateRepository[TEntity any] interface {
	Update(context.Context, *TEntity) error
}

type IReadSingleRepository[TKey any, TEntity any] interface {
	// Read single entity
	// 	context.Context: current Context
	// 	bool: by pass buffer (true)
	// 	TKey: type key
	Read(context.Context, TKey) (*TEntity, error)
}

type IReadByExternalIDRepository[TExtID any, TEntity any] interface {
	// Read single entity by external ID
	// 	context.Context: current Context
	// 	bool: by pass buffer (true)
	// 	TKey: type key
	ReadByExternalID(context.Context, TExtID) (*TEntity, error)
}

type IReadByParentIDRepository[TParentKey any, TEntity any] interface {
	// Read entities by parent key
	// 	context.Context: current Context
	// 	bool: by pass buffer (true)
	// 	TKey: type key
	ReadByParentID(context.Context, TParentKey) ([]TEntity, error)
}
