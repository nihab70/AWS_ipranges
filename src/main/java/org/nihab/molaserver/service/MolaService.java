package org.nihab.molaserver.service;

import org.springframework.stereotype.Service;

import org.nihab.molaserver.model.MolaRepository;
import org.nihab.molaserver.model.MolaModel;

import java.util.List;

@Service
public class MolaService {

    private final MolaRepository molaRepository;

    public MolaService(MolaRepository molaRepository) {this.molaRepository = molaRepository;}

    public MolaModel saveMola(MolaModel molaModel) {
        return molaRepository.save(molaModel);
    }

    public List<MolaModel> findAll() {
        return molaRepository.findAll();
    }
}
