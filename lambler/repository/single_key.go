package repository

type SingleKeyRepository[E any] interface {
	Create(factory func() E) error
}

type SingleKeyTable[E any] interface {
	Repository() SingleKeyRepository[E]
}
