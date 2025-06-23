package log

import (
	"context"
	"time"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/clickhouse"
	"github.com/fivemanage/lite/pkg/kafkaqueue"
	"github.com/uptrace/bun"
)

type Service struct {
	db            *bun.DB
	kafkaProducer *kafkaqueue.Queue
}

func NewService(db *bun.DB, kafkaProducer *kafkaqueue.Queue, clickhouseClient *clickhouse.Client) *Service {
	return &Service{
		db:            db,
		kafkaProducer: kafkaProducer,
	}
}

func (r *Service) SubmitLogs(ctx context.Context, organizationId string, datasetID string, logs []api.Log) {
	clickhouseLogs := make([]*clickhouse.Log, len(logs))

	for i, log := range logs {
		timestamp := time.Now().UTC()
		metadata := make(map[string]string)

		for k, v := range log.Metadata {
			for key, value := range buildLogAttributes(k, v, 0) {
				metadata[key] = value
			}
		}

		// fix log message because it can be silly sometimes
		logMessage := formatMessage(log.Message)

		// these should override any metadata that matches
		metadata["message"] = logMessage
		metadata["severity"] = log.Level
		metadata["_resource"] = log.Resource

		clickhouseLogs[i] = &clickhouse.Log{
			TraceID:    "",
			Timestamp:  timestamp,
			TeamID:     organizationId,
			DatasetID:  datasetID,
			Body:       log.Message,
			Attributes: log.Metadata,
		}
	}

	r.kafkaProducer.Submit(ctx, clickhouseLogs)
}
