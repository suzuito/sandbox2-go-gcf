package crawler

import (
	"context"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
)

func CrawlerStartPipelinePeriodically(ctx context.Context, e event.Event) error {
	ctx = context.WithValue(ctx, "traceId", uuid.New().String())
	if err := u.StartPipelinePeriodically(ctx); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		return err
	}
	return nil
}
