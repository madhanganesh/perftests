package com.mg.perftests;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "songs")
public class Song {
    @Id
    private String id;
    private String title;
    private String[] lyrics;

    public String getId() {
        return id;
    }

    public String getTitle() {
        return title;
    }

    /**
     * @return the lyrics
     */
    public String[] getLyrics() {
        return lyrics;
    }

    /**
     * @param lyrics the lyrics to set
     */
    public void setLyrics(String[] lyrics) {
        this.lyrics = lyrics;
    }

    public void setTitle(String title) {
        this.title = title;
    }
}
