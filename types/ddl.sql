CREATE TABLE public.player_summary (
	id int4 NOT NULL,
	"name" varchar(100) NULL,
	normalizedname varchar(100) NULL,
	teamshortname varchar(3) NULL,
	teamlongname varchar(50) NULL,
	"position" varchar(5) NULL,
	points int4 NULL,
	goals int4 NULL,
	assists int4 NULL,
	minutes int4 NULL,
	reds int4 NULL,
	yellows int4 NULL,
	saves int4 NULL,
	penssaved int4 NULL,
	pensmissed int4 NULL,
	bonus int4 NULL,
	cleansheets int4 NULL,
	lastcost numeric(3,1) NULL,
	season varchar(10) NULL,
	CONSTRAINT player_summary_pkey PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
) ;

CREATE TABLE public.player_detail (
	opponent varchar(3) NULL,
	elementteam varchar(30) NULL,
	gameweek int4 NULL,
	points int4 NULL,
	goals int4 NULL,
	assists int4 NULL,
	minutes int4 NULL,
	reds int4 NULL,
	yellows int4 NULL,
	pensmissed int4 NULL,
	bonus int4 NULL,
	bps int4 NULL,
	cleansheets int4 NULL,
	cost numeric(3,1) NULL,
	goalsconceded int4 NULL,
	owngoals int4 NULL,
	penaltiessaved int4 NULL,
	penaltiesmissed int4 NULL,
	saves int4 NULL,
	influence numeric(7,3) NULL,
	creativity numeric(7,3) NULL,
	threat int4 NULL,
	ictindex numeric(7,3) NULL,
	selectedby int4 NULL,
	nettransfers int4 NULL,
	season varchar(10) NULL,
	id int4 NOT NULL DEFAULT nextval('player_detail_id_seq'::regclass),
	CONSTRAINT player_detail_pk PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
) ;
