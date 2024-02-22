package daemon

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rpolnx/go-asynq-poc/internal/configs"
	"github.com/sirupsen/logrus"
)

type Processor struct {
	ServerCache *asynq.Server
}

func (c *Processor) HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	logrus.Infof("Sending Email to User: user_email=%s, job_id=%d", p.EmailAddress, p.JobId)
	// Email delivery code ...
	return nil
}

func (c *Processor) HandleImageResizeTask(ctx context.Context, t *asynq.Task) error {
	var p ImageResizePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	logrus.Infof("Resizing image: source_url=%d, job_id=%d", p.SourceURL, p.JobId)
	// Resizing image code ...
	return nil
}

func (c *Processor) HandleScheduledJob(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	logrus.Infof("Sending Email to User: user_id=%d, job_id=%d", p.UserId, p.JobId)
	// Scheduled Job code ...
	return nil
}

func NewProcessor(appConfig *configs.AppConfig, serverCache *asynq.Server) *Processor {
	return &Processor{
		ServerCache: serverCache,
	}
}
