package repository

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CredentialsStore struct {
	memoryStore memoryStore
	logger      *zap.Logger
	engine      *gin.Engine
}

func NewCredentialsStore(e *gin.Engine, l *zap.Logger) CredentialsStore {
	// This can be changed to a database to store all the credentials that the server handles.
	// For this code-challenge, it's just used the following pair username-password as example.

	credentials := memoryStore{
		"user-test": {
			Username: "user-test",
			Password: "password-test",
		},
	}

	return CredentialsStore{
		memoryStore: credentials,
		logger:      l,
		engine:      e,
	}
}
