-- View: public.mv_all_flights

-- DROP MATERIALIZED VIEW IF EXISTS public.mv_all_flights;

CREATE MATERIALIZED VIEW IF NOT EXISTS public.mv_all_flights
TABLESPACE pg_default
AS
  SELECT flight.hex_code,
    flight.flight_code,
    flight.flight_time,
    flight.type_code,
    flight.category,
    flight.squawk,
    adsb_reference.operator_icao,
	adsb_reference.owner_name,
	master.name,
	adsb_reference.manufacturer_name,
    reference.mfr,
	adsb_reference.model_code,
    reference.model
	
   FROM flight
   		join adsb_reference on lower(flight.hex_code::text) = adsb_reference.hex_code
     JOIN master ON flight.hex_code::text = master.mode_s_code_hex
     JOIN reference ON master.mfr_mdl_code = reference.code
	 ORDER by flight_time desc
WITH DATA;

ALTER TABLE IF EXISTS public.mv_all_flights
    OWNER TO pi;