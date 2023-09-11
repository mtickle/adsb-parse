-- PROCEDURE: public.add_flight(character varying, character varying, character varying, timestamp with time zone, bigint, character varying, character varying, character varying)

-- DROP PROCEDURE IF EXISTS public.add_flight(character varying, character varying, character varying, timestamp with time zone, bigint, character varying, character varying, character varying);

CREATE OR REPLACE PROCEDURE public.add_flight(
	_hex_code character varying,
	_type_code character varying,
	_flight_code character varying,
	_flight_time timestamp with time zone,
	_alt_baro bigint,
	_category character varying,
	_lat character varying,
	_lon character varying)
LANGUAGE 'sql'
AS $BODY$
INSERT INTO
  public.flight (
    hex_code,
    type_code,
    flight_code,
    flight_time,
	alt_baro,
    category,
    lat,
    lon
  )
VALUES
  (
    _hex_code,
    _type_code,
    _flight_code,
    _flight_time,
    _alt_baro,
    _category,
    _lat,
    _lon
    
  ) 
  ON CONFLICT (hex_code, lat, lon) DO NOTHING;
$BODY$;

ALTER PROCEDURE public.add_flight(character varying, character varying, character varying, timestamp with time zone, bigint, character varying, character varying, character varying)
    OWNER TO pi;
