package repository

import "github.com/amartha-shorty/pkg/shorty"

type shortenRepository struct {
}

func NewShortyRepository() shorty.Repository {
	return &shortenRepository{}
}
