package org.nihab.molaserver.model;

import lombok.Builder;
import lombok.Data;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document
@Data
@Builder
public class MolaModel {

    @Id
    private String id;
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
}