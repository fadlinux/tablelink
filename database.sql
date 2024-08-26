

CREATE TABLE public.role_rights (
    id bigint NOT NULL,
    role_id bigint NOT NULL,
    route character varying,
    section character varying,
    path character varying,
    r_create integer DEFAULT 0,
    r_read integer DEFAULT 0,
    r_update integer DEFAULT 0,
    r_delete integer DEFAULT 0,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

CREATE TABLE public.roles (
    id bigint NOT NULL,
    name character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);

CREATE TABLE public.users (
    id bigint NOT NULL,
    role_id bigint NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    name character varying,
    last_access time without time zone,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);