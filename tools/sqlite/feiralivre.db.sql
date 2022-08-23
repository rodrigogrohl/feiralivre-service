BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "street_market" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"longitude"	INTEGER,
	"latitude"	INTEGER,
	"sector_cense"	INTEGER,
	"area_ponderate"	INTEGER,
	"district_code"	INTEGER,
	"district"	TEXT,
	"sub_town_code"	INTEGER,
	"sub_town"	TEXT,
	"region5"	TEXT,
	"region8"	TEXT,
	"name_alias"	TEXT,
	"registry"	TEXT,
	"addr"	TEXT,
	"addr_number"	TEXT,
	"neighborhood"	TEXT,
	"reference"	TEXT
);
CREATE INDEX IF NOT EXISTS "neighborhood_idx" ON "street_market" (
	"neighborhood"	ASC
);
CREATE INDEX IF NOT EXISTS "name_idx" ON "street_market" (
	"name_alias"	ASC
);
CREATE INDEX IF NOT EXISTS "region5_idx" ON "street_market" (
	"region5"	ASC
);
CREATE INDEX IF NOT EXISTS "district_idx" ON "street_market" (
	"district"	ASC
);
COMMIT;
