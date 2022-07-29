package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/gymnasy55/adviser-bot/lib/e"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

var ErrNoSavedPage = errors.New("no saved page")

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("Can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("Can't calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
