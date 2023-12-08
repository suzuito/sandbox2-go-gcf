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
	crawlerStarterSettingID := string(e.Data())
	if err := u.StartPipelinePeriodically(ctx, crawler.CrawlerStarterSettingID(crawlerStarterSettingID)); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		return err
	}
	return nil
}
