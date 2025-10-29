--
-- PostgreSQL database dump
--

\restrict h1Jhr5F3wvcmuM7HDNp1es9bzO6AbKplCX7G8rHCKAeqRdcyXChLZDq3MZIesnr

-- Dumped from database version 18.0 (Debian 18.0-1.pgdg13+3)
-- Dumped by pg_dump version 18.0 (Debian 18.0-1.pgdg13+3)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: accounts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.accounts (
    id bigint NOT NULL,
    user_name text NOT NULL,
    password text NOT NULL,
    role_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.accounts OWNER TO postgres;

--
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.accounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.accounts_id_seq OWNER TO postgres;

--
-- Name: accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.accounts_id_seq OWNED BY public.accounts.id;


--
-- Name: attendances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.attendances (
    id bigint NOT NULL,
    shift_schedule_id bigint,
    actual_start_time timestamp with time zone,
    actual_end_time timestamp with time zone,
    hours bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.attendances OWNER TO postgres;

--
-- Name: attendances_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.attendances_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.attendances_id_seq OWNER TO postgres;

--
-- Name: attendances_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.attendances_id_seq OWNED BY public.attendances.id;


--
-- Name: availibilities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.availibilities (
    id bigint NOT NULL,
    employee_id bigint,
    shift_id bigint NOT NULL,
    day_of_week character varying(10) NOT NULL,
    is_available boolean DEFAULT false,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.availibilities OWNER TO postgres;

--
-- Name: availibilities_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.availibilities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.availibilities_id_seq OWNER TO postgres;

--
-- Name: availibilities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.availibilities_id_seq OWNED BY public.availibilities.id;


--
-- Name: bookings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bookings (
    id bigint NOT NULL,
    customer_name text NOT NULL,
    phone_number text NOT NULL,
    email text NOT NULL,
    booking_date text NOT NULL,
    booking_time text NOT NULL,
    total_persons bigint NOT NULL,
    status text DEFAULT 'pending'::text NOT NULL,
    memo text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.bookings OWNER TO postgres;

--
-- Name: bookings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bookings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.bookings_id_seq OWNER TO postgres;

--
-- Name: bookings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bookings_id_seq OWNED BY public.bookings.id;


--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id bigint NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: customers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customers (
    id bigint NOT NULL,
    full_name text,
    phone_number text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.customers OWNER TO postgres;

--
-- Name: customers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.customers_id_seq OWNER TO postgres;

--
-- Name: customers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.customers_id_seq OWNED BY public.customers.id;


--
-- Name: employees; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.employees (
    id bigint NOT NULL,
    full_name text,
    gender boolean,
    birthday text,
    phone_number text,
    email text,
    schedule_type text,
    address text,
    join_date text,
    base_salary bigint,
    salary_per_hour bigint,
    account_id bigint,
    avatar_file_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.employees OWNER TO postgres;

--
-- Name: employees_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.employees_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.employees_id_seq OWNER TO postgres;

--
-- Name: employees_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.employees_id_seq OWNED BY public.employees.id;


--
-- Name: files; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.files (
    id bigint NOT NULL,
    file_name text NOT NULL,
    url text NOT NULL,
    mime_type text,
    size bigint,
    public_id text NOT NULL,
    resource_type text,
    folder text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.files OWNER TO postgres;

--
-- Name: files_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.files_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.files_id_seq OWNER TO postgres;

--
-- Name: files_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.files_id_seq OWNED BY public.files.id;


--
-- Name: ingredients; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ingredients (
    id bigint NOT NULL,
    name text NOT NULL,
    description text,
    quantity bigint,
    warning_quantity bigint NOT NULL,
    supplier text,
    unit text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.ingredients OWNER TO postgres;

--
-- Name: ingredients_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ingredients_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ingredients_id_seq OWNER TO postgres;

--
-- Name: ingredients_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ingredients_id_seq OWNED BY public.ingredients.id;


--
-- Name: menu_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.menu_items (
    id bigint NOT NULL,
    name text NOT NULL,
    description text,
    price bigint NOT NULL,
    category_id bigint,
    file_id bigint,
    status text DEFAULT 'Available'::text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.menu_items OWNER TO postgres;

--
-- Name: menu_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.menu_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.menu_items_id_seq OWNER TO postgres;

--
-- Name: menu_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.menu_items_id_seq OWNED BY public.menu_items.id;


--
-- Name: order_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_items (
    id bigint NOT NULL,
    order_id bigint NOT NULL,
    menu_item_id bigint NOT NULL,
    quantity bigint NOT NULL,
    amount bigint,
    memo text,
    status text DEFAULT 'Pending'::text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.order_items OWNER TO postgres;

--
-- Name: order_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_items_id_seq OWNER TO postgres;

--
-- Name: order_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_items_id_seq OWNED BY public.order_items.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id bigint NOT NULL,
    customer_id bigint,
    table_id bigint NOT NULL,
    amount bigint NOT NULL,
    status text DEFAULT 'UnPaid'::text NOT NULL,
    memo text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.orders_id_seq OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- Name: permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.permissions (
    id bigint NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.permissions OWNER TO postgres;

--
-- Name: permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.permissions_id_seq OWNER TO postgres;

--
-- Name: permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.permissions_id_seq OWNED BY public.permissions.id;


--
-- Name: role_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role_permissions (
    permission_id bigint NOT NULL,
    role_id bigint NOT NULL
);


ALTER TABLE public.role_permissions OWNER TO postgres;

--
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id bigint NOT NULL,
    role_name text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.roles_id_seq OWNER TO postgres;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: shift_schedules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shift_schedules (
    id bigint NOT NULL,
    employee_id bigint NOT NULL,
    shift_id bigint NOT NULL,
    date text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.shift_schedules OWNER TO postgres;

--
-- Name: shift_schedules_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.shift_schedules_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.shift_schedules_id_seq OWNER TO postgres;

--
-- Name: shift_schedules_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.shift_schedules_id_seq OWNED BY public.shift_schedules.id;


--
-- Name: shifts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shifts (
    id bigint NOT NULL,
    shift_name text NOT NULL,
    code text NOT NULL,
    start_time text NOT NULL,
    end_time text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.shifts OWNER TO postgres;

--
-- Name: shifts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.shifts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.shifts_id_seq OWNER TO postgres;

--
-- Name: shifts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.shifts_id_seq OWNED BY public.shifts.id;


--
-- Name: tables; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tables (
    id bigint NOT NULL,
    table_name text NOT NULL,
    "position" text,
    seats bigint,
    memo text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.tables OWNER TO postgres;

--
-- Name: tables_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tables_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tables_id_seq OWNER TO postgres;

--
-- Name: tables_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tables_id_seq OWNED BY public.tables.id;


--
-- Name: tickets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tickets (
    id bigint NOT NULL,
    ingredient_id bigint NOT NULL,
    quantity bigint,
    unit text NOT NULL,
    ticket_type text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.tickets OWNER TO postgres;

--
-- Name: tickets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tickets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tickets_id_seq OWNER TO postgres;

--
-- Name: tickets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tickets_id_seq OWNED BY public.tickets.id;


--
-- Name: accounts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts ALTER COLUMN id SET DEFAULT nextval('public.accounts_id_seq'::regclass);


--
-- Name: attendances id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.attendances ALTER COLUMN id SET DEFAULT nextval('public.attendances_id_seq'::regclass);


--
-- Name: availibilities id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.availibilities ALTER COLUMN id SET DEFAULT nextval('public.availibilities_id_seq'::regclass);


--
-- Name: bookings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings ALTER COLUMN id SET DEFAULT nextval('public.bookings_id_seq'::regclass);


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: customers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers ALTER COLUMN id SET DEFAULT nextval('public.customers_id_seq'::regclass);


--
-- Name: employees id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees ALTER COLUMN id SET DEFAULT nextval('public.employees_id_seq'::regclass);


--
-- Name: files id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files ALTER COLUMN id SET DEFAULT nextval('public.files_id_seq'::regclass);


--
-- Name: ingredients id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ingredients ALTER COLUMN id SET DEFAULT nextval('public.ingredients_id_seq'::regclass);


--
-- Name: menu_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.menu_items ALTER COLUMN id SET DEFAULT nextval('public.menu_items_id_seq'::regclass);


--
-- Name: order_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items ALTER COLUMN id SET DEFAULT nextval('public.order_items_id_seq'::regclass);


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: permissions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.permissions ALTER COLUMN id SET DEFAULT nextval('public.permissions_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: shift_schedules id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shift_schedules ALTER COLUMN id SET DEFAULT nextval('public.shift_schedules_id_seq'::regclass);


--
-- Name: shifts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shifts ALTER COLUMN id SET DEFAULT nextval('public.shifts_id_seq'::regclass);


--
-- Name: tables id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tables ALTER COLUMN id SET DEFAULT nextval('public.tables_id_seq'::regclass);


--
-- Name: tickets id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tickets ALTER COLUMN id SET DEFAULT nextval('public.tickets_id_seq'::regclass);


--
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.accounts (id, user_name, password, role_id, created_at, updated_at, deleted_at) FROM stdin;
1	admin	$2a$10$skVoACozhx1E5Gp67XK4W.m76tLlPOVF9qRbzUjx40zTMUVo7J2uq	1	2025-10-16 01:27:44.282523+00	2025-10-16 01:27:44.282523+00	\N
10	25101802	$2a$10$WP3xGdzOojSUrnz.1Yvhfu5eAoSm9ayUoFYaG6jRDydu93fPohUN2	5	2025-10-18 02:46:36.741487+00	2025-10-18 02:46:59.932903+00	\N
11	25101803	$2a$10$lls.mAjL5qw0IUFajIFNqezjhU0y/a6vVz.sjwX2cu.v5wQSOlsFS	4	2025-10-18 02:50:21.883107+00	2025-10-18 02:50:21.883107+00	\N
12	25101804	$2a$10$lSHc9COCqhxhPavmjEfmQOEu/r7qgFG9YLsvKLEey0RsrV2Lh.At2	3	2025-10-18 02:52:38.726117+00	2025-10-18 02:52:38.726117+00	\N
13	25101805	$2a$10$KkgkzySPBQDNwm3.zKCEzOhkO8h3thHYswJ4/pZtaNQFVHMaTmz6.	6	2025-10-18 02:54:50.008661+00	2025-10-18 02:54:50.008661+00	\N
7	25101606	$2a$10$96/df21mJdQCk7gI1G7UNeeDzB5TMp/zqJ4jj/xwTeiAkdU2XHJO.	2	2025-10-16 02:24:55.375797+00	2025-10-17 03:07:47.796069+00	2025-10-18 02:58:00.280001+00
6	25101605	$2a$10$SKaHG3fISzuqxNCrmdtq2eHJ1ZMtiKuLmgr/bs3XaKIBx1cOcaiTi	7	2025-10-16 02:24:04.738466+00	2025-10-16 02:24:04.738466+00	2025-10-18 02:58:13.602747+00
5	25101604	$2a$10$tNUj5bIa6dkOpLIZ7aD31.ZEjvPDk0eXNKAC1fetaEpn/kSsPaSo2	6	2025-10-16 02:22:36.166687+00	2025-10-16 02:23:11.788101+00	2025-10-18 02:58:16.131337+00
4	25101603	$2a$10$akQboB3FG7WX6d5mngCu8OX7fdBQrmxmrP6S/3x2TLP3TdLpp6nOK	3	2025-10-16 02:20:59.51809+00	2025-10-16 02:20:59.51809+00	2025-10-18 02:58:20.261987+00
3	25101602	$2a$10$9R6tz1MsVuhwg2zaEHP/1OYYM0Aw6XEdnlJHVkioGtooJ3WtgddCa	5	2025-10-16 02:19:47.395635+00	2025-10-16 02:21:42.858914+00	2025-10-18 02:58:23.408454+00
2	25101601	$2a$10$Iv3Oa8dZi6bXGfzXTeTy1.xvRzMpCwF6iO7tFJ3aIBQ4nKv1Bi/iW	4	2025-10-16 01:33:30.871117+00	2025-10-16 01:33:30.871117+00	2025-10-18 02:58:27.204845+00
14	25101806	$2a$10$S4A.JXceqYfZBTdiCX7w9uEaNlOJ3.PHtRM3PmODPLKiYdl6UbHGq	6	2025-10-18 03:41:41.031437+00	2025-10-18 03:41:58.698841+00	\N
15	25101807	$2a$10$TWCEnPBtQnK0xWgHdUt5tep1g4..mq93oYUPHSbOxWEGtGO9fi0YW	1	2025-10-18 03:45:06.712633+00	2025-10-18 03:45:06.712633+00	\N
9	25101801	$2a$10$/fqqfzzoYxB.FJWxPoF9Weiea5QQTzi4Ln2URp9dEW1RaJQBKqG7G	2	2025-10-18 02:44:36.206162+00	2025-10-18 03:46:01.318082+00	\N
16	25101808	$2a$10$Ye5qFmTs2Tg7NpSPzZkkL.H4cF3NLkM4WQr4OLCPmc69burOkZKUq	4	2025-10-18 03:47:10.654397+00	2025-10-18 03:47:10.654397+00	\N
17	25101809	$2a$10$vtQSZRKM7r6z2gGD7bmI7.Oz2L67Ti6q7nCC0k0Dh6mS10ysBiN1W	4	2025-10-18 03:48:52.143454+00	2025-10-18 03:48:52.143454+00	\N
\.


--
-- Data for Name: attendances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.attendances (id, shift_schedule_id, actual_start_time, actual_end_time, hours, created_at, updated_at, deleted_at) FROM stdin;
1	157	2025-10-20 03:05:07.784+00	2025-10-20 03:05:07.784+00	0	2025-10-20 03:05:07.826667+00	2025-10-20 03:05:07.826667+00	2025-10-20 04:20:46.5691+00
2	157	2025-10-20 03:05:26.929+00	2025-10-20 04:44:30.65+00	2	2025-10-20 03:05:26.918918+00	2025-10-20 04:44:30.675119+00	\N
\.


--
-- Data for Name: availibilities; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.availibilities (id, employee_id, shift_id, day_of_week, is_available, created_at, updated_at, deleted_at) FROM stdin;
1	1	1	Monday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
2	1	1	Tuesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
3	1	1	Wednesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
4	1	1	Thursday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
5	1	1	Friday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
6	1	1	Saturday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
7	1	1	Sunday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
8	1	2	Monday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
9	1	2	Tuesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
10	1	2	Wednesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
11	1	2	Thursday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
12	1	2	Friday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
13	1	2	Saturday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
14	1	2	Sunday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
15	1	3	Monday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
16	1	3	Tuesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
17	1	3	Wednesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
18	1	3	Thursday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
19	1	3	Friday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
20	1	3	Saturday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
21	1	3	Sunday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
64	4	1	Monday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
65	4	1	Tuesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
66	4	1	Wednesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
67	4	1	Thursday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
68	4	1	Friday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
69	4	1	Saturday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
70	4	1	Sunday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
71	4	2	Monday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
72	4	2	Tuesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
73	4	2	Wednesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
74	4	2	Thursday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
75	4	2	Friday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
76	4	2	Saturday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
77	4	2	Sunday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
78	4	3	Monday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
79	4	3	Tuesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
80	4	3	Wednesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
81	4	3	Thursday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
82	4	3	Friday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
83	4	3	Saturday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
84	4	3	Sunday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
85	5	1	Monday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
86	5	1	Tuesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
87	5	1	Wednesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
88	5	1	Thursday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
89	5	1	Friday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
90	5	1	Saturday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
91	5	1	Sunday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
92	5	2	Monday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
93	5	2	Tuesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
94	5	2	Wednesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
95	5	2	Thursday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
96	5	2	Friday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
97	5	2	Saturday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
98	5	2	Sunday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
99	5	3	Monday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
100	5	3	Tuesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
101	5	3	Wednesday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
30	2	2	Tuesday	t	2025-10-16 02:29:29.715813+00	2025-10-16 02:39:52.655011+00	\N
38	2	3	Wednesday	t	2025-10-16 02:29:29.715813+00	2025-10-16 02:39:56.738001+00	\N
25	2	1	Thursday	t	2025-10-16 02:29:29.715813+00	2025-10-16 02:39:57.999435+00	\N
34	2	2	Saturday	t	2025-10-16 02:29:29.715813+00	2025-10-16 02:39:59.416201+00	\N
35	2	2	Sunday	t	2025-10-16 02:29:29.715813+00	2025-10-16 02:40:00.150207+00	\N
59	3	3	Wednesday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.228161+00	\N
53	3	2	Thursday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.333867+00	\N
60	3	3	Thursday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.340855+00	\N
51	3	2	Tuesday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.021732+00	\N
46	3	1	Thursday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.325535+00	\N
58	3	3	Tuesday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.064602+00	\N
54	3	2	Friday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.541724+00	\N
61	3	3	Friday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.583978+00	\N
55	3	2	Saturday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.610382+00	\N
48	3	1	Saturday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.605034+00	\N
62	3	3	Saturday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.638864+00	\N
49	3	1	Sunday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.840993+00	\N
56	3	2	Sunday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.861168+00	\N
63	3	3	Sunday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.871843+00	\N
29	2	2	Monday	t	2025-10-16 02:29:29.715813+00	2025-10-16 13:34:06.944269+00	\N
45	3	1	Wednesday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.077935+00	\N
47	3	1	Friday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.348921+00	\N
36	2	3	Monday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:56.72579+00	\N
37	2	3	Tuesday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:56.734856+00	\N
23	2	1	Tuesday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:56.735835+00	\N
24	2	1	Wednesday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:56.939403+00	\N
31	2	2	Wednesday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:56.9922+00	\N
32	2	2	Thursday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:57.008564+00	\N
39	2	3	Thursday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:57.130542+00	\N
26	2	1	Friday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:57.199139+00	\N
33	2	2	Friday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:57.2095+00	\N
40	2	3	Friday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:57.262729+00	\N
27	2	1	Saturday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:57.328764+00	\N
41	2	3	Saturday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:57.383521+00	\N
28	2	1	Sunday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:57.480864+00	\N
42	2	3	Sunday	t	2025-10-16 02:29:29.715813+00	2025-10-17 01:53:57.531456+00	\N
57	3	3	Monday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:21.814446+00	\N
52	3	2	Wednesday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:22.078188+00	\N
102	5	3	Thursday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
103	5	3	Friday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
104	5	3	Saturday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
105	5	3	Sunday	f	2025-10-16 02:29:29.715813+00	2025-10-16 02:29:29.715813+00	\N
22	2	1	Monday	t	2025-10-16 02:29:29.715813+00	2025-10-16 02:39:51.635071+00	\N
136	2	4	Wednesday	t	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:54.495609+00	\N
134	2	4	Monday	t	2025-10-16 13:35:37.984037+00	2025-10-17 01:53:56.748562+00	\N
135	2	4	Tuesday	t	2025-10-16 13:35:37.984037+00	2025-10-17 01:53:56.930805+00	\N
137	2	4	Thursday	t	2025-10-16 13:35:37.984037+00	2025-10-17 01:53:57.136192+00	\N
138	2	4	Friday	t	2025-10-16 13:35:37.984037+00	2025-10-17 01:53:57.282194+00	\N
139	2	4	Saturday	t	2025-10-16 13:35:37.984037+00	2025-10-17 01:53:57.474192+00	\N
140	2	4	Sunday	t	2025-10-16 13:35:37.984037+00	2025-10-17 01:53:57.546895+00	\N
141	3	4	Monday	t	2025-10-16 13:35:37.984037+00	2025-10-17 03:08:21.812245+00	\N
120	6	3	Monday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:12.584129+00	\N
114	6	2	Tuesday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:12.585306+00	\N
106	6	1	Monday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:12.585863+00	\N
107	6	1	Tuesday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:12.586547+00	\N
113	6	2	Monday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:12.593709+00	\N
121	6	3	Tuesday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:12.593822+00	\N
115	6	2	Wednesday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:12.933861+00	\N
108	6	1	Wednesday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:12.935167+00	\N
109	6	1	Thursday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:12.938725+00	\N
122	6	3	Wednesday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:12.940188+00	\N
116	6	2	Thursday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:12.952863+00	\N
123	6	3	Thursday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:12.957539+00	\N
117	6	2	Friday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:13.285492+00	\N
111	6	1	Saturday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:13.291247+00	\N
124	6	3	Friday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:13.293477+00	\N
110	6	1	Friday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:13.302898+00	\N
118	6	2	Saturday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:13.313627+00	\N
125	6	3	Saturday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:13.315246+00	\N
119	6	2	Sunday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:13.63539+00	\N
112	6	1	Sunday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:13.636071+00	\N
126	6	3	Sunday	t	2025-10-16 02:29:29.715813+00	2025-10-16 03:20:13.646006+00	\N
127	1	4	Monday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
128	1	4	Tuesday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
129	1	4	Wednesday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
130	1	4	Thursday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
131	1	4	Friday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
132	1	4	Saturday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
133	1	4	Sunday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
148	4	4	Monday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
149	4	4	Tuesday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
150	4	4	Wednesday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
151	4	4	Thursday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
152	4	4	Friday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
153	4	4	Saturday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
154	4	4	Sunday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
155	5	4	Monday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
156	5	4	Tuesday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
157	5	4	Wednesday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
158	5	4	Thursday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
159	5	4	Friday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
160	5	4	Saturday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
161	5	4	Sunday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
162	6	4	Monday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
163	6	4	Tuesday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
164	6	4	Wednesday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
165	6	4	Thursday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
166	6	4	Friday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
43	3	1	Monday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:21.812557+00	\N
50	3	2	Monday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:21.814746+00	\N
44	3	1	Tuesday	t	2025-10-16 02:29:29.715813+00	2025-10-17 03:08:21.825034+00	\N
142	3	4	Tuesday	t	2025-10-16 13:35:37.984037+00	2025-10-17 03:08:22.073588+00	\N
143	3	4	Wednesday	t	2025-10-16 13:35:37.984037+00	2025-10-17 03:08:22.284798+00	\N
144	3	4	Thursday	t	2025-10-16 13:35:37.984037+00	2025-10-17 03:08:22.433745+00	\N
145	3	4	Friday	t	2025-10-16 13:35:37.984037+00	2025-10-17 03:08:22.599922+00	\N
146	3	4	Saturday	t	2025-10-16 13:35:37.984037+00	2025-10-17 03:08:22.799124+00	\N
147	3	4	Sunday	t	2025-10-16 13:35:37.984037+00	2025-10-17 03:08:22.871619+00	\N
167	6	4	Saturday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
168	6	4	Sunday	f	2025-10-16 13:35:37.984037+00	2025-10-16 13:35:37.984037+00	\N
169	7	1	Monday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
170	7	1	Tuesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
171	7	1	Wednesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
172	7	1	Thursday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
173	7	1	Friday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
174	7	1	Saturday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
175	7	1	Sunday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
176	7	2	Monday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
177	7	2	Tuesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
178	7	2	Wednesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
179	7	2	Thursday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
180	7	2	Friday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
181	7	2	Saturday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
182	7	2	Sunday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
183	7	3	Monday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
184	7	3	Tuesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
185	7	3	Wednesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
186	7	3	Thursday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
187	7	3	Friday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
188	7	3	Saturday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
189	7	3	Sunday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
190	8	1	Monday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
191	8	1	Tuesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
192	8	1	Wednesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
193	8	1	Thursday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
194	8	1	Friday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
195	8	1	Saturday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
196	8	1	Sunday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
197	8	2	Monday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
198	8	2	Tuesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
199	8	2	Wednesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
200	8	2	Thursday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
201	8	2	Friday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
202	8	2	Saturday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
203	8	2	Sunday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
204	8	3	Monday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
205	8	3	Tuesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
206	8	3	Wednesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
207	8	3	Thursday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
208	8	3	Friday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
209	8	3	Saturday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
210	8	3	Sunday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
254	11	1	Tuesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
256	11	1	Thursday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
257	11	1	Friday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
258	11	1	Saturday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
259	11	1	Sunday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
263	11	2	Thursday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
264	11	2	Friday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
265	11	2	Saturday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
266	11	2	Sunday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
267	11	3	Monday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
260	11	2	Monday	t	2025-10-18 03:24:01.02804+00	2025-10-18 03:37:57.881867+00	\N
255	11	1	Wednesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:38:08.290049+00	\N
261	11	2	Tuesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:38:09.249183+00	\N
262	11	2	Wednesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:38:07.553222+00	\N
213	9	1	Wednesday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.341749+00	\N
220	9	2	Wednesday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.345848+00	\N
227	9	3	Wednesday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.347949+00	\N
214	9	1	Thursday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.350588+00	\N
215	9	1	Friday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.403373+00	\N
222	9	2	Friday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.414649+00	\N
230	9	3	Saturday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.428737+00	\N
221	9	2	Thursday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.355358+00	\N
216	9	1	Saturday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.419079+00	\N
229	9	3	Friday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.416838+00	\N
217	9	1	Sunday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.465378+00	\N
224	9	2	Sunday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.481755+00	\N
211	9	1	Monday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:15.79775+00	\N
231	9	3	Sunday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.48453+00	\N
225	9	3	Monday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.27784+00	\N
218	9	2	Monday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.280085+00	\N
223	9	2	Saturday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.420312+00	\N
219	9	2	Tuesday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.281544+00	\N
246	10	3	Monday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:12.932366+00	\N
233	10	1	Tuesday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:12.9333+00	\N
247	10	3	Tuesday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:12.934839+00	\N
239	10	2	Monday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:12.936993+00	\N
240	10	2	Tuesday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.132418+00	\N
232	10	1	Monday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.13999+00	\N
234	10	1	Wednesday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.180774+00	\N
241	10	2	Wednesday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.192224+00	\N
235	10	1	Thursday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.193859+00	\N
242	10	2	Thursday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.202931+00	\N
249	10	3	Thursday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.331073+00	\N
248	10	3	Wednesday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.347553+00	\N
236	10	1	Friday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.434311+00	\N
250	10	3	Friday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.451631+00	\N
243	10	2	Friday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.452661+00	\N
237	10	1	Saturday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.470885+00	\N
244	10	2	Saturday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.535746+00	\N
251	10	3	Saturday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.555477+00	\N
238	10	1	Sunday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.690911+00	\N
245	10	2	Sunday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.712347+00	\N
252	10	3	Sunday	t	2025-10-18 03:24:01.02804+00	2025-10-20 03:04:13.715638+00	\N
269	11	3	Wednesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
270	11	3	Thursday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
271	11	3	Friday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
272	11	3	Saturday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
273	11	3	Sunday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:24:01.02804+00	\N
286	12	2	Saturday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:52.251885+00	\N
253	11	1	Monday	t	2025-10-18 03:24:01.02804+00	2025-10-18 03:37:52.262499+00	\N
293	12	3	Saturday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:52.270064+00	\N
268	11	3	Tuesday	f	2025-10-18 03:24:01.02804+00	2025-10-18 03:38:09.950985+00	\N
316	15	1	Monday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
317	15	1	Tuesday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
318	15	1	Wednesday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
319	15	1	Thursday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
320	15	1	Friday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
321	15	1	Saturday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
322	15	1	Sunday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
323	15	2	Monday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
324	15	2	Tuesday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
325	15	2	Wednesday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
326	15	2	Thursday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
327	15	2	Friday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
328	15	2	Saturday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
329	15	2	Sunday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
330	15	3	Monday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
331	15	3	Tuesday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
332	15	3	Wednesday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
333	15	3	Thursday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
334	15	3	Friday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
335	15	3	Saturday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
336	15	3	Sunday	f	2025-10-18 11:00:32.205183+00	2025-10-18 11:00:32.205183+00	\N
288	12	3	Monday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:51.715792+00	\N
274	12	1	Monday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:51.71638+00	\N
275	12	1	Tuesday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:51.72201+00	\N
282	12	2	Tuesday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:51.725222+00	\N
281	12	2	Monday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:51.728024+00	\N
289	12	3	Tuesday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:51.732788+00	\N
276	12	1	Wednesday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:51.977652+00	\N
290	12	3	Wednesday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:51.990015+00	\N
291	12	3	Thursday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:51.991524+00	\N
283	12	2	Wednesday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:51.991796+00	\N
277	12	1	Thursday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:51.9951+00	\N
284	12	2	Thursday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:51.995186+00	\N
278	12	1	Friday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:52.237213+00	\N
285	12	2	Friday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:52.240818+00	\N
292	12	3	Friday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:52.24709+00	\N
279	12	1	Saturday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:52.251599+00	\N
294	12	3	Sunday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:52.512535+00	\N
287	12	2	Sunday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:52.51264+00	\N
280	12	1	Sunday	t	2025-10-18 11:00:32.205183+00	2025-10-20 03:06:52.531051+00	\N
337	11	1	Monday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
338	11	2	Monday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
339	11	3	Monday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
340	11	1	Tuesday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
341	11	2	Tuesday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
342	11	3	Tuesday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
343	11	1	Wednesday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
344	11	2	Wednesday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
345	11	3	Wednesday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
346	11	1	Thursday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
347	11	2	Thursday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
348	11	3	Thursday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
349	11	1	Friday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
350	11	2	Friday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
351	11	3	Friday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
352	11	1	Saturday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
353	11	2	Saturday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
354	11	3	Saturday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
355	11	1	Sunday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
356	11	2	Sunday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
357	11	3	Sunday	f	2025-10-20 03:27:41.788216+00	2025-10-20 03:27:41.788216+00	\N
304	14	2	Wednesday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.494745+00	\N
302	14	2	Monday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.42808+00	\N
311	14	3	Wednesday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.49706+00	\N
310	14	3	Tuesday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.453212+00	\N
298	14	1	Thursday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.498775+00	\N
305	14	2	Thursday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.509262+00	\N
312	14	3	Thursday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.524367+00	\N
306	14	2	Friday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.564377+00	\N
299	14	1	Friday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.566129+00	\N
313	14	3	Friday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.568125+00	\N
300	14	1	Saturday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.570414+00	\N
307	14	2	Saturday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.581394+00	\N
314	14	3	Saturday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.595304+00	\N
308	14	2	Sunday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.631757+00	\N
315	14	3	Sunday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.637818+00	\N
301	14	1	Sunday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.634054+00	\N
309	14	3	Monday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.425186+00	\N
295	14	1	Monday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.426871+00	\N
303	14	2	Tuesday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.428361+00	\N
297	14	1	Wednesday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.497172+00	\N
358	12	1	Monday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
359	12	2	Monday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
360	12	3	Monday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
361	12	1	Tuesday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
362	12	2	Tuesday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
363	12	3	Tuesday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
226	9	3	Tuesday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.289494+00	\N
364	12	1	Wednesday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
365	12	2	Wednesday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
366	12	3	Wednesday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
367	12	1	Thursday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
368	12	2	Thursday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
369	12	3	Thursday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
370	12	1	Friday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
371	12	2	Friday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
372	12	3	Friday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
373	12	1	Saturday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
374	12	2	Saturday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
375	12	3	Saturday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
376	12	1	Sunday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
377	12	2	Sunday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
378	12	3	Sunday	f	2025-10-20 03:38:12.658075+00	2025-10-20 03:38:12.658075+00	\N
379	12	1	Monday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
380	12	2	Monday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
381	12	3	Monday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
382	12	1	Tuesday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
383	12	2	Tuesday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
384	12	3	Tuesday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
385	12	1	Wednesday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
386	12	2	Wednesday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
387	12	3	Wednesday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
388	12	1	Thursday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
389	12	2	Thursday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
390	12	3	Thursday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
391	12	1	Friday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
392	12	2	Friday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
393	12	3	Friday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
394	12	1	Saturday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
395	12	2	Saturday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
396	12	3	Saturday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
397	12	1	Sunday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
398	12	2	Sunday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
399	12	3	Sunday	f	2025-10-20 03:41:01.632275+00	2025-10-20 03:41:01.632275+00	\N
212	9	1	Tuesday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.289923+00	\N
228	9	3	Thursday	t	2025-10-18 03:24:01.02804+00	2025-10-24 14:49:14.359649+00	\N
400	7	5	Monday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
401	7	5	Tuesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
402	7	5	Wednesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
403	7	5	Thursday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
404	7	5	Friday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
405	7	5	Saturday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
406	7	5	Sunday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
407	8	5	Monday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
408	8	5	Tuesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
409	8	5	Wednesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
410	8	5	Thursday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
411	8	5	Friday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
412	8	5	Saturday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
413	8	5	Sunday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
414	9	5	Monday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
415	9	5	Tuesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
416	9	5	Wednesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
417	9	5	Thursday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
418	9	5	Friday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
419	9	5	Saturday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
420	9	5	Sunday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
421	10	5	Monday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
422	10	5	Tuesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
423	10	5	Wednesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
424	10	5	Thursday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
425	10	5	Friday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
426	10	5	Saturday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
427	10	5	Sunday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
428	11	5	Monday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
429	11	5	Tuesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
430	11	5	Wednesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
431	11	5	Thursday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
432	11	5	Friday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
433	11	5	Saturday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
434	11	5	Sunday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
435	12	5	Monday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
436	12	5	Tuesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
437	12	5	Wednesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
438	12	5	Thursday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
439	12	5	Friday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
440	12	5	Saturday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
441	12	5	Sunday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
442	14	5	Monday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
443	14	5	Tuesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
444	14	5	Wednesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
445	14	5	Thursday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
446	14	5	Friday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
447	14	5	Saturday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
448	14	5	Sunday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
449	15	5	Monday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
450	15	5	Tuesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
451	15	5	Wednesday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
452	15	5	Thursday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
453	15	5	Friday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
454	15	5	Saturday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
455	15	5	Sunday	f	2025-10-29 05:23:44.981166+00	2025-10-29 05:23:44.981166+00	\N
296	14	1	Tuesday	f	2025-10-18 11:00:32.205183+00	2025-10-29 06:04:28.43711+00	\N
\.


--
-- Data for Name: bookings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bookings (id, customer_name, phone_number, email, booking_date, booking_time, total_persons, status, memo, created_at, updated_at, deleted_at) FROM stdin;
1	Nguyễn Hữu Nam	0123456789	nam123@gmail.com	2025/10/04	20:30	4	confirmed	Sinh nhật	2025-10-16 02:28:09.603731+00	2025-10-18 03:50:19.135468+00	\N
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, name, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customers (id, full_name, phone_number, created_at, updated_at, deleted_at) FROM stdin;
1	Nguyễn Hữu Nam	0123456789	2025-10-16 02:28:10.084126+00	2025-10-16 02:28:10.084126+00	\N
\.


--
-- Data for Name: employees; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.employees (id, full_name, gender, birthday, phone_number, email, schedule_type, address, join_date, base_salary, salary_per_hour, account_id, avatar_file_id, created_at, updated_at, deleted_at) FROM stdin;
1	Lê Đỗ Thanh Sang	t	2002-09-30T17:00:00.000Z	0123456789	sang@gmail.com	FULLTIME	98/2 Đoàn Thị Kia	2025-09-30T17:00:00.000Z	5000000	0	2	\N	2025-10-16 01:33:30.872184+00	2025-10-16 01:33:30.872184+00	2025-10-17 15:51:47.856254+00
2	Nguyễn Thành Công	t	2000-10-01T17:00:00.000Z	0123456789	cong@gmail.com	FULLTIME	Nhà trọ Phạm Tiến Đô, Đường DJ9, Tổ 14, khu phố 3b, Phường Thới Hòa, Thành phố Hồ Chí Minh, Phường Bến Cát	2025-09-30T17:00:00.000Z	10000000	0	3	\N	2025-10-16 02:19:47.396654+00	2025-10-16 02:21:42.858622+00	2025-10-18 02:38:16.628971+00
3	Đỗ Thị Lan	f	2003-09-28T17:00:00.000Z	0456789123	lan@gmail.com	PARTTIME	DJ5, khu phố 3B, Thới Hòa, Bến Cát, Bình Dương	2025-09-30T17:00:00.000Z	4000000	24000	4	\N	2025-10-16 02:20:59.518509+00	2025-10-16 02:20:59.518509+00	2025-10-18 02:38:33.582995+00
4	Nguyễn Thị Ninh	f		0456789123	ninh@gmail.com	FULLTIME	123 Lê Lợi, Phường Bến Thành	2025-09-30T17:00:00.000Z	5000000	0	5	\N	2025-10-16 02:22:36.167071+00	2025-10-16 02:23:11.787759+00	2025-10-18 02:38:35.583197+00
5	Phan Minh Huy	t	2004-08-01T17:00:00.000Z	0359088784	huy5902@gmail.com	PARTTIME	Thoi Hoa, Ben Cat	2025-09-30T17:00:00.000Z	3000000	0	6	\N	2025-10-16 02:24:04.738841+00	2025-10-16 02:24:04.738841+00	2025-10-18 02:38:37.419022+00
6	Trần Thành Nam	t	2004-09-26T17:00:00.000Z	0147258369	nam@gmail.com	FULLTIME		2025-09-30T17:00:00.000Z	15000000	0	7	\N	2025-10-16 02:24:55.376143+00	2025-10-17 03:07:47.795815+00	2025-10-18 02:38:39.074735+00
8	Trần Quang Dũng	t	2003-07-10T17:00:00.000Z	094638512	Dungga@gmail.com	FULLTIME	Bình Dương	2025-10-17T17:00:00.000Z	4000000	20000	10	\N	2025-10-18 02:46:36.741857+00	2025-10-18 02:46:59.932619+00	\N
9	Trần Bảo Ninh	t	2003-11-20T17:00:00.000Z	091547635	Ninh123@gmail.com	FULLTIME	Đăk Nông	2025-10-17T17:00:00.000Z	4000000	20000	11	\N	2025-10-18 02:50:21.883523+00	2025-10-18 02:50:21.883523+00	\N
10	Hà Kim Ngân	f	2003-05-25T17:00:00.000Z	094712354	Ngan2605@gmail.com	FULLTIME	Bình Phước	2025-10-17T17:00:00.000Z	5000000	25000	12	\N	2025-10-18 02:52:38.728014+00	2025-10-18 02:52:38.728014+00	\N
11	Dương Gia Huy	t	2003-05-13T17:00:00.000Z	094563895	Huy321@gmail.com	FULLTIME	Bình Phước	2025-10-17T17:00:00.000Z	4000000	20000	13	\N	2025-10-18 02:54:50.009031+00	2025-10-18 02:54:50.009031+00	\N
12	Bùi Tùng Lâm	t	2003-10-03T17:00:00.000Z	094876556	Lamlori03@gmail.com	PARTTIME	Quảng Ngãi	2025-10-17T17:00:00.000Z	4000000	20000	14	\N	2025-10-18 03:41:41.033501+00	2025-10-18 03:41:58.698423+00	\N
13	Trần Quang Việt	t	2003-08-12T17:00:00.000Z	0965954249	qzit1308@gmail.com	FULLTIME	Quảng Ngãi	2025-10-17T17:00:00.000Z	10000000	35000	15	\N	2025-10-18 03:45:06.713067+00	2025-10-18 03:45:06.713067+00	2025-10-18 03:45:21.077486+00
7	Trần Quang Việt	t	2003-08-12T17:00:00.000Z	0965954249	qzit1308@gmail.com	FULLTIME	Quảng Ngãi\n	2025-10-17T17:00:00.000Z	6000000	25000	9	\N	2025-10-18 02:44:36.207248+00	2025-10-18 03:46:01.317771+00	\N
14	Trần Quang Cường	t	2003-06-22T17:00:00.000Z	094876612	cuongdz@gmail.com	PARTTIME	Quảng Ngãi	2025-10-17T17:00:00.000Z	4000000	20000	16	\N	2025-10-18 03:47:10.654726+00	2025-10-18 03:47:10.654726+00	\N
15	Nguyễn Xuân Quỳnh	t	2003-04-07T17:00:00.000Z	094633118	quynhari600@gmail.com	PARTTIME	Quảng Ngãi	2025-10-17T17:00:00.000Z	4000000	200000	17	\N	2025-10-18 03:48:52.143818+00	2025-10-18 03:48:52.143818+00	\N
\.


--
-- Data for Name: files; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.files (id, file_name, url, mime_type, size, public_id, resource_type, folder, created_at, updated_at, deleted_at) FROM stdin;
1	Bánh hỏi heo quay.webp	https://res.cloudinary.com/dbuulcdnd/image/upload/v1760578661/uploads/uploads/1760578659_B%C3%A1nh%20h%E1%BB%8Fi%20heo%20quay.webp.webp	image/webp	32186	uploads/uploads/1760578659_Bánh hỏi heo quay.webp	image		2025-10-16 01:37:41.686012+00	2025-10-16 01:37:41.686012+00	\N
2	chả giò.jpg	https://res.cloudinary.com/dbuulcdnd/image/upload/v1760578663/uploads/uploads/1760578662_ch%E1%BA%A3%20gi%C3%B2.jpg.jpg	image/jpeg	116029	uploads/uploads/1760578662_chả giò.jpg	image		2025-10-16 01:37:43.824625+00	2025-10-16 01:37:43.824625+00	\N
3	download (1).webp	https://res.cloudinary.com/dbuulcdnd/image/upload/v1760578664/uploads/uploads/1760578664_download%20%281%29.webp.webp	image/webp	6976	uploads/uploads/1760578664_download (1).webp	image		2025-10-16 01:37:44.823876+00	2025-10-16 01:37:44.823876+00	\N
4	OIF.jpg	https://res.cloudinary.com/dbuulcdnd/image/upload/v1760578665/uploads/uploads/1760578665_OIF.jpg.jpg	image/jpeg	16865	uploads/uploads/1760578665_OIF.jpg	image		2025-10-16 01:37:46.330877+00	2025-10-16 01:37:46.330877+00	\N
5	OIP (1).webp	https://res.cloudinary.com/dbuulcdnd/image/upload/v1760578666/uploads/uploads/1760578666_OIP%20%281%29.webp.webp	image/webp	9762	uploads/uploads/1760578666_OIP (1).webp	image		2025-10-16 01:37:47.579733+00	2025-10-16 01:37:47.579733+00	\N
6	OIP.jpg	https://res.cloudinary.com/dbuulcdnd/image/upload/v1760578668/uploads/uploads/1760578667_OIP.jpg.jpg	image/jpeg	10002	uploads/uploads/1760578667_OIP.jpg	image		2025-10-16 01:37:48.675287+00	2025-10-16 01:37:48.675287+00	\N
7	OIP.webp	https://res.cloudinary.com/dbuulcdnd/image/upload/v1760578669/uploads/uploads/1760578668_OIP.webp.webp	image/webp	11400	uploads/uploads/1760578668_OIP.webp	image		2025-10-16 01:37:49.813708+00	2025-10-16 01:37:49.813708+00	\N
8	th.jpg	https://res.cloudinary.com/dbuulcdnd/image/upload/v1760578670/uploads/uploads/1760578670_th.jpg.jpg	image/jpeg	12486	uploads/uploads/1760578670_th.jpg	image		2025-10-16 01:37:51.157772+00	2025-10-16 01:37:51.157772+00	\N
9	shopping.webp	https://res.cloudinary.com/dbuulcdnd/image/upload/v1760761193/uploads/uploads/1760761192_shopping.webp.webp	image/webp	2828	uploads/uploads/1760761192_shopping.webp	image		2025-10-18 04:19:54.016273+00	2025-10-18 04:19:54.016273+00	\N
\.


--
-- Data for Name: ingredients; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ingredients (id, name, description, quantity, warning_quantity, supplier, unit, created_at, updated_at, deleted_at) FROM stdin;
2	Bắp cải		20	5	Rau sạch	cai	2025-10-16 02:41:39.847749+00	2025-10-16 02:41:39.847749+00	2025-10-16 02:43:03.472113+00
3	Bắp cải		20	5	Rau sạch	cai	2025-10-16 02:41:55.927025+00	2025-10-16 02:41:55.927025+00	2025-10-16 02:43:06.169182+00
4	Bắp cải		20	5	Rau sạch	cai	2025-10-16 02:42:50.309282+00	2025-10-16 02:42:50.309282+00	2025-10-16 02:43:08.103914+00
6	Thịt heo		100	20	Thịt tươi	kg	2025-10-16 02:43:47.393554+00	2025-10-16 02:43:47.393554+00	\N
7	Cá chép		20	5	Cá tươi	con	2025-10-16 02:44:35.990021+00	2025-10-16 02:44:35.990021+00	\N
8	Tỏi		50	10	Gia vị	kg	2025-10-16 02:45:13.447152+00	2025-10-16 02:45:13.447152+00	\N
9	Hành tím		50	10	Gia vị	kg	2025-10-16 02:45:17.035555+00	2025-10-16 02:45:17.035555+00	\N
10	Tiêu		50	10	Gia vị	kg	2025-10-16 02:45:22.687973+00	2025-10-16 02:45:22.687973+00	\N
11	Ớt chuông		50	10	Gia vị	kg	2025-10-16 02:45:36.105873+00	2025-10-16 02:45:36.105873+00	\N
12	Dưa muối		50	10	Gia vị	kg	2025-10-16 02:45:41.409266+00	2025-10-16 02:45:41.409266+00	\N
13	Gạo		100	10	Gia vị	kg	2025-10-16 02:46:12.996506+00	2025-10-16 02:46:12.996506+00	\N
14	Bún		100	10	Gia vị	kg	2025-10-16 02:46:34.277282+00	2025-10-16 02:46:34.277282+00	\N
15	Dưa hấu		20	1	Trái cây	qua	2025-10-16 02:47:05.728345+00	2025-10-16 02:47:05.728345+00	\N
16	Cam		20	1	Trái cây	qua	2025-10-16 02:47:08.023045+00	2025-10-16 02:47:08.023045+00	\N
17	Mít		20	1	Trái cây	qua	2025-10-16 02:47:19.036492+00	2025-10-16 02:47:19.036492+00	\N
18	Hàu sữa		50	20	Hải sản	kg	2025-10-16 02:47:43.347302+00	2025-10-16 02:47:43.347302+00	\N
19	Bạch tuộc		50	20	Hải sản	kg	2025-10-16 02:47:51.482246+00	2025-10-16 02:47:51.482246+00	\N
20	Mực ống		50	20	Hải sản	kg	2025-10-16 02:47:56.814677+00	2025-10-16 02:47:56.814677+00	\N
21	Cá đuối		50	20	Hải sản	kg	2025-10-16 02:48:03.78642+00	2025-10-16 02:48:03.78642+00	\N
22	Ốc hương		50	20	Hải sản	kg	2025-10-16 02:48:16.301436+00	2025-10-16 02:48:16.301436+00	\N
5	Thịt bò		110	20	Thịt tươi	kg	2025-10-16 02:43:38.05121+00	2025-10-16 02:56:39.059891+00	\N
1	Bắp cải		18	5	Rau sạch	cai	2025-10-16 02:41:18.779147+00	2025-10-16 02:56:54.899984+00	\N
24	nước suối		10	2	aquafina	thung	2025-10-18 03:14:04.724227+00	2025-10-18 03:14:04.724227+00	2025-10-18 03:14:25.244094+00
23	nước suối		10	2	aquafina	thung	2025-10-18 03:14:02.967671+00	2025-10-18 03:14:02.967671+00	2025-10-18 03:14:29.751097+00
25	nước suối	nước uống	7	2	aquafina	thung	2025-10-18 03:14:16.171193+00	2025-10-18 03:15:38.322209+00	\N
\.


--
-- Data for Name: menu_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.menu_items (id, name, description, price, category_id, file_id, status, created_at, updated_at, deleted_at) FROM stdin;
1	Cá chép chiên giòn		200000	\N	4	Available	2025-10-16 02:48:55.277194+00	2025-10-16 02:48:55.277194+00	\N
2	Bánh hỏi heo quay		159000	\N	1	Available	2025-10-16 02:49:17.32517+00	2025-10-16 02:49:17.32517+00	\N
3	Chả chiên giòn		150000	\N	2	Available	2025-10-16 02:49:34.399581+00	2025-10-16 02:49:34.399581+00	\N
4	Gỏi rau trộn		50000	\N	8	Available	2025-10-16 02:49:51.384527+00	2025-10-16 02:49:51.384527+00	\N
5	Xôi ngũ vị		59000	\N	5	Available	2025-10-16 02:50:07.106628+00	2025-10-16 02:50:07.106628+00	\N
6	Xiên heo nướng 		99000	\N	7	Available	2025-10-16 02:50:25.587057+00	2025-10-16 02:50:25.587057+00	\N
7	nước suối	nước uống	20000	\N	9	Available	2025-10-18 04:20:39.029053+00	2025-10-18 04:20:39.029053+00	\N
\.


--
-- Data for Name: order_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.order_items (id, order_id, menu_item_id, quantity, amount, memo, status, created_at, updated_at, deleted_at) FROM stdin;
2	1	1	1	200000		Done	2025-10-16 02:52:36.268926+00	2025-10-16 02:56:17.238529+00	\N
1	1	5	1	59000		Done	2025-10-16 02:52:36.26126+00	2025-10-16 02:56:18.278438+00	\N
3	2	6	1	99000		Done	2025-10-16 03:05:09.135706+00	2025-10-16 13:23:05.595841+00	\N
6	2	5	1	59000		Done	2025-10-16 03:05:09.341563+00	2025-10-16 13:37:48.661354+00	\N
5	2	2	1	159000		Done	2025-10-16 03:05:09.140691+00	2025-10-18 03:19:02.057161+00	\N
4	2	1	1	200000		Done	2025-10-16 03:05:09.139922+00	2025-10-18 03:19:03.47219+00	\N
8	3	4	1	50000		Done	2025-10-18 03:18:31.983581+00	2025-10-29 05:27:40.646652+00	\N
7	3	3	1	150000		Done	2025-10-18 03:18:31.982143+00	2025-10-29 05:28:36.317915+00	\N
11	4	2	1	159000		Pending	2025-10-29 05:30:47.232577+00	2025-10-29 05:30:47.232577+00	\N
12	4	7	1	20000		Pending	2025-10-29 05:30:47.244283+00	2025-10-29 05:30:47.244283+00	\N
13	4	5	1	59000		Pending	2025-10-29 05:32:28.044145+00	2025-10-29 05:32:28.044145+00	\N
9	3	5	1	59000		Done	2025-10-18 03:18:32.038452+00	2025-10-29 05:36:37.852454+00	\N
10	4	3	1	150000		Done	2025-10-29 05:30:47.230423+00	2025-10-29 05:45:00.932627+00	\N
16	5	7	1	20000		Done	2025-10-29 05:46:48.422314+00	2025-10-29 05:46:58.301421+00	\N
15	5	4	1	50000		Done	2025-10-29 05:46:48.41893+00	2025-10-29 05:46:59.503402+00	\N
14	5	2	1	159000		Done	2025-10-29 05:46:48.352037+00	2025-10-29 05:47:01.446716+00	\N
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, customer_id, table_id, amount, status, memo, created_at, updated_at, deleted_at) FROM stdin;
1	\N	1	259000	Paid		2025-10-16 02:52:35.784811+00	2025-10-16 02:53:01.756752+00	\N
2	\N	3	517000	Paid		2025-10-16 03:05:08.721683+00	2025-10-16 03:05:14.238277+00	\N
3	\N	2	259000	UnPaid		2025-10-18 03:18:31.866396+00	2025-10-18 03:18:31.866396+00	\N
4	\N	1	388000	Paid		2025-10-29 05:30:47.085016+00	2025-10-29 05:32:33.450877+00	\N
5	\N	3	229000	Paid		2025-10-29 05:46:48.204294+00	2025-10-29 05:47:08.676901+00	\N
\.


--
-- Data for Name: permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.permissions (id, name, created_at, updated_at, deleted_at) FROM stdin;
1	view_accounts	2025-10-16 01:27:44.167655+00	2025-10-16 01:27:44.167655+00	\N
2	create_account	2025-10-16 01:27:44.168626+00	2025-10-16 01:27:44.168626+00	\N
3	edit_account	2025-10-16 01:27:44.169235+00	2025-10-16 01:27:44.169235+00	\N
4	delete_account	2025-10-16 01:27:44.169746+00	2025-10-16 01:27:44.169746+00	\N
5	view_employees	2025-10-16 01:27:44.170433+00	2025-10-16 01:27:44.170433+00	\N
6	create_employee	2025-10-16 01:27:44.171119+00	2025-10-16 01:27:44.171119+00	\N
7	edit_employee	2025-10-16 01:27:44.17162+00	2025-10-16 01:27:44.17162+00	\N
8	delete_employee	2025-10-16 01:27:44.172147+00	2025-10-16 01:27:44.172147+00	\N
9	view_shifts	2025-10-16 01:27:44.172777+00	2025-10-16 01:27:44.172777+00	\N
10	create_shift	2025-10-16 01:27:44.173408+00	2025-10-16 01:27:44.173408+00	\N
11	edit_shift	2025-10-16 01:27:44.174036+00	2025-10-16 01:27:44.174036+00	\N
12	delete_shift	2025-10-16 01:27:44.17608+00	2025-10-16 01:27:44.17608+00	\N
13	view_shift_schedules	2025-10-16 01:27:44.176673+00	2025-10-16 01:27:44.176673+00	\N
14	create_shift_schedule	2025-10-16 01:27:44.17725+00	2025-10-16 01:27:44.17725+00	\N
15	edit_shift_schedule	2025-10-16 01:27:44.17781+00	2025-10-16 01:27:44.17781+00	\N
16	delete_shift_schedule	2025-10-16 01:27:44.179175+00	2025-10-16 01:27:44.179175+00	\N
17	view_menu_items	2025-10-16 01:27:44.181019+00	2025-10-16 01:27:44.181019+00	\N
18	create_menu_item	2025-10-16 01:27:44.18167+00	2025-10-16 01:27:44.18167+00	\N
19	edit_menu_item	2025-10-16 01:27:44.182302+00	2025-10-16 01:27:44.182302+00	\N
20	delete_menu_item	2025-10-16 01:27:44.184081+00	2025-10-16 01:27:44.184081+00	\N
21	view_ingredients	2025-10-16 01:27:44.184639+00	2025-10-16 01:27:44.184639+00	\N
22	create_ingredient	2025-10-16 01:27:44.185207+00	2025-10-16 01:27:44.185207+00	\N
23	edit_ingredient	2025-10-16 01:27:44.187113+00	2025-10-16 01:27:44.187113+00	\N
24	delete_ingredient	2025-10-16 01:27:44.187712+00	2025-10-16 01:27:44.187712+00	\N
25	view_inventory	2025-10-16 01:27:44.188285+00	2025-10-16 01:27:44.188285+00	\N
26	edit_inventory	2025-10-16 01:27:44.188859+00	2025-10-16 01:27:44.188859+00	\N
27	view_orders	2025-10-16 01:27:44.190487+00	2025-10-16 01:27:44.190487+00	\N
28	create_order	2025-10-16 01:27:44.191089+00	2025-10-16 01:27:44.191089+00	\N
29	edit_order	2025-10-16 01:27:44.19312+00	2025-10-16 01:27:44.19312+00	\N
30	delete_order	2025-10-16 01:27:44.193711+00	2025-10-16 01:27:44.193711+00	\N
31	view_tables	2025-10-16 01:27:44.194299+00	2025-10-16 01:27:44.194299+00	\N
32	create_table	2025-10-16 01:27:44.194897+00	2025-10-16 01:27:44.194897+00	\N
33	edit_table	2025-10-16 01:27:44.196506+00	2025-10-16 01:27:44.196506+00	\N
34	delete_table	2025-10-16 01:27:44.1971+00	2025-10-16 01:27:44.1971+00	\N
35	view_customers	2025-10-16 01:27:44.199196+00	2025-10-16 01:27:44.199196+00	\N
36	create_customer	2025-10-16 01:27:44.199771+00	2025-10-16 01:27:44.199771+00	\N
37	edit_customer	2025-10-16 01:27:44.200388+00	2025-10-16 01:27:44.200388+00	\N
38	delete_customer	2025-10-16 01:27:44.200983+00	2025-10-16 01:27:44.200983+00	\N
39	view_roles	2025-10-16 01:27:44.202132+00	2025-10-16 01:27:44.202132+00	\N
40	create_role	2025-10-16 01:27:44.202715+00	2025-10-16 01:27:44.202715+00	\N
41	edit_role	2025-10-16 01:27:44.203323+00	2025-10-16 01:27:44.203323+00	\N
42	delete_role	2025-10-16 01:27:44.203875+00	2025-10-16 01:27:44.203875+00	\N
43	view_permissions	2025-10-16 01:27:44.204479+00	2025-10-16 01:27:44.204479+00	\N
44	view_tickets	2025-10-16 01:27:44.205087+00	2025-10-16 01:27:44.205087+00	\N
45	create_ticket	2025-10-16 01:27:44.205613+00	2025-10-16 01:27:44.205613+00	\N
46	edit_ticket	2025-10-16 01:27:44.206133+00	2025-10-16 01:27:44.206133+00	\N
47	delete_ticket	2025-10-16 01:27:44.206736+00	2025-10-16 01:27:44.206736+00	\N
48	view_attendance	2025-10-16 01:27:44.207422+00	2025-10-16 01:27:44.207422+00	\N
49	create_attendance	2025-10-16 01:27:44.208022+00	2025-10-16 01:27:44.208022+00	\N
50	edit_attendance	2025-10-16 01:27:44.208564+00	2025-10-16 01:27:44.208564+00	\N
51	delete_attendance	2025-10-16 01:27:44.209172+00	2025-10-16 01:27:44.209172+00	\N
52	view_files	2025-10-16 01:27:44.209793+00	2025-10-16 01:27:44.209793+00	\N
53	upload_file	2025-10-16 01:27:44.210442+00	2025-10-16 01:27:44.210442+00	\N
54	delete_file	2025-10-16 01:27:44.212117+00	2025-10-16 01:27:44.212117+00	\N
\.


--
-- Data for Name: role_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.role_permissions (permission_id, role_id) FROM stdin;
1	1
2	1
3	1
4	1
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roles (id, role_name, created_at, updated_at, deleted_at) FROM stdin;
2	Quản lý	2025-10-16 01:27:44.215588+00	2025-10-16 01:27:44.215588+00	\N
3	Thu ngân	2025-10-16 01:27:44.216199+00	2025-10-16 01:27:44.216199+00	\N
4	Phục vụ	2025-10-16 01:27:44.216827+00	2025-10-16 01:27:44.216827+00	\N
5	Đầu bếp	2025-10-16 01:27:44.21749+00	2025-10-16 01:27:44.21749+00	\N
6	Nhân viên bếp	2025-10-16 01:27:44.218123+00	2025-10-16 01:27:44.218123+00	\N
7	Nhân viên kho	2025-10-16 01:27:44.218689+00	2025-10-16 01:27:44.218689+00	\N
8	Khách hàng	2025-10-16 01:27:44.219286+00	2025-10-16 01:27:44.219286+00	\N
1	Admin	2025-10-16 01:27:44.213246+00	2025-10-20 03:59:26.659941+00	\N
\.


--
-- Data for Name: shift_schedules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shift_schedules (id, employee_id, shift_id, date, created_at, updated_at, deleted_at) FROM stdin;
1	2	1	2025-10-20	2025-10-16 02:40:15.166482+00	2025-10-16 02:40:15.166482+00	\N
2	2	2	2025-10-21	2025-10-16 02:40:15.173431+00	2025-10-16 02:40:15.173431+00	\N
3	2	3	2025-10-22	2025-10-16 02:40:15.180002+00	2025-10-16 02:40:15.180002+00	\N
4	2	1	2025-10-24	2025-10-16 02:40:15.18481+00	2025-10-16 02:40:15.18481+00	\N
5	2	2	2025-10-25	2025-10-16 02:40:15.189413+00	2025-10-16 02:40:15.189413+00	\N
6	6	1	2025-10-20	2025-10-16 03:20:48.137269+00	2025-10-16 03:20:48.137269+00	\N
7	6	2	2025-10-20	2025-10-16 03:20:48.142488+00	2025-10-16 03:20:48.142488+00	\N
154	3	1	2025-10-23	2025-10-18 03:23:55.255207+00	2025-10-18 03:23:55.255207+00	\N
9	6	1	2025-10-21	2025-10-16 03:20:48.153473+00	2025-10-16 03:20:48.153473+00	\N
10	6	2	2025-10-21	2025-10-16 03:20:48.158538+00	2025-10-16 03:20:48.158538+00	\N
155	3	1	2025-10-25	2025-10-18 03:23:55.326142+00	2025-10-18 03:23:55.326142+00	\N
12	6	1	2025-10-22	2025-10-16 03:20:48.167508+00	2025-10-16 03:20:48.167508+00	\N
13	6	2	2025-10-22	2025-10-16 03:20:48.171821+00	2025-10-16 03:20:48.171821+00	\N
156	3	1	2025-10-26	2025-10-18 03:23:55.366559+00	2025-10-18 03:23:55.366559+00	\N
15	6	1	2025-10-23	2025-10-16 03:20:48.18105+00	2025-10-16 03:20:48.18105+00	\N
16	6	2	2025-10-23	2025-10-16 03:20:48.184819+00	2025-10-16 03:20:48.184819+00	\N
17	2	1	2025-10-23	2025-10-16 03:20:48.188398+00	2025-10-16 03:20:48.188398+00	\N
18	6	1	2025-10-24	2025-10-16 03:20:48.192648+00	2025-10-16 03:20:48.192648+00	\N
19	6	2	2025-10-24	2025-10-16 03:20:48.196266+00	2025-10-16 03:20:48.196266+00	\N
20	6	1	2025-10-25	2025-10-16 03:20:48.200172+00	2025-10-16 03:20:48.200172+00	\N
21	6	2	2025-10-25	2025-10-16 03:20:48.204012+00	2025-10-16 03:20:48.204012+00	\N
157	11	1	2025-10-20	2025-10-19 02:00:58.185874+00	2025-10-19 02:00:58.185874+00	\N
23	6	1	2025-10-26	2025-10-16 03:20:48.212632+00	2025-10-16 03:20:48.212632+00	\N
24	6	2	2025-10-26	2025-10-16 03:20:48.216166+00	2025-10-16 03:20:48.216166+00	\N
25	2	2	2025-10-26	2025-10-16 03:20:48.22011+00	2025-10-16 03:20:48.22011+00	\N
158	9	1	2025-10-27	2025-10-20 03:02:37.758395+00	2025-10-20 03:02:37.758395+00	\N
159	11	1	2025-10-27	2025-10-20 03:02:37.772497+00	2025-10-20 03:02:37.772497+00	\N
160	9	1	2025-10-28	2025-10-20 03:02:37.784877+00	2025-10-20 03:02:37.784877+00	\N
29	2	2	2025-10-20	2025-10-16 13:36:57.882432+00	2025-10-16 13:36:57.882432+00	\N
161	11	1	2025-10-29	2025-10-20 03:02:37.797319+00	2025-10-20 03:02:37.797319+00	\N
162	9	1	2025-10-29	2025-10-20 03:02:37.804004+00	2025-10-20 03:02:37.804004+00	\N
163	9	1	2025-10-30	2025-10-20 03:02:37.813558+00	2025-10-20 03:02:37.813558+00	\N
164	11	1	2025-10-30	2025-10-20 03:02:37.827273+00	2025-10-20 03:02:37.827273+00	\N
165	9	1	2025-10-31	2025-10-20 03:02:37.833992+00	2025-10-20 03:02:37.833992+00	\N
166	9	1	2025-11-01	2025-10-20 03:02:37.846086+00	2025-10-20 03:02:37.846086+00	\N
36	2	4	2025-10-22	2025-10-16 13:36:57.932366+00	2025-10-16 13:36:57.932366+00	\N
167	11	1	2025-11-01	2025-10-20 03:02:37.865019+00	2025-10-20 03:02:37.865019+00	\N
168	9	1	2025-11-02	2025-10-20 03:02:37.873269+00	2025-10-20 03:02:37.873269+00	\N
169	2	1	2025-10-27	2025-10-20 03:02:37.883852+00	2025-10-20 03:02:37.883852+00	\N
170	2	2	2025-10-28	2025-10-20 03:02:37.890745+00	2025-10-20 03:02:37.890745+00	\N
171	2	1	2025-10-29	2025-10-20 03:02:37.897155+00	2025-10-20 03:02:37.897155+00	\N
172	2	2	2025-10-30	2025-10-20 03:02:37.902258+00	2025-10-20 03:02:37.902258+00	\N
173	2	1	2025-10-31	2025-10-20 03:02:37.907499+00	2025-10-20 03:02:37.907499+00	\N
174	2	2	2025-11-01	2025-10-20 03:02:37.913009+00	2025-10-20 03:02:37.913009+00	\N
175	2	1	2025-11-02	2025-10-20 03:02:37.919801+00	2025-10-20 03:02:37.919801+00	\N
176	3	1	2025-10-27	2025-10-20 03:02:37.934696+00	2025-10-20 03:02:37.934696+00	\N
177	3	2	2025-10-28	2025-10-20 03:02:37.953239+00	2025-10-20 03:02:37.953239+00	\N
178	3	1	2025-10-29	2025-10-20 03:02:37.962442+00	2025-10-20 03:02:37.962442+00	\N
179	3	2	2025-10-30	2025-10-20 03:02:37.968368+00	2025-10-20 03:02:37.968368+00	\N
180	3	1	2025-10-31	2025-10-20 03:02:37.976773+00	2025-10-20 03:02:37.976773+00	\N
181	3	2	2025-11-01	2025-10-20 03:02:37.982493+00	2025-10-20 03:02:37.982493+00	\N
182	3	1	2025-11-02	2025-10-20 03:02:37.991646+00	2025-10-20 03:02:37.991646+00	\N
183	6	1	2025-10-27	2025-10-20 03:02:37.996607+00	2025-10-20 03:02:37.996607+00	\N
184	6	2	2025-10-28	2025-10-20 03:02:38.002403+00	2025-10-20 03:02:38.002403+00	\N
185	6	1	2025-10-29	2025-10-20 03:02:38.007749+00	2025-10-20 03:02:38.007749+00	\N
186	6	2	2025-10-30	2025-10-20 03:02:38.014734+00	2025-10-20 03:02:38.014734+00	\N
187	6	1	2025-10-31	2025-10-20 03:02:38.020443+00	2025-10-20 03:02:38.020443+00	\N
188	6	2	2025-11-01	2025-10-20 03:02:38.02656+00	2025-10-20 03:02:38.02656+00	\N
189	6	1	2025-11-02	2025-10-20 03:02:38.033866+00	2025-10-20 03:02:38.033866+00	\N
197	11	1	2025-10-28	2025-10-26 02:00:59.972192+00	2025-10-26 02:00:59.972192+00	\N
198	12	1	2025-10-27	2025-10-26 02:01:00.116134+00	2025-10-26 02:01:00.116134+00	\N
199	14	1	2025-10-27	2025-10-26 02:01:00.122288+00	2025-10-26 02:01:00.122288+00	\N
200	12	1	2025-10-28	2025-10-26 02:01:00.138812+00	2025-10-26 02:01:00.138812+00	\N
201	14	1	2025-10-28	2025-10-26 02:01:00.142927+00	2025-10-26 02:01:00.142927+00	\N
202	12	1	2025-10-29	2025-10-26 02:01:00.165268+00	2025-10-26 02:01:00.165268+00	\N
203	14	1	2025-10-29	2025-10-26 02:01:00.173471+00	2025-10-26 02:01:00.173471+00	\N
74	2	1	2025-10-21	2025-10-17 02:09:31.187626+00	2025-10-17 02:09:31.187626+00	\N
78	2	1	2025-10-22	2025-10-17 02:09:31.22049+00	2025-10-17 02:09:31.22049+00	\N
79	2	2	2025-10-22	2025-10-17 02:09:31.22893+00	2025-10-17 02:09:31.22893+00	\N
83	2	2	2025-10-23	2025-10-17 02:09:31.258421+00	2025-10-17 02:09:31.258421+00	\N
87	2	2	2025-10-24	2025-10-17 02:09:31.292345+00	2025-10-17 02:09:31.292345+00	\N
90	2	1	2025-10-25	2025-10-17 02:09:31.323144+00	2025-10-17 02:09:31.323144+00	\N
190	10	1	2025-10-27	2025-10-20 03:04:35.264008+00	2025-10-20 03:04:35.264008+00	\N
191	10	1	2025-10-28	2025-10-20 03:04:35.332378+00	2025-10-20 03:04:35.332378+00	\N
192	10	1	2025-10-29	2025-10-20 03:04:35.388519+00	2025-10-20 03:04:35.388519+00	\N
193	10	1	2025-10-30	2025-10-20 03:04:35.451319+00	2025-10-20 03:04:35.451319+00	\N
194	10	1	2025-10-31	2025-10-20 03:04:35.514927+00	2025-10-20 03:04:35.514927+00	\N
195	10	1	2025-11-01	2025-10-20 03:04:35.571453+00	2025-10-20 03:04:35.571453+00	\N
196	10	1	2025-11-02	2025-10-20 03:04:35.626194+00	2025-10-20 03:04:35.626194+00	\N
118	2	1	2025-10-26	2025-10-17 02:50:23.101405+00	2025-10-17 02:50:23.101405+00	\N
150	3	1	2025-10-20	2025-10-17 03:08:43.009445+00	2025-10-17 03:08:43.009445+00	\N
151	3	1	2025-10-21	2025-10-17 03:08:43.028939+00	2025-10-17 03:08:43.028939+00	\N
152	3	1	2025-10-22	2025-10-17 03:08:43.060104+00	2025-10-17 03:08:43.060104+00	\N
153	3	1	2025-10-24	2025-10-17 03:08:43.185673+00	2025-10-17 03:08:43.185673+00	\N
\.


--
-- Data for Name: shifts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shifts (id, shift_name, code, start_time, end_time, created_at, updated_at, deleted_at) FROM stdin;
2	Ca 2	ca-2	00:00	17:00	2025-10-16 02:26:03.345859+00	2025-10-16 02:26:03.345859+00	\N
3	Ca 3	ca-3	17:00	22:00	2025-10-16 02:26:22.557291+00	2025-10-16 02:26:22.557291+00	\N
1	Ca 1	ca-1	08:00	00:00	2025-10-16 02:25:46.342076+00	2025-10-18 02:59:16.610509+00	\N
4	Ca4	ca4	08:30	00:31	2025-10-16 13:35:08.740414+00	2025-10-16 13:35:08.740414+00	2025-10-18 03:00:08.374044+00
5	Ca4	ca4	12:33	14:23	2025-10-29 05:23:41.701016+00	2025-10-29 05:23:41.701016+00	\N
\.


--
-- Data for Name: tables; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tables (id, table_name, "position", seats, memo, created_at, updated_at, deleted_at) FROM stdin;
1	Bàn 1	Trong nhà	4		2025-10-16 02:14:47.218487+00	2025-10-16 02:14:47.218487+00	\N
2	Bàn 2	Trong nhà	4		2025-10-16 02:15:11.837749+00	2025-10-16 02:15:11.837749+00	\N
3	Bàn 3	Trong nhà	4		2025-10-16 02:16:16.101679+00	2025-10-16 02:16:16.101679+00	\N
4	Bàn 4	Trong nhà	4		2025-10-16 02:16:20.206846+00	2025-10-16 02:16:20.206846+00	\N
5	Bàn 5	Trong nhà	4		2025-10-16 02:16:22.013709+00	2025-10-16 02:16:22.013709+00	\N
6	Bàn 6	Trong nhà	4		2025-10-16 02:16:23.350291+00	2025-10-16 02:16:23.350291+00	\N
7	Bàn 7	Trong nhà	4		2025-10-16 02:16:25.063847+00	2025-10-16 02:16:25.063847+00	\N
8	Bàn 8	Trong nhà	4		2025-10-16 02:16:26.905229+00	2025-10-16 02:16:26.905229+00	\N
9	Bàn 9	Trong nhà	4		2025-10-16 02:16:29.046529+00	2025-10-16 02:16:29.046529+00	\N
10	Bàn 10	Trong nhà	4		2025-10-16 02:16:30.696245+00	2025-10-16 02:16:30.696245+00	\N
11	Bàn 11	Trong nhà	4		2025-10-16 02:16:32.424015+00	2025-10-16 02:16:32.424015+00	\N
12	Bàn 12	Trong nhà	4		2025-10-16 02:16:33.551932+00	2025-10-16 02:16:33.551932+00	\N
13	Bàn 13	Trong nhà	4		2025-10-16 02:16:34.799088+00	2025-10-16 02:16:34.799088+00	\N
14	Bàn 14	Trong nhà	4		2025-10-16 02:16:35.759689+00	2025-10-16 02:16:35.759689+00	\N
15	Bàn 15	Trong nhà	4		2025-10-16 02:16:36.752592+00	2025-10-16 02:16:36.752592+00	\N
16	Bàn 16	Trong nhà	4		2025-10-16 02:16:37.64844+00	2025-10-16 02:16:37.64844+00	\N
17	Bàn 17	Trong nhà	4		2025-10-16 02:16:38.752458+00	2025-10-16 02:16:38.752458+00	\N
18	Bàn 18	Trong nhà	4		2025-10-16 02:16:40.121098+00	2025-10-16 02:16:40.121098+00	\N
19	Bàn 19	Trong nhà	4		2025-10-16 02:16:42.199998+00	2025-10-16 02:16:42.199998+00	\N
20	Bàn 20	Trong nhà	4		2025-10-16 02:16:44.09706+00	2025-10-16 02:18:36.790047+00	\N
\.


--
-- Data for Name: tickets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tickets (id, ingredient_id, quantity, unit, ticket_type, created_at, updated_at, deleted_at) FROM stdin;
1	5	10	kg	Import	2025-10-16 02:56:39.058041+00	2025-10-16 02:56:39.058041+00	\N
2	1	2	qua	Export	2025-10-16 02:56:54.899108+00	2025-10-16 02:56:54.899108+00	\N
3	25	3	thung	Export	2025-10-18 03:15:38.318776+00	2025-10-18 03:15:38.318776+00	\N
\.


--
-- Name: accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.accounts_id_seq', 17, true);


--
-- Name: attendances_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.attendances_id_seq', 2, true);


--
-- Name: availibilities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.availibilities_id_seq', 455, true);


--
-- Name: bookings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bookings_id_seq', 1, true);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 1, false);


--
-- Name: customers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.customers_id_seq', 1, true);


--
-- Name: employees_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.employees_id_seq', 15, true);


--
-- Name: files_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.files_id_seq', 9, true);


--
-- Name: ingredients_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ingredients_id_seq', 25, true);


--
-- Name: menu_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.menu_items_id_seq', 7, true);


--
-- Name: order_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_items_id_seq', 16, true);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_id_seq', 5, true);


--
-- Name: permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.permissions_id_seq', 54, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.roles_id_seq', 8, true);


--
-- Name: shift_schedules_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.shift_schedules_id_seq', 203, true);


--
-- Name: shifts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.shifts_id_seq', 5, true);


--
-- Name: tables_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tables_id_seq', 20, true);


--
-- Name: tickets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tickets_id_seq', 3, true);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: attendances attendances_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.attendances
    ADD CONSTRAINT attendances_pkey PRIMARY KEY (id);


--
-- Name: availibilities availibilities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.availibilities
    ADD CONSTRAINT availibilities_pkey PRIMARY KEY (id);


--
-- Name: bookings bookings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_pkey PRIMARY KEY (id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


--
-- Name: employees employees_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_pkey PRIMARY KEY (id);


--
-- Name: files files_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_pkey PRIMARY KEY (id);


--
-- Name: ingredients ingredients_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ingredients
    ADD CONSTRAINT ingredients_pkey PRIMARY KEY (id);


--
-- Name: menu_items menu_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.menu_items
    ADD CONSTRAINT menu_items_pkey PRIMARY KEY (id);


--
-- Name: order_items order_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: permissions permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);


--
-- Name: role_permissions role_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT role_permissions_pkey PRIMARY KEY (permission_id, role_id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: shift_schedules shift_schedules_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shift_schedules
    ADD CONSTRAINT shift_schedules_pkey PRIMARY KEY (id);


--
-- Name: shifts shifts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shifts
    ADD CONSTRAINT shifts_pkey PRIMARY KEY (id);


--
-- Name: tables tables_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT tables_pkey PRIMARY KEY (id);


--
-- Name: tickets tickets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tickets
    ADD CONSTRAINT tickets_pkey PRIMARY KEY (id);


--
-- Name: accounts uni_accounts_user_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT uni_accounts_user_name UNIQUE (user_name);


--
-- Name: permissions uni_permissions_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT uni_permissions_name UNIQUE (name);


--
-- Name: roles uni_roles_role_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT uni_roles_role_name UNIQUE (role_name);


--
-- Name: idx_accounts_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_accounts_deleted_at ON public.accounts USING btree (deleted_at);


--
-- Name: idx_attendances_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_attendances_deleted_at ON public.attendances USING btree (deleted_at);


--
-- Name: idx_availibilities_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_availibilities_deleted_at ON public.availibilities USING btree (deleted_at);


--
-- Name: idx_bookings_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bookings_deleted_at ON public.bookings USING btree (deleted_at);


--
-- Name: idx_categories_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_categories_deleted_at ON public.categories USING btree (deleted_at);


--
-- Name: idx_customers_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_customers_deleted_at ON public.customers USING btree (deleted_at);


--
-- Name: idx_customers_phone_number; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_customers_phone_number ON public.customers USING btree (phone_number);


--
-- Name: idx_employees_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_employees_deleted_at ON public.employees USING btree (deleted_at);


--
-- Name: idx_files_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_files_deleted_at ON public.files USING btree (deleted_at);


--
-- Name: idx_ingredients_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ingredients_deleted_at ON public.ingredients USING btree (deleted_at);


--
-- Name: idx_menu_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_menu_items_deleted_at ON public.menu_items USING btree (deleted_at);


--
-- Name: idx_order_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_order_items_deleted_at ON public.order_items USING btree (deleted_at);


--
-- Name: idx_orders_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_orders_deleted_at ON public.orders USING btree (deleted_at);


--
-- Name: idx_permissions_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_permissions_deleted_at ON public.permissions USING btree (deleted_at);


--
-- Name: idx_roles_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_roles_deleted_at ON public.roles USING btree (deleted_at);


--
-- Name: idx_shift_schedules_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_shift_schedules_deleted_at ON public.shift_schedules USING btree (deleted_at);


--
-- Name: idx_shifts_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_shifts_deleted_at ON public.shifts USING btree (deleted_at);


--
-- Name: idx_tables_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tables_deleted_at ON public.tables USING btree (deleted_at);


--
-- Name: idx_tickets_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tickets_deleted_at ON public.tickets USING btree (deleted_at);


--
-- Name: accounts fk_accounts_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT fk_accounts_role FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- Name: attendances fk_attendances_shift_schedule; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.attendances
    ADD CONSTRAINT fk_attendances_shift_schedule FOREIGN KEY (shift_schedule_id) REFERENCES public.shift_schedules(id);


--
-- Name: availibilities fk_availibilities_employee; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.availibilities
    ADD CONSTRAINT fk_availibilities_employee FOREIGN KEY (employee_id) REFERENCES public.employees(id);


--
-- Name: availibilities fk_availibilities_shift; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.availibilities
    ADD CONSTRAINT fk_availibilities_shift FOREIGN KEY (shift_id) REFERENCES public.shifts(id);


--
-- Name: employees fk_employees_account; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT fk_employees_account FOREIGN KEY (account_id) REFERENCES public.accounts(id);


--
-- Name: employees fk_employees_avatar_file; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT fk_employees_avatar_file FOREIGN KEY (avatar_file_id) REFERENCES public.files(id);


--
-- Name: menu_items fk_menu_items_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.menu_items
    ADD CONSTRAINT fk_menu_items_category FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- Name: menu_items fk_menu_items_file; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.menu_items
    ADD CONSTRAINT fk_menu_items_file FOREIGN KEY (file_id) REFERENCES public.files(id);


--
-- Name: order_items fk_order_items_menu_item; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT fk_order_items_menu_item FOREIGN KEY (menu_item_id) REFERENCES public.menu_items(id);


--
-- Name: order_items fk_order_items_order; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT fk_order_items_order FOREIGN KEY (order_id) REFERENCES public.orders(id);


--
-- Name: orders fk_orders_customer; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_orders_customer FOREIGN KEY (customer_id) REFERENCES public.customers(id) ON UPDATE RESTRICT ON DELETE SET NULL;


--
-- Name: orders fk_orders_table; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_orders_table FOREIGN KEY (table_id) REFERENCES public.tables(id);


--
-- Name: role_permissions fk_role_permissions_permission; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES public.permissions(id);


--
-- Name: role_permissions fk_role_permissions_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- Name: shift_schedules fk_shift_schedules_employee; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shift_schedules
    ADD CONSTRAINT fk_shift_schedules_employee FOREIGN KEY (employee_id) REFERENCES public.employees(id);


--
-- Name: shift_schedules fk_shift_schedules_shift; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shift_schedules
    ADD CONSTRAINT fk_shift_schedules_shift FOREIGN KEY (shift_id) REFERENCES public.shifts(id);


--
-- Name: tickets fk_tickets_ingredient; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tickets
    ADD CONSTRAINT fk_tickets_ingredient FOREIGN KEY (ingredient_id) REFERENCES public.ingredients(id);


--
-- PostgreSQL database dump complete
--

\unrestrict h1Jhr5F3wvcmuM7HDNp1es9bzO6AbKplCX7G8rHCKAeqRdcyXChLZDq3MZIesnr

