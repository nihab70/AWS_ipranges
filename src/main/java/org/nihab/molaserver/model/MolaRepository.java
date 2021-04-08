package org.nihab.molaserver.model;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Component;

@Component
public interface MolaRepository extends MongoRepository<MolaModel ,String> {
}
