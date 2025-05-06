CREATE TABLE favorites (
    user_id INTEGER NOT NULL,
    song_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, song_id),
    FOREIGN KEY (song_id) REFERENCES song(song_id)
);
