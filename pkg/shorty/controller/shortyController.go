package controller

import (
	"context"
	"fmt"
	"net/http"
	stdUrl "net/url"
	"regexp"
	"time"

	"github.com/amartha-shorty/pkg/model"
	"github.com/amartha-shorty/pkg/shorty"
	"github.com/lucasjones/reggen"
)

var inMemoryMap map[string]model.ShortySpec

type shortyController struct {
	timeout          time.Duration
	shortyRepository shorty.Repository
}

func NewShortyController(timeout time.Duration, shortyRepository shorty.Repository) shorty.Controller {
	return &shortyController{
		timeout:          timeout,
		shortyRepository: shortyRepository,
	}
}

func (uc *shortyController) Shorten(ctx context.Context, url string, desiredShortCode string) (shortCode string, httpStatusCode int, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	shortCode = desiredShortCode
	if shortCode == "" {
		generatedString, err := reggen.Generate(model.RegexFormula, 6)
		if err != nil {
			return "", http.StatusInternalServerError, err
		}
		shortCode = generatedString
	}

	_, err = stdUrl.ParseRequestURI(url)
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	if _, ok := inMemoryMap[shortCode]; ok {
		return "", http.StatusConflict, fmt.Errorf("shortcode is already in use (%v)", shortCode)
	}

	constraint := regexp.MustCompile(model.RegexFormula)
	if !constraint.MatchString(shortCode) {
		return "", http.StatusUnprocessableEntity, fmt.Errorf("invalid shortcode format (%v)", shortCode)
	}

	inMemoryMap = map[string]model.ShortySpec{
		shortCode: model.ShortySpec{
			URL:           url,
			StartDate:     time.Now().String(),
			LastSeenDate:  "",
			RedirectCount: 0,
		},
	}

	return shortCode, http.StatusCreated, nil
}

func (uc *shortyController) ShortCode(ctx context.Context, shortCode string) (url string, httpStatusCode int, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	if _, ok := inMemoryMap[shortCode]; !ok {
		return "", http.StatusNotFound, fmt.Errorf("shortcode cannot be found in the system (%v)", shortCode)
	}

	inMemoryMap[shortCode] = model.ShortySpec{
		URL:           inMemoryMap[shortCode].URL,
		StartDate:     inMemoryMap[shortCode].StartDate,
		LastSeenDate:  time.Now().String(),
		RedirectCount: inMemoryMap[shortCode].RedirectCount + 1,
	}

	return inMemoryMap[shortCode].URL, http.StatusFound, nil
}

func (uc *shortyController) ShortCodeStats(ctx context.Context, shortCode string) (shortyData model.ShortySpec, httpStatusCode int, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	if _, ok := inMemoryMap[shortCode]; !ok {
		return model.ShortySpec{}, http.StatusNotFound, fmt.Errorf("shortcode cannot be found in the system (%v)", shortCode)
	}

	return inMemoryMap[shortCode], http.StatusOK, nil
}
