package daemon

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/rpolnx/go-asynq-poc/internal/configs"
	"github.com/sirupsen/logrus"
)

type Enqueuer struct {
	CachePool *asynq.Client
}

const (
	TypeEmailDelivery = "email:deliver"
	TypeImageResize   = "image:resize"
	TypeScheduledJob      = "job:scheduled"
)

type EmailDeliveryPayload struct {
	JobId        int
	EmailAddress string
	UserId       int
	CustomerId   int
}

type ImageResizePayload struct {
	JobId     int
	SourceURL string
}

func (d *Enqueuer) NewEmailDeliveryTask(idx int) (*asynq.Task, error) {
	logrus.Infof("Enqueued %s - %d", TypeEmailDelivery, idx)

	payload, err := json.Marshal(EmailDeliveryPayload{
		JobId:        idx,
		EmailAddress: fmt.Sprintf("test%d@example.com", idx),
		UserId:       100,
		CustomerId:   idx,
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func (d *Enqueuer) NewImageResizeTask(idx int, src string) (*asynq.Task, error) {
	logrus.Infof("Enqueued %s - %d", TypeImageResize, idx)

	payload, err := json.Marshal(ImageResizePayload{SourceURL: src, JobId: idx})
	if err != nil {
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.
		NewTask(TypeImageResize, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}

func (d *Enqueuer) NewScheduledJob(idx int, secondsToSchedule int) (*asynq.Task, error) {
	logrus.Infof("Enqueued %s, id: %d - executing in %ds ", TypeScheduledJob, idx, secondsToSchedule)

	payload, err := json.Marshal(EmailDeliveryPayload{
		JobId:        idx,
		EmailAddress: fmt.Sprintf("test%d@example.com", idx),
		UserId:       100,
		CustomerId:   idx,
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeScheduledJob, payload, asynq.MaxRetry(1),
		asynq.ProcessIn(time.Duration(secondsToSchedule)*time.Second)), nil
}

func (d *Enqueuer) EnqueueAll(tasks ...*asynq.Task) error {

	for _, t := range tasks {
		info, err := d.CachePool.Enqueue(t)
		if err != nil {
			return err
		}
		fmt.Println(info)
	}

	return nil
}

func NewEnqueuer(appConfig *configs.AppConfig, cachePool *asynq.Client) *Enqueuer {
	return &Enqueuer{
		CachePool: cachePool,
	}
}
