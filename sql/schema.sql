create table teams(
	id int auto_increment PRIMARY KEY,
	name varchar(255) not null,
	shortname varchar(255) not null,
    country varchar(2) not null default 'US',
    conference varchar(4),
    division varchar(20)
);

create table transfers(
	id int auto_increment PRIMARY KEY,
    date date not null,
    seller int not null,
    buyer int not null,
    player varchar(255) not null,
    status varchar(50) not null,
    FOREIGN KEY (seller) REFERENCES teams(id),
    FOREIGN KEY (buyer) REFERENCES teams(id)
);

-- deal, rumour and rejected (status)

-- Atlantic East
insert into teams (name, shortname, conference, division) values ('Boston Celtics', 'BOS', 'East', 'Atlantic');
insert into teams (name, shortname, conference, division) values ('Brooklyn Nets', 'BKN', 'East', 'Atlantic');
insert into teams (name, shortname, conference, division) values ('New York Knicks', 'NYK', 'East', 'Atlantic');
insert into teams (name, shortname, conference, division) values ('Philadelphia 76ers', 'PHI', 'East', 'Atlantic');
insert into teams (name, shortname, conference, division) values ('Toronto Raptors', 'TOR', 'East', 'Atlantic');


-- Central East
insert into teams (name, shortname, conference, division) values ('Chicago Bulls', 'CHI', 'East', 'Central');
insert into teams (name, shortname, conference, division) values ('Cleveland Cavaliers', 'CLE', 'East', 'Central');
insert into teams (name, shortname, conference, division) values ('Detroit Pistons', 'DET', 'East', 'Central');
insert into teams (name, shortname, conference, division) values ('Indiana Pacers', 'IND', 'East', 'Central');
insert into teams (name, shortname, conference, division) values ('Milwaukee Bucks', 'MIL', 'East', 'Central');

-- Southeast East
insert into teams (name, shortname, conference, division) values ('Atlanta Hawks', 'ATL', 'East', 'Southeast');
insert into teams (name, shortname, conference, division) values ('Charlote Hornets', 'CHA', 'East', 'Southeast');
insert into teams (name, shortname, conference, division) values ('Miami Heat', 'MIA', 'East', 'Southeast');
insert into teams (name, shortname, conference, division) values ('Orlando Magic', 'ORL', 'East', 'Southeast');
insert into teams (name, shortname, conference, division) values ('Washington Wizards', 'WAS', 'East', 'Southeast');

-- Northwest West
insert into teams (name, shortname, conference, division) values ('Denver Nuggets', 'DEN', 'West', 'Northwest');
insert into teams (name, shortname, conference, division) values ('Minnesota Timberwolves', 'MIN', 'West', 'Northwest');
insert into teams (name, shortname, conference, division) values ('Oklahoma City Thunders', 'OKC', 'West', 'Northwest');
insert into teams (name, shortname, conference, division) values ('Portland Trailblazers', 'POR', 'West', 'Northwest');
insert into teams (name, shortname, conference, division) values ('Utah Jazz', 'UTA', 'West', 'Northwest');

-- Pacific West
insert into teams (name, shortname, conference, division) values ('Golden State Warriors', 'GSW', 'West', 'Pacific');
insert into teams (name, shortname, conference, division) values ('Los Angeles Clippers', 'LAC', 'West', 'Pacific');
insert into teams (name, shortname, conference, division) values ('Los Angeles Lakers', 'LAL', 'West', 'Pacific');
insert into teams (name, shortname, conference, division) values ('Phoenix Suns', 'PHX', 'West', 'Pacific');
insert into teams (name, shortname, conference, division) values ('Sacramento Kings', 'SAC', 'West', 'Pacific');

-- Southwest West
insert into teams (name, shortname, conference, division) values ('Dallas Mavericks', 'DAL', 'West', 'Southwest');
insert into teams (name, shortname, conference, division) values ('Houston Rockets', 'HOU', 'West', 'Southwest');
insert into teams (name, shortname, conference, division) values ('Memphis Grizzlies', 'MEM', 'West', 'Southwest');
insert into teams (name, shortname, conference, division) values ('New Orleans Pelicans', 'NOP', 'West', 'Southwest');
insert into teams (name, shortname, conference, division) values ('San Antonio Spurs', 'SAS', 'West', 'Southwest');