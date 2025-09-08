ALTER TABLE interesting_aircraft ADD COLUMN commit_hash VARCHAR(40);
ALTER TABLE interesting_aircraft ADD CONSTRAINT interesting_aircraft_icao_unique UNIQUE (icao);