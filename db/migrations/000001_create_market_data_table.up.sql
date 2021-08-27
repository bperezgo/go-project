CREATE TABLE IF NOT EXISTS market_data(
	id UUID PRIMARY KEY,
	time TIMESTAMP NOT NULL,
	open_value REAL,
	high_value REAL,
	low_value REAL,
	close_value REAL,
	symbol VARCHAR (5) NOT NULL
);
