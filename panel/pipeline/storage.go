package pipeline

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"pluto"
	"time"

	"github.com/dgraph-io/badger/v4"
	"go.uber.org/zap"
)

var storage *Storage

func GetStorage() *Storage {
	return storage
}

func init() {
	db, err := badger.Open(badger.DefaultOptions(Env.PipelinesPath))
	if err != nil {
		pluto.Log.Fatal("Open storage of panel", zap.Error(err))
	}
	storage = &Storage{db}
}

type Storage struct {
	*badger.DB
}

func (s *Storage) Find(name string) (p Pipeline, err error) {
	return p, s.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(name))
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				return &pluto.Error{
					HTTPCode: http.StatusNotFound,
					Message:  fmt.Sprintf("Pipeline (%s) not found", name),
				}
			}

			pluto.Log.Error("Get pipeline", zap.String("pipeline_name", name), zap.Error(err))
			return fmt.Errorf("get pipeline: %v", err)
		}

		return item.Value(func(val []byte) error {
			if err := json.Unmarshal(val, &p); err != nil {
				pluto.Log.Error("Unmarshal pipeline", zap.String("pipeline_name", name), zap.Error(err))
				return fmt.Errorf("unmarshal pipeline: %v", err)
			}
			return nil
		})
	})
}

func (s *Storage) save(pipeline *Pipeline) error {
	pipeline.SavedAt = time.Now()

	b, err := json.Marshal(pipeline)
	if err != nil {
		return fmt.Errorf("marshal pipeline: %v", err)
	}

	return s.Update(func(txn *badger.Txn) error {
		if err := txn.Set([]byte(pipeline.Name), b); err != nil {
			return fmt.Errorf("set pipeline: %v", err)
		}
		return nil
	})
}

func (s *Storage) delete(name string) error {
	return s.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(name))
		if err != nil {
			return fmt.Errorf("delete pipeline: %v", err)
		}
		return nil
	})
}

func (s *Storage) List() (pipelines []Pipeline, err error) {
	pipelines = []Pipeline{}
	return pipelines, s.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			if err := it.Item().Value(func(v []byte) error {
				var p Pipeline
				// TODO: move it out of transaction
				if err := json.Unmarshal(v, &p); err != nil {
					pluto.Log.Error("Unmarshal pipeline", zap.Error(err))
					return fmt.Errorf("unmarshal pipeline: %v", err)
				}
				pipelines = append(pipelines, p)
				return nil
			}); err != nil {
				pluto.Log.Error("Retrieve pipeline", zap.Error(err))
				return fmt.Errorf("retrieve pipeline: %v", err)
			}
		}

		return nil
	})
}

func (s *Storage) ReloadExecutionCache() error {
	storedPipelines, err := GetStorage().List()
	if err != nil {
		return fmt.Errorf("list pipelines: %v", err)
	}

	pipelines := make(map[string]pluto.Pipeline)
	for _, storedPipeline := range storedPipelines {
		pipeline, err := storedPipeline.Create()
		if err != nil {
			return err
		}

		pipelines[pipeline.Name] = pipeline
	}

	pluto.ReloadExecutionCache(pipelines)
	return nil
}
