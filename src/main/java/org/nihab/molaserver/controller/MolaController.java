package org.nihab.molaserver.controller;

import java.util.List;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import org.nihab.molaserver.service.MolaService;
import org.nihab.molaserver.model.MolaModel;

import lombok.extern.slf4j.Slf4j;

@RestController
@Slf4j
public class MolaController {

    private final MolaService molaService;

    public MolaController(MolaService molaService) {this.molaService = molaService;}

    @PostMapping("/molas")
    public ResponseEntity<String> addMola(@RequestBody MolaRequest molaRequest) {
        log.info("Request : {}", molaRequest);
        molaService.saveMola(molaRequest.toMolaModel());
        return ResponseEntity.status(HttpStatus.CREATED).build();
    }

    @GetMapping("/molas")
    public List<MolaModel> getAllMola() {
        return molaService.findAll();
    }
}
