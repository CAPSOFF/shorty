package shorty

import "context"

type Controller interface {
	Shorten(ctx context.Context, url string, desiredShortCode string) (shortCode string, errorCode int, err error)
	ShortCode(ctx context.Context) error
	ShortCodeStats(ctx context.Context) error
}
