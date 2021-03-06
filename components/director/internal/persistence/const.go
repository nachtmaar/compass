package persistence

import "github.com/lib/pq"

type persistenceCtxKey string

const (
	// PersistenceCtxKey is a key used in context to store the persistance object
	PersistenceCtxKey persistenceCtxKey = "PersistenceCtx"
	// UniqueViolation is the error code that happens when the Unique Key is violated
	UniqueViolation pq.ErrorCode = "23505"
)
