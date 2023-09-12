-- View: public.mv_flight

-- DROP MATERIALIZED VIEW IF EXISTS public.mv_flight;

CREATE MATERIALIZED VIEW IF NOT EXISTS public.mv_all_flight
TABLESPACE pg_default
AS
 SELECT flight.hex_code,
    flight.flight_code,
    flight.flight_time,
    master.name,
    reference.mfr,
    reference.model
   FROM flight
     JOIN master ON flight.hex_code::text = master.mode_s_code_hex
     JOIN reference ON master.mfr_mdl_code = reference.code
WITH DATA;

ALTER TABLE IF EXISTS public.mv_flight
    OWNER TO pi;