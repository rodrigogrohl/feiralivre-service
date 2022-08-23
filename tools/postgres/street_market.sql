CREATE TABLE public.street_market (
	id bigserial NOT NULL,
	longitude float8 NULL,
	latitude float8 NULL,
	sector_cense int8 NULL,
	area_ponderate int8 NULL,
	district_code int8 NULL,
	district varchar NULL,
	sub_town_code int8 NULL,
	sub_town varchar NULL,
	region5 varchar NULL,
	region8 varchar NULL,
	name_alias varchar NULL,
	registry varchar NULL,
	addr varchar NULL,
	addr_number varchar NULL,
	neighborhood varchar NULL,
	reference varchar NULL,
	CONSTRAINT street_market_pkey PRIMARY KEY (id)
);
CREATE INDEX street_market_district_idx ON public.street_market USING btree (district);
CREATE INDEX street_market_name_alias_idx ON public.street_market USING btree (name_alias);
CREATE INDEX street_market_neighborhood_idx ON public.street_market USING btree (neighborhood);
CREATE INDEX street_market_region5_idx ON public.street_market USING btree (region5);
