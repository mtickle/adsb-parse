-- FUNCTION: public.get_flight(text)
-- DROP FUNCTION IF EXISTS public.get_flight(text);
CREATE
OR REPLACE FUNCTION public.get_flight(_flight_code text) RETURNS TABLE(
    hex_code character varying(6),
    type_code character varying(12),
    flight_code character varying(255),
    flight_time timestamp with time zone,
    alt_baro bigint, 
    category character varying(255),
    lat character varying(255),
    lon character varying(255),
    owner_name text,
    mfr text,
    model text
) LANGUAGE 'plpgsql' COST 100 VOLATILE PARALLEL UNSAFE ROWS 1000 AS $BODY$ begin return query
SELECT
    flight.hex_code,
    flight.type_code,
    flight.flight_code,
    flight.flight_time,
    flight.alt_baro,
    flight.category,
    flight.lat,
    flight.lon,
    master.name,
    reference.mfr,
    reference.model
FROM
    flight
    JOIN master ON flight.hex_code :: text = master.mode_s_code_hex
    JOIN reference ON master.mfr_mdl_code = reference.code
where
    flight.flight_code = _flight_code
ORDER BY
    flight_time desc;

end;

$BODY$;

ALTER FUNCTION public.get_flight(text) OWNER TO postgres;