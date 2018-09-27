CREATE TABLE player_summary(
   id             INTEGER  NOT NULL PRIMARY KEY 
  ,name           VARCHAR(9)
  ,normalizedName VARCHAR(9)
  ,teamShortName  VARCHAR(3)
  ,teamLongName   VARCHAR(7)
  ,position       VARCHAR(2)
  ,points         INTEGER 
  ,goals          INTEGER
  ,assists        INTEGER
  ,minutes        INTEGER 
  ,reds           INTEGER
  ,yellows        INTEGER 
  ,saves          INTEGER 
  ,pensSaved      INTEGER 
  ,pensMissed     INTEGER
  ,bonus          INTEGER 
  ,cleanSheets    INTEGER 
  ,lastCost       NUMERIC(3,1)
);

CREATE TABLE player_detail(
   opponent        VARCHAR(3) NOT NULL PRIMARY KEY
  ,elementTeam     VARCHAR(30)
  ,gameWeek        INT 
  ,points          INT 
  ,goals           INT 
  ,assists         INT 
  ,minutes         INT 
  ,reds            INT 
  ,yellows         INT 
  ,pensMissed      INT 
  ,bonus           INT 
  ,bps             INT 
  ,cleanSheets     INT 
  ,cost            INT 
  ,goalsConceded   INT 
  ,ownGoals        INT 
  ,penaltiesSaved  INT 
  ,penaltiesMissed INT 
  ,saves           INT 
  ,influence       INT 
  ,creativity      INT 
  ,threat          INT 
  ,ictIndex        INT 
  ,selectedBy      INT 
  ,netTransfers    INT 
);