package crawler

import (
	"context"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
)

func CrawlerDispatchCrawl(ctx context.Context, e event.Event) error {
	msg, err := initCloudEventFunction(ctx, &e)
	if err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		return err
	}
	if err := u.DispatchCrawlOnGCF(ctx, msg.Message.Data); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		return err
	}
	return nil
}
