package kafkaqueue

import (
	"context"
	"fmt"
	"time"

	"github.com/fivemanage/lite/internal/clickhouse"
	"github.com/sirupsen/logrus"
)

const (
	DefaultBatchFlushSize      = 10000
	DefaultBatchedFlushTimeout = 3 * time.Second
)

type KafkaBatchWorker struct {
	consumer         *Queue
	logs             []*clickhouse.Log
	lastFlush        time.Time
	clickhouseClient *clickhouse.Client
}

func NewBatchWorker(clickhouseClient *clickhouse.Client, kafkaConsumer *Queue) *KafkaBatchWorker {
	return &KafkaBatchWorker{
		consumer:         kafkaConsumer,
		logs:             []*clickhouse.Log{},
		lastFlush:        time.Now(),
		clickhouseClient: clickhouseClient,
	}
}

func (r *KafkaBatchWorker) ProcessMessages() {
	for {
		func(ctx context.Context) {
			readCtx, readCancel := context.WithTimeout(ctx, DefaultBatchedFlushTimeout)
			defer readCancel()

			msg := r.consumer.ReadMessage(readCtx)
			if msg != nil {
				fmt.Println("Received message:", msg)
				r.logs = append(r.logs, msg)
			}

			if time.Since(r.lastFlush) > DefaultBatchedFlushTimeout || len(r.logs) > DefaultBatchFlushSize {
				err := r.clickhouseClient.BatchWriteLogRows(ctx, r.logs)
				if err != nil {
					fmt.Println("failed to batch write logs", err.Error())
				}

				r.lastFlush = time.Now()
				r.logs = []*clickhouse.Log{}
			}
		}(context.Background())
	}
}
