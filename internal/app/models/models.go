package models

import (
	"github.com/Agniy/shortener/internal/app/storage"
	"time"
)

type Links struct {
	ID        uint64 `gorm:"primaryKey"`
	URL       string
	ShortURL  string
	UrlKey    string    `gorm:"index"`
	CreatedAt time.Time `gorm:"index"`
}

func (l *Links) TableName() string {
	return "links"
}

func GetLink(urlKey string) (Links, error) {
	psql, err := storage.GetDbClient()
	if err != nil {
		return Links{}, err
	}
	var link Links
	result := psql.First(&link, "url_key = ?", urlKey)
	if result.Error != nil {
		return Links{}, result.Error
	}
	return link, nil
}

// create a new link
func CreateLink(url, shortUrl, urlKey string) error {
	psql, err := storage.GetDbClient()
	if err != nil {
		return err
	}

	result := psql.FirstOrCreate(&Links{UrlKey: urlKey}, Links{
		URL:      url,
		ShortURL: shortUrl,
		UrlKey:   urlKey,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
