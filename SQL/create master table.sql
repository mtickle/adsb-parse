-- Table: public.master

-- DROP TABLE IF EXISTS public.master;

CREATE TABLE IF NOT EXISTS public.master
(
    "N_NUMBER" text COLLATE pg_catalog."default",
    "SERIAL_NUMBER" text COLLATE pg_catalog."default",
    "MFR_MDL_CODE" text COLLATE pg_catalog."default",
    "ENG_MFR_MDL" text COLLATE pg_catalog."default",
    "YEAR_MFR" text COLLATE pg_catalog."default",
    "NAME" text COLLATE pg_catalog."default",
    "UNIQUE_ID" text COLLATE pg_catalog."default",
    "MODE_S_CODE_HEX" text COLLATE pg_catalog."default"
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.master
    OWNER to pi;