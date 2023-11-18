package crawler

import (
	"context"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
)

// MessagePublishedData contains the full Pub/Sub message
// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/eventarc/docs/cloudevents#pubsub
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type MessagePublishedData struct {
	Message struct {
		Data       []byte            `json:"data"`
		Attributes map[string]string `json:"attributes"`
	}
}

func initCloudEventFunction(ctx context.Context, e *event.Event) (*MessagePublishedData, error) {
	msg := MessagePublishedData{}
	if err := e.DataAs(&msg); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		return nil, err
	}
	traceID, exist := msg.Message.Attributes["traceId"]
	if exist {
		ctx = context.WithValue(ctx, "traceId", traceID)
	}
	return &msg, nil
}
