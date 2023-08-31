-- Table: public.flight

-- DROP TABLE IF EXISTS public.flight;

CREATE TABLE IF NOT EXISTS public.flight
(
     id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    hex_code character varying(6) COLLATE pg_catalog."default",
    type_code character varying(12) COLLATE pg_catalog."default",
    flight character varying(255) COLLATE pg_catalog."default",
    alt_baro bigint,
    alt_geom bigint,
    gs character varying(255) COLLATE pg_catalog."default",
    track character varying(255) COLLATE pg_catalog."default",
    baro_rate character varying(255) COLLATE pg_catalog."default",
    category character varying(255) COLLATE pg_catalog."default",
    nav_qnh character varying(255) COLLATE pg_catalog."default",
    nav_altitude_mcp character varying(255) COLLATE pg_catalog."default",
    lat character varying(255) COLLATE pg_catalog."default",
    lon character varying(255) COLLATE pg_catalog."default",
    nic character varying(255) COLLATE pg_catalog."default",
    rc character varying(255) COLLATE pg_catalog."default",
    seen_pos character varying(255) COLLATE pg_catalog."default",
    r_dst character varying(255) COLLATE pg_catalog."default",
    r_dir character varying(255) COLLATE pg_catalog."default",
    version character varying(255) COLLATE pg_catalog."default",
    nic_baro character varying(255) COLLATE pg_catalog."default",
    nac_p character varying(255) COLLATE pg_catalog."default",
    nac_v character varying(255) COLLATE pg_catalog."default",
    sil character varying(255) COLLATE pg_catalog."default",
    sil_type character varying(255) COLLATE pg_catalog."default",
    gva character varying(255) COLLATE pg_catalog."default",
    sda character varying(255) COLLATE pg_catalog."default",
    alert character varying(255) COLLATE pg_catalog."default",
    spi character varying(255) COLLATE pg_catalog."default",
    mlat character varying(255) COLLATE pg_catalog."default",
    tisb character varying(255) COLLATE pg_catalog."default",
    messages character varying(255) COLLATE pg_catalog."default",
    seen character varying(255) COLLATE pg_catalog."default",
    rssi character varying(255) COLLATE pg_catalog."default",
    flight_time timestamp with time zone,
    squawk character varying(255) COLLATE pg_catalog."default",
    emergency character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT flight_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.flight
    OWNER to pi;