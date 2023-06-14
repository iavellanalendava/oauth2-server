package repository

import (
	"github.com/google/uuid"
	"go.uber.org/zap"

	"oauth2-server/internal/model"
)

type keysStore map[string]model.Key

type KeysStore struct {
	keysStore keysStore
	logger    *zap.Logger
}

func NewKeysStore(l *zap.Logger) KeysStore {
	keysStore := keysStore{}
	return KeysStore{
		keysStore: keysStore,
		logger:    l,
	}
}

func (r *KeysStore) Save(key *model.Key) error {
	if key != nil {
		r.keysStore[key.Id.String()] = *key
	}

	return nil
}

func (r *KeysStore) GetById(id uuid.UUID) (*model.Key, bool) {
	key, found := r.keysStore[id.String()]
	return &key, found
}

func (r *KeysStore) GetAll() []*model.Key {
	keys := make([]*model.Key, 0, len(r.keysStore))

	for _, key := range r.keysStore {
		keys = append(keys, &key)
	}

	return keys
}
