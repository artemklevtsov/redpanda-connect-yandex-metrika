package stat_table

import (
	"github.com/artemklevtsov/redpanda-connect-yandex-metrika/internal/misc"
	"github.com/redpanda-data/benthos/v4/public/service"
)

type apiResponse struct {
	Query     *apiQuery          `json:"query"`
	Data      []apiResponseEntry `json:"data"`
	TotalRows int                `json:"total_rows"`
}

type apiResponseEntry struct {
	Dimensions []struct {
		Name string `json:"name"`
		Id   string `json:"id,omitempty"`
	} `json:"dimensions"`
	Metrics []float64 `json:"metrics"`
}

func (r *apiResponse) Batch(formatKeys bool) (service.MessageBatch, error) {
	msgs := make(service.MessageBatch, len(r.Data))

	for i, e := range r.Data {
		row := make(map[string]any)

		for di, d := range e.Dimensions {
			k := r.Query.Dimensions[di]
			if formatKeys {
				k = misc.FormatKey(k)
			}

			row[k] = d.Name
		}

		for mi, m := range e.Metrics {
			k := r.Query.Metrics[mi]
			if formatKeys {
				k = misc.FormatKey(k)
			}

			row[k] = m
		}

		msg := service.NewMessage(nil)
		msg.SetStructuredMut(row)
		msg.MetaSetMut("query", r.Query)
		msg.MetaSetMut("total_rows", r.TotalRows)

		msgs[i] = msg
	}

	return msgs, nil
}
