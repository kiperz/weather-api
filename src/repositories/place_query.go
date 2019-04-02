package repositories

import (
	"database/sql"
	"errors"
	"github.com/kiperz/weather-api/queries"
)

type PlaceQuery struct {
	Name string
	Date string
	Type string
	Temperature float64
	MinimumTemperature float64
	MaximumTemperature float64
}

type Stats struct {
	Count int `json:"count"`
	Stats map[string]*MonthStats `json:"stats"`
}

type MonthStats struct {
	 Avarage float64 `json:"avg"`
	 Lowest float64 `json:"low"`
	 Highest float64 `json:"high"`
	 Types map[string]int `json:"types"`
}

type PlaceQueryRepository struct {
	conn *sql.DB
}

func NewPlaceQueryRepository(db *sql.DB) (*PlaceQueryRepository, error) {
	PlaceQueryRepository := &PlaceQueryRepository{
		conn: db,
	}

	if err := PlaceQueryRepository.ensureTableIsCreated(); err != nil {
		return nil, err
	}

	return PlaceQueryRepository, nil
}

func (pqr *PlaceQueryRepository) ensureTableIsCreated() (error) {
	if _, err := pqr.conn.Exec(queries.CreatePlaceQueriesTableQuery); err != nil {
		return err
	}
	return nil
}

func (pqr *PlaceQueryRepository) Put(placeQuery *PlaceQuery) (error) {
	if placeQuery == nil {
		return errors.New("placeQuery is nil")
	}

	if _, err := pqr.conn.Exec(queries.InsertPlaceQueryQuery, placeQuery.Name, placeQuery.Date, placeQuery.Type,
		placeQuery.Temperature, placeQuery.MinimumTemperature, placeQuery.MaximumTemperature); err != nil {
		return err
	}
	return nil
}

func (pqr *PlaceQueryRepository) GetStatistics(name string) (*Stats, error) {
	rows, err := pqr.conn.Query(queries.SelectPlaceQueriesQuery, name);
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	statsPerMonth := make(map[string]*MonthStats)
	overallQueries := 0
	for rows.Next() {
		var (
			stat MonthStats
			date string
			queriesCount int
		)
		if err = rows.Scan(&date, &queriesCount, &stat.Avarage, &stat.Lowest, &stat.Highest); err != nil {
			return nil, err
		}
		if stat.Types, err = pqr.getTypes(name, date); err != nil {
			return nil, err
		}
		overallQueries += queriesCount
		statsPerMonth[date] = &stat
	}
	return &Stats{
		Count: overallQueries,
		Stats: statsPerMonth,
	}, nil
}

func (pqr *PlaceQueryRepository) getTypes(name string, month string) (map[string]int, error) {
	rows, err := pqr.conn.Query(queries.SelectPlaceQueriesTypesOccurances, name, month);
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	res := make(map[string]int)
	for rows.Next() {
		var (
			count int
			queryType string
		)
		if err := rows.Scan(&count, &queryType); err != nil {
			return nil, err
		}
		res[queryType] = count

	}
	return res, nil
}