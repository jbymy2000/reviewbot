CREATE TABLE IF NOT EXISTS public.ratings
(
    id integer NOT NULL DEFAULT nextval('ratings_id_seq'::regclass),
    user_id bigint,
    rating integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    psid bigint DEFAULT 0,
    CONSTRAINT ratings_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.ratings
    OWNER to postgres;