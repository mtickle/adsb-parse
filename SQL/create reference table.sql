-- Table: public.adsb_reference



-- DROP TABLE IF EXISTS public.adsb_reference;

CREATE TABLE IF NOT EXISTS public.adsb_reference
(
    hex_code character varying COLLATE pg_catalog."default",
	reg_code character varying COLLATE pg_catalog."default",
    manufacturer_icao character varying COLLATE pg_catalog."default",
    manufacturer_name character varying COLLATE pg_catalog."default",
    type_code character varying COLLATE pg_catalog."default",
    model_code character varying COLLATE pg_catalog."default",
    serial_number character varying COLLATE pg_catalog."default",
    tbd character varying COLLATE pg_catalog."default",
    icao_aircraft_type character varying COLLATE pg_catalog."default",
    operator_name character varying COLLATE pg_catalog."default",
    operator_icao character varying COLLATE pg_catalog."default",
    operator_callsign character varying COLLATE pg_catalog."default",
    tbd2 character varying COLLATE pg_catalog."default",
    owner_name character varying COLLATE pg_catalog."default",
    tbd3 character varying COLLATE pg_catalog."default"
    
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.adsb_reference
    OWNER to pi;



    COPY adsb_reference(hex_code, reg_code, manufacturer_icao, manufacturer_name, type_code, model_code, serial_number, tbd, icao_aircraft_type, operator_name, operator_icao, operator_callsign, tbd2, owner_name, tbd3)
FROM '/tmp/db2.csv'
DELIMITER ','
CSV HEADER;