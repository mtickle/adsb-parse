-- Table: public.reference

-- DROP TABLE IF EXISTS public.reference;

CREATE TABLE IF NOT EXISTS public.reference
(
    "CODE" text COLLATE pg_catalog."default",
    "MFR" text COLLATE pg_catalog."default",
    "MODEL" text COLLATE pg_catalog."default"
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.reference
    OWNER to pi;