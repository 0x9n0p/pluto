package pipeline

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"pluto"
	"pluto/panel/database"

	"go.uber.org/zap"
)

var bucket = []byte("pipeline")

func init() {
	transaction, err := database.Get().NewTransaction(true)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Create bucket (%s): %v\n", bucket, err)
		os.Exit(1)
	}

	defer func() {
		if err := transaction.CommitOrRollback(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Create bucket (%s): %v\n", bucket, err)
			os.Exit(1)
		}
	}()

	if _, err := transaction.CreateBucketIfNotExists(bucket); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Create bucket (%s): %v\n", bucket, err)
		os.Exit(1)
	}
}

func Find(tx *database.Transaction, name string) (a Pipeline, err error) {
	b := tx.Bucket(bucket).Get([]byte(name))
	if b == nil {
		return Pipeline{}, &pluto.Error{
			HTTPCode: http.StatusNotFound,
			Message:  fmt.Sprintf("Pipeline (%s) not found", name),
		}
	}

	if err := json.Unmarshal(b, &a); err != nil {
		pluto.Log.Error("Can not unmarshal pipeline", zap.String("pipeline_name", name), zap.Error(err))
		return Pipeline{}, &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "Failed to find pipeline",
		}
	}

	a.Transaction = tx
	return
}

func All(tx *database.Transaction) ([]Pipeline, error) {
	pipelines := []Pipeline{}
	c := tx.Bucket(bucket).Cursor()

	for k, v := c.First(); k != nil; k, v = c.Next() {
		var p Pipeline
		if err := json.Unmarshal(v, &p); err != nil {
			pluto.Log.Error("Can not unmarshal pipeline", zap.String("pipeline_name", string(k)), zap.Error(err))
			return nil, &pluto.Error{
				HTTPCode: http.StatusInternalServerError,
				Message:  "An internal server error occurred",
			}
		}

		p.Transaction = tx
		pipelines = append(pipelines, p)
	}

	return pipelines, nil
}

func ReloadExecutionCache() (err error) {
	tx, err := database.Get().NewTransaction(false)
	if err != nil {
		return err
	}

	defer func() {
		if err2 := tx.Rollback(); err2 != nil {
			err = err2
		}
	}()

	storedPipelines, err := All(tx)
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
