package crawler

import (
	"context"

	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/crawler/crawler/pkg/inject"
	"github.com/suzuito/sandbox2-go/crawler/crawler/pkg/usecase"
)

var u usecase.Usecase

func init() {
	clog.L.AddKey("traceId")
	clog.L.AddKey("crawlerId")
	ctx := context.Background()
	var err error
	u, err = inject.NewUsecaseGCP(ctx)
	if err != nil {
		clog.L.Errorf(ctx, "%+v", err)
	}
}
