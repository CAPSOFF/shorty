package controller

import (
	"context"
	"fmt"
	stdUrl "net/url"
	"regexp"
	"time"

	"github.com/amartha-shorty/pkg/model"
	"github.com/amartha-shorty/pkg/shorty"
	"github.com/lucasjones/reggen"
)

var inMemoryMap map[string]model.ShortySpec

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

func (uc *shortenController) Shorten(ctx context.Context, url string, desiredShortCode string) (shortCode string, errorCode int, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	shortCode = desiredShortCode
	if shortCode == "" {
		generatedString, err := reggen.Generate(model.RegexFormula, 6)
		if err != nil {
			return "", 500, err
		}
		shortCode = generatedString
	}

	_, err = stdUrl.ParseRequestURI(url)
	if err != nil {
		return "", 400, err
	}

	if _, ok := inMemoryMap[shortCode]; ok {
		return "", 409, fmt.Errorf("shortcode is already in use (%v)", shortCode)
	}

	constraint := regexp.MustCompile(model.RegexFormula)
	if !constraint.MatchString(shortCode) {
		return "", 422, fmt.Errorf("invalid shortcode format (%v)", shortCode)
	}

	inMemoryMap = map[string]model.ShortySpec{
		shortCode: model.ShortySpec{
			URL:           url,
			StartDate:     time.Now().String(),
			LastSeenDate:  "",
			RedirectCount: 0,
		},
	}

	return shortCode, 201, nil
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
