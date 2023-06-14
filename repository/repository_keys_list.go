package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"oauth2-server/internal/model"
)

type keysStore map[string][]model.Key

type KeysStore struct {
	keysStore keysStore
	logger    *zap.Logger
}

func NewKeysStore(l *zap.Logger) KeysStore {
	keysStore := keysStore{}

	keyTest := model.Key{
		Id:        uuid.New(),
		PublicKey: []byte("example-public-key"),
		CreatedAt: time.Now(),
	}
	keysStore["client-id-example"] = []model.Key{
		keyTest,
	}

	return KeysStore{
		keysStore: keysStore,
		logger:    l,
	}
}

func (r *KeysStore) Save(clientId string, key *model.Key) error {
	if _, found := r.keysStore[clientId]; !found {
		if key != nil {
			r.keysStore[clientId] = []model.Key{
				*key,
			}
		}
		return nil
	}

	r.keysStore[clientId] = append(r.keysStore[clientId], *key)
	return nil
}

func (r *KeysStore) GetAllById(clientId string) ([]*model.Key, error) {
	if _, found := r.keysStore[clientId]; !found {
		return nil, fmt.Errorf("no keys were found for cliendId provided")
	}

	var keys []*model.Key
	for _, key := range r.keysStore[clientId] {
		keys = append(keys, &key)
	}

	return keys, nil
}
