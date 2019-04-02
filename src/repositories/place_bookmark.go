package repositories

import (
	"database/sql"
	"errors"
	"github.com/kiperz/weather-api/queries"
)

type PlaceBookmark struct {
	Name string `json:"name"`
}

type PlaceBookmarkRepository struct {
	conn *sql.DB
}

func NewPlaceBookmarkRepository(db *sql.DB) (*PlaceBookmarkRepository, error) {
	PlaceBookmarkRepository := &PlaceBookmarkRepository{
		conn: db,
	}

	if err := PlaceBookmarkRepository.ensureTableIsCreated(); err != nil {
		return nil, err
	}

	return PlaceBookmarkRepository, nil
}

func (pbr *PlaceBookmarkRepository) ensureTableIsCreated() error {
	if _, err := pbr.conn.Exec(queries.CreatePlaceBookmarksTableQuery); err != nil {
		return err
	}
	return nil
}

func (pbr *PlaceBookmarkRepository) Get() ([]PlaceBookmark, error) {
	placeBookmarks := []PlaceBookmark{}

	rows, err := pbr.conn.Query(queries.SelectPlaceBookmarksQuery);
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			return nil, err
		}
		placeBookmarks = append(placeBookmarks, PlaceBookmark{
			Name: name,
		})
	}
	return placeBookmarks, nil
}

func (pbr *PlaceBookmarkRepository) Put(placeBookmark *PlaceBookmark) error {
	if placeBookmark == nil {
		return errors.New("placeBookmark is nil")
	}

	if _, err := pbr.conn.Exec(queries.InsertPlaceBookmarkQuery, placeBookmark.Name); err != nil {
		return err
	}
	return nil
}
