package gcfcrawler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/functions/metadata"
	"github.com/google/uuid"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	CreateTime time.Time `json:"createTime"`
	// Fields is the data for this value. The type depends on the format of your
	// database. Log the interface{} value and inspect the result to see a JSON
	// representation of your database fields.
	Fields     interface{} `json:"fields"`
	Name       string      `json:"name"`
	UpdateTime time.Time   `json:"updateTime"`
}

func NotifierNotify(ctx context.Context, e FirestoreEvent) error {
	ctx = context.WithValue(ctx, "traceId", uuid.New().String())
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %w", err)
	}
	// clog.L.Infof(ctx, "Function triggered by change to: %v", meta.Resource)
	// clog.L.Infof(ctx, "Old value: %+v", e.OldValue)
	// clog.L.Infof(ctx, "New value: %+v", e.Value)
	fullPath := strings.Split(e.Value.Name, "/documents/")[1]
	if err := u.NotifyOnGCF(ctx, fullPath); err != nil {
		// return terrors.Wrap(err)
		return terrors.Wrap(err)
	}
	return nil
}
