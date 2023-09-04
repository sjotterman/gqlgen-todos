--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE IF EXISTS postgres;
--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: user
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE postgres OWNER TO "user";

\connect postgres

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: user
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: restaurant; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.restaurant (
    id integer NOT NULL,
    name character varying(140) NOT NULL,
    description character varying(1024) NOT NULL,
    phone_number character varying(40) NOT NULL
);


ALTER TABLE public.restaurant OWNER TO "user";

--
-- Name: restaurant_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.restaurant_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.restaurant_id_seq OWNER TO "user";

--
-- Name: restaurant_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.restaurant_id_seq OWNED BY public.restaurant.id;


--
-- Name: restaurant id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.restaurant ALTER COLUMN id SET DEFAULT nextval('public.restaurant_id_seq'::regclass);


--
-- Data for Name: restaurant; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.restaurant (id, name, description, phone_number) FROM stdin;
1	Applebees's	American Casual	330-666-2261
2	Romeo's	Pizza	(330) 836-6000
\.


--
-- Name: restaurant_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.restaurant_id_seq', 1, true);


--
-- Name: restaurant restaurant_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.restaurant
    ADD CONSTRAINT restaurant_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

