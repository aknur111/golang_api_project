-- CREATE DATABASE ourproject;

-- Creating tables
CREATE TABLE groups (
    group_id INTEGER PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    num_of_members INTEGER NOT NULL,
    launch_date DATE DEFAULT CURRENT_DATE
);

CREATE TABLE singer (
    singer_id INTEGER PRIMARY KEY,
    first_name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    birthday DATE NOT NULL,
    group_id INTEGER REFERENCES groups(group_id)
);

CREATE TABLE album (
    album_id INTEGER PRIMARY KEY,
    title VARCHAR(25) NOT NULL,
    genre VARCHAR(25) NOT NULL,
    num_of_tracks INTEGER NOT NULL,
    group_id INTEGER REFERENCES groups(group_id) NOT NULL
);

CREATE TABLE song (
    song_id INTEGER PRIMARY KEY,
    title VARCHAR(25) NOT NULL,
    length INTEGER,
    album_id INTEGER REFERENCES album(album_id) NOT NULL
);


-- Inserting initial values for created tables

INSERT INTO groups(group_id, name, num_of_members)
VALUES (1, 'BTS', 7),
       (2, 'BlackPink', 4),
       (3, 'EXO', 9);

INSERT INTO singer(singer_id, first_name, last_name, birthday, group_id)
VALUES (1, 'Nam-joon', 'Kim','1994-09-12',1),
       (2,'Ji-min', 'Park','1995-10-13',1),
       (3,'Yoon-gi', 'Min','1993-03-09',1),
       (4,'Ji-soo', 'Kim','1995-01-03',2),
       (5,'Jennie', 'Kim','1996-01-16',2),
       (6,'Lisa', 'Manoban','1997-03-27',2),
       (7,'Roseanne', 'Park','1997-02-11',2),
       (8,'Baek-hyun', 'Byun','1992-05-06',3),
       (9,'Chan-yeol', 'Park','1992-11-27',3),
       (10,'Min-seok', 'Kim','1990-03-26',3);


INSERT INTO album(album_id, title, genre, num_of_tracks, group_id)
VALUES (100,'Map of the Soul', 'EDM',16,1),
       (101,'Born Pink','Pop',8,2),
       (102,'Dont mess up my Tempo', 'Dance',11,3);

INSERT INTO song(song_id, title, length, album_id)
VALUES (111,'ON',250,100),
       (112,'Pink Venom',Null,101),
       (113,'Love Shot', NULL,102);