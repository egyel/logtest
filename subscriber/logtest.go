package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	logtest "gitlab.com/krayohu/logtest/proto/logtest"
)

type Logtest struct{}

func (e *Logtest) Handle(ctx context.Context, msg *logtest.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *logtest.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
