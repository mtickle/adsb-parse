-- Table: public.flight

-- DROP TABLE IF EXISTS public.flight;

CREATE TABLE IF NOT EXISTS public.flight
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    hex_code character varying(6) COLLATE pg_catalog."default",
    type_code character varying(12) COLLATE pg_catalog."default",
    flight_code character varying(255) COLLATE pg_catalog."default",
    flight_time timestamp with time zone,
    alt_baro bigint,
    category character varying(255) COLLATE pg_catalog."default",
    lat character varying(255) COLLATE pg_catalog."default",
    lon character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT flight_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.flight
    OWNER to pi;