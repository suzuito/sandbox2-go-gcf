package crawler

import (
	"context"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func CrawlerStartPipelinePeriodically(ctx context.Context, e event.Event) error {
	ctx = context.WithValue(ctx, "traceId", uuid.New().String())
	msg, err := initCloudEventFunction(ctx, &e)
	if err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		return err
	}
	if err := u.StartPipelinePeriodically(ctx, crawler.CrawlerStarterSettingID(msg.Message.Data)); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		return err
	}
	return nil
}
