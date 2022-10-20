package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("SchemaGet", SchemaGet)
}

type Table struct {
	ProjectID string `json:"project_id"`
	DatasetID string `json:"dataset_id"`
	TableID   string `json:"table_id"`
}

type Resp struct {
	TableInfo string                  `json:"table_info"`
	ColumnNum int                     `json:"column_num"`
	Schemas   []*bigquery.FieldSchema `json:"schemas"`
}

func SchemaGet(w http.ResponseWriter, r *http.Request) {
	var table Table
	if err := json.NewDecoder(r.Body).Decode(&table); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, table.ProjectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer client.Close()

	// fetch the table metadata and print the schema info
	meta, err := client.Dataset(table.DatasetID).Table(table.TableID).Metadata(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var schema Resp
	schema.TableInfo = fmt.Sprintf("%s:%s.%s", table.ProjectID, table.DatasetID, table.TableID)
	schema.ColumnNum = len(meta.Schema)
	schema.Schemas = meta.Schema
	// var scs []byte
	// for _, field := range meta.Schema {
	// 	data, err := json.Marshal(field)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}
	// 	scs = append(scs, data...)
	// }
	// schema.Schemas = string(scs)
	json.NewEncoder(w).Encode(schema)
}
