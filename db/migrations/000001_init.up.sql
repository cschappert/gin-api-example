 CREATE TABLE teams(
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(255),
    created_at datetime,
    updated_at datetime,
    PRIMARY KEY (id));

CREATE TABLE accounts(
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(255),
    email varchar(255),
    created_at datetime,
    updated_at datetime,
    team_id int(11),
    PRIMARY KEY (id),
    CONSTRAINT fk_team_account FOREIGN KEY (team_id) REFERENCES teams(id));

