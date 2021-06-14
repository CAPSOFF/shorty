package controller

import (
	"context"
	"time"

	"github.com/amartha-shorty/pkg/shorty"
)

type shortenController struct {
	timeout           time.Duration
	shortenRepository shorty.Repository
}

func NewShortyController(timeout time.Duration, shortenRepository shorty.Repository) shorty.Controller {
	return &shortenController{
		timeout:           timeout,
		shortenRepository: shortenRepository,
	}
}

func (uc *shortenController) Shorten(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return nil
}

func (uc *shortenController) ShortCode(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return nil
}

func (uc *shortenController) ShortCodeStats(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	return nil
}
