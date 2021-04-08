package org.nihab.molaserver.controller;


import org.nihab.molaserver.model.MolaModel;

import lombok.Data;

@Data
public class MolaRequest {
    private String name;
    private String firstName;
    private String address;
    private String birthday;
    private String placeOfBirth;
    private String nationality;
    private String numberOfId;
    private String dateOfId;
    private String validThruOfId;
    private String typeOfIdentification;

    public MolaModel toMolaModel(){
       return MolaModel.builder()
                                .name(this.name)
                                .firstName(this.firstName)
                                .address(this.address)
                                .birthday(this.birthday)
                                .placeOfBirth(this.placeOfBirth)
                                .nationality(this.nationality)
                                .numberOfId(this.numberOfId)
                                .dateOfId(this.dateOfId)
                                .validThruOfId(this.validThruOfId)
                                .typeOfIdentification(this.typeOfIdentification)
                                .build();
    }
}
