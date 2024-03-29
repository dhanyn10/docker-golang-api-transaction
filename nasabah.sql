-- Adminer 4.8.1 PostgreSQL 14.2 (Debian 14.2-1.pgdg110+1) dump

\connect "admin";

DROP TABLE IF EXISTS "history";
DROP SEQUENCE IF EXISTS history_id_seq;
CREATE SEQUENCE history_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."history" (
    "id" integer DEFAULT nextval('history_id_seq') NOT NULL,
    "event" character varying(50) NOT NULL,
    "cond" smallint NOT NULL,
    "activity" text NOT NULL,
    "datetime" text NOT NULL,
    CONSTRAINT "history_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "nasabah";
DROP SEQUENCE IF EXISTS nasabah_id_seq;
CREATE SEQUENCE nasabah_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."nasabah" (
    "id" integer DEFAULT nextval('nasabah_id_seq') NOT NULL,
    "username" character varying(250) NOT NULL,
    "password" character varying(250) NOT NULL,
    "token" text NOT NULL,
    "tabungan" integer NOT NULL,
    CONSTRAINT "nasabah_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "nasabah_username_key" UNIQUE ("username")
) WITH (oids = false);

INSERT INTO "nasabah" ("id", "username", "password", "token", "tabungan") VALUES
(1,	'hero',	'$2a$14$Ecpo5cbpO0TVTNHRxiHrS.OjzZjOxLKTVUOEV/dBfx8zfsbuhCivK',	'token',	150000000),
(2,	'zeus',	'$2a$14$Eqqe7.mHNWAZWGk6FcauSem6tn6odv6UYCORBDNU57I/6H74q9gAu',	'token',	35000000);

DROP TABLE IF EXISTS "transaksi";
DROP SEQUENCE IF EXISTS transaksi_id_seq;
CREATE SEQUENCE transaksi_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."transaksi" (
    "id" integer DEFAULT nextval('transaksi_id_seq') NOT NULL,
    "from" character varying(250) NOT NULL,
    "to" character varying(250) NOT NULL,
    "amount" money NOT NULL,
    "datetime" text NOT NULL,
    CONSTRAINT "transaksi_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


-- 2022-04-10 05:51:00.049491+00
