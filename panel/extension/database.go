package extension

import (
	"encoding/json"
	"pluto"
	ext_http "pluto/extensions/http"
	"pluto/panel/database"

	"go.uber.org/zap"
)

var bucket_InstalledExtensions = []byte("installed_extensions")

var Descriptors = []Descriptor{
	{
		ID:   ext_http.ExtensionID_V1,
		Name: "Restful",
		Processors: []string{
			ext_http.ProcessorName_WriteResponse,
		},
		Pipelines: []string{},
	},
}

func FindDescriptor(id string) (Descriptor, bool) {
	for _, descriptor := range Descriptors {
		if descriptor.ID == id {
			return descriptor, true
		}
	}
	return Descriptor{}, false
}

func FindInstallationDetail(tx *database.Transaction, id string) (InstallationDetail, bool) {
	b := tx.Bucket(bucket_InstalledExtensions).Get([]byte(id))
	if b == nil {
		return InstallationDetail{}, false
	}

	var d InstallationDetail
	if err := json.Unmarshal(b, &d); err != nil {
		pluto.Log.Error("Failed to unmarshal installation detail", zap.Error(err))
		return InstallationDetail{}, false
	}

	return d, true
}
