--
-- PostgreSQL database dump
--

-- Dumped from database version 12.15 (Ubuntu 12.15-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.15 (Ubuntu 12.15-0ubuntu0.20.04.1)

-- Started on 2023-07-18 14:04:30 WIB

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 203 (class 1259 OID 32795)
-- Name: mst_customer; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mst_customer (
    id integer NOT NULL,
    customer_name character varying(50) NOT NULL,
    phone_number character varying(20) NOT NULL,
    jumlah_transaksi integer
);


ALTER TABLE public.mst_customer OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 32793)
-- Name: mst_customer_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mst_customer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mst_customer_id_seq OWNER TO postgres;

--
-- TOC entry 2988 (class 0 OID 0)
-- Dependencies: 202
-- Name: mst_customer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mst_customer_id_seq OWNED BY public.mst_customer.id;


--
-- TOC entry 204 (class 1259 OID 32811)
-- Name: mst_service; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mst_service (
    id character varying(20) NOT NULL,
    service_name character varying(20) NOT NULL,
    price integer NOT NULL
);


ALTER TABLE public.mst_service OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 40999)
-- Name: trx_transaction; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trx_transaction (
    id integer NOT NULL,
    customer_id integer,
    service_id character varying(20),
    quantity integer,
    transaction_date date,
    total_transaksi integer
);


ALTER TABLE public.trx_transaction OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 40997)
-- Name: trx_transaction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.trx_transaction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.trx_transaction_id_seq OWNER TO postgres;

--
-- TOC entry 2989 (class 0 OID 0)
-- Dependencies: 205
-- Name: trx_transaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.trx_transaction_id_seq OWNED BY public.trx_transaction.id;


--
-- TOC entry 2842 (class 2604 OID 32798)
-- Name: mst_customer id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mst_customer ALTER COLUMN id SET DEFAULT nextval('public.mst_customer_id_seq'::regclass);


--
-- TOC entry 2843 (class 2604 OID 41002)
-- Name: trx_transaction id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trx_transaction ALTER COLUMN id SET DEFAULT nextval('public.trx_transaction_id_seq'::regclass);


--
-- TOC entry 2979 (class 0 OID 32795)
-- Dependencies: 203
-- Data for Name: mst_customer; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mst_customer (id, customer_name, phone_number, jumlah_transaksi) FROM stdin;
12	anna	076485936478	3
15	maria	089765678656	1
14	luna	0948576390	1
16	bear	089764536271	1
\.


--
-- TOC entry 2980 (class 0 OID 32811)
-- Dependencies: 204
-- Data for Name: mst_service; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mst_service (id, service_name, price) FROM stdin;
S001	Cuci + Setrika	7000
S002	Laundry Bedcover	50000
S003	Laundry Boneka	25000
\.


--
-- TOC entry 2982 (class 0 OID 40999)
-- Dependencies: 206
-- Data for Name: trx_transaction; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trx_transaction (id, customer_id, service_id, quantity, transaction_date, total_transaksi) FROM stdin;
20	12	S001	2	2022-01-01	14000
18	12	S001	2	2022-10-10	14000
19	12	S002	2	2022-10-10	100000
21	14	S003	3	2022-10-10	75000
22	15	S003	1	2022-12-12	25000
23	16	S003	4	2022-12-12	100000
\.


--
-- TOC entry 2990 (class 0 OID 0)
-- Dependencies: 202
-- Name: mst_customer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mst_customer_id_seq', 16, true);


--
-- TOC entry 2991 (class 0 OID 0)
-- Dependencies: 205
-- Name: trx_transaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.trx_transaction_id_seq', 23, true);


--
-- TOC entry 2845 (class 2606 OID 32800)
-- Name: mst_customer mst_customer_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mst_customer
    ADD CONSTRAINT mst_customer_pkey PRIMARY KEY (id);


--
-- TOC entry 2847 (class 2606 OID 32815)
-- Name: mst_service mst_service_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mst_service
    ADD CONSTRAINT mst_service_pkey PRIMARY KEY (id);


--
-- TOC entry 2849 (class 2606 OID 41004)
-- Name: trx_transaction trx_transaction_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trx_transaction
    ADD CONSTRAINT trx_transaction_pkey PRIMARY KEY (id);


--
-- TOC entry 2850 (class 2606 OID 41005)
-- Name: trx_transaction trx_transaction_customer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trx_transaction
    ADD CONSTRAINT trx_transaction_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES public.mst_customer(id);


--
-- TOC entry 2851 (class 2606 OID 41010)
-- Name: trx_transaction trx_transaction_service_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trx_transaction
    ADD CONSTRAINT trx_transaction_service_id_fkey FOREIGN KEY (service_id) REFERENCES public.mst_service(id);


-- Completed on 2023-07-18 14:04:31 WIB

--
-- PostgreSQL database dump complete
--

