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
	msg := struct {
		CrawlerStarterSettingID crawler.CrawlerStarterSettingID
	}{}
	if err := e.DataAs(&msg); err != nil {
		return err
	}
	if err := u.StartPipelinePeriodically(ctx, msg.CrawlerStarterSettingID); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		return err
	}
	return nil
}
