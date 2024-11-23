package com.mg.perftests;

import org.springframework.stereotype.Repository;
import org.springframework.data.mongodb.repository.MongoRepository;

@Repository
public interface SongRepository extends MongoRepository<Song, String> {
	
}
