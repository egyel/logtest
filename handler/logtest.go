package handler

import (
	"context"

	//"github.com/micro/go-micro/v2/util/log"
	log "github.com/micro/go-micro/v2/logger"

	logtest "gitlab.com/krayohu/logtest/proto/logtest"
)

// Logtest is test
type Logtest struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Logtest) Call(ctx context.Context, req *logtest.Request, rsp *logtest.Response) error {
	log.Debug("Received Logtest.Call request")
	log.Info("Info in handler")
	log.Infof("log infof inside %q", req.Name)
	log.Error("Error in handler")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Logtest) Stream(ctx context.Context, req *logtest.StreamingRequest, stream logtest.Logtest_StreamStream) error {
	log.Debugf("Received Logtest.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Debugf("Responding: %d", i)
		if err := stream.Send(&logtest.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Logtest) PingPong(ctx context.Context, stream logtest.Logtest_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Debugf("Got ping %v", req.Stroke)
		if err := stream.Send(&logtest.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
