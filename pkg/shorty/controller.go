package shorty

import (
	"context"

	"github.com/amartha-shorty/pkg/model"
)

type Controller interface {
	Shorten(ctx context.Context, url string, desiredShortCode string) (shortCode string, httpStatusCode int, err error)
	ShortCode(ctx context.Context, shortCode string) (url string, httpStatusCode int, err error)
	ShortCodeStats(ctx context.Context, shortCode string) (shortyData model.ShortySpec, httpStatusCode int, err error)
}
