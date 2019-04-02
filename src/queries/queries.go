package queries

var (
	CreatePlaceBookmarksTableQuery = `CREATE TABLE IF NOT EXISTS place_bookmarks (
		id SERIAL PRIMARY KEY,
		name varchar(100) NOT NULL)`
	SelectPlaceBookmarksQuery = `SELECT name FROM place_bookmarks`
	InsertPlaceBookmarkQuery = `INSERT INTO place_bookmarks (name) VALUES ($1)`
	CreatePlaceQueriesTableQuery = `CREATE TABLE IF NOT EXISTS place_queries (
		id SERIAL PRIMARY KEY,
		name varchar(100) NOT NULL,
		query_date DATE NOT NULL,
		type VARCHAR(100) NOT NULL,
		temp NUMERIC,
		min_temp NUMERIC,
		max_temp NUMERIC)`
	SelectPlaceQueriesQuery = `SELECT to_char(query_date, 'MM-YYYY') as date, count(*) as queries_count, AVG(temp) as average_temp, MIN(min_temp) as min_temp, MAX(max_temp) as max_temp 
FROM place_queries WHERE name=$1 GROUP BY date`
	SelectPlaceQueriesTypesOccurances = `SELECT count(*) as queries_count, type FROM place_queries WHERE name=$1 and to_char(query_date, 'MM-YYYY')=$2 GROUP BY type;`
	InsertPlaceQueryQuery = `INSERT INTO place_queries (name, query_date, type, temp, min_temp, max_temp) values ($1, $2, $3, $4, $5, $6)`
)