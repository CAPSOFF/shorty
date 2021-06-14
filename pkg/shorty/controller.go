package shorty

import "context"

type Controller interface {
	Shorten(ctx context.Context) error
	ShortCode(ctx context.Context) error
	ShortCodeStats(ctx context.Context) error
}
