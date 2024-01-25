package extension

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pluto"
	"pluto/extensions"
	"pluto/panel/database"
	"time"

	"go.uber.org/zap"
)

type Extension struct {
	Descriptor         `json:"descriptor"`
	InstallationDetail `json:"installation_detail"`
	Transaction        *database.Transaction `json:"-"`
}

func (e *Extension) Install() error {
	bucket := e.Transaction.Bucket(bucket_InstalledExtensions)

	if bucket.Get([]byte(e.ID)) != nil {
		return &pluto.Error{
			HTTPCode: http.StatusConflict,
			Message:  fmt.Sprintf("Extension (%s) already installed", e.ID),
		}
	}

	ext, found := extensions.Extensions[e.ID]
	if !found {
		return &pluto.Error{
			HTTPCode: http.StatusNotFound,
			Message:  fmt.Sprintf("Extension (%s) not found", e.ID),
		}
	}

	if err := ext.Install(); err != nil {
		pluto.Log.Debug("Failed to install extension", zap.String("extension_id", e.ID), zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  fmt.Sprintf("Failed to install extension (%s)", e.ID),
		}
	}

	{
		e.InstallationDetail = InstallationDetail{
			Installed:   true,
			InstalledAt: time.Now(),
		}

		b, err := json.Marshal(e.InstallationDetail)
		if err != nil {
			pluto.Log.Error("Failed to marshal installation detail", zap.Error(err))
			return &pluto.Error{
				HTTPCode: http.StatusInternalServerError,
				Message:  fmt.Sprintf("Failed to install extension (%s)", e.ID),
			}
		}

		if err := bucket.Put([]byte(e.ID), b); err != nil {
			pluto.Log.Error("Failed to put the installation detail on the bucket", zap.Error(err))
			return &pluto.Error{
				HTTPCode: http.StatusInternalServerError,
				Message:  fmt.Sprintf("Failed to install extension (%s)", e.ID),
			}
		}
	}

	return nil
}

func (e *Extension) Uninstall() error {
	bucket := e.Transaction.Bucket(bucket_InstalledExtensions)

	v := bucket.Get([]byte(e.ID))
	if v == nil {
		return &pluto.Error{
			HTTPCode: http.StatusNotFound,
			Message:  fmt.Sprintf("Extension (%s) is not installed", e.ID),
		}
	}

	if err := json.Unmarshal(v, &e.InstallationDetail); err != nil {
		pluto.Log.Error("Failed to marshall installation detail", zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  "An internal server occurred",
		}
	}

	ext, found := extensions.Extensions[e.ID]
	if !found {
		return &pluto.Error{
			HTTPCode: http.StatusNotFound,
			Message:  fmt.Sprintf("Extension (%s) not found", e.ID),
		}
	}

	if err := ext.Uninstall(); err != nil {
		pluto.Log.Debug("Failed to uninstall extension", zap.String("extension_id", e.ID), zap.Error(err))
		return &pluto.Error{
			HTTPCode: http.StatusInternalServerError,
			Message:  fmt.Sprintf("Failed to uninstall extension (%s)", e.ID),
		}
	}

	e.Installed = false
	return nil
}

type Descriptor struct {
	// ID such as extension_name-v1.2.3
	ID   string `json:"id"`
	Name string `json:"name"`

	/*
		Below pipelines and processors are added/deleted during the installation/uninstallation process.
	*/

	Processors []string `json:"processors"`
	Pipelines  []string `json:"pipelines"`
}

type InstallationDetail struct {
	Installed   bool      `json:"installed"`
	InstalledAt time.Time `json:"installed_at"`
}
