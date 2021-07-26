package com.diegorosa.kafka.manager.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;

import java.io.Serializable;
import java.util.UUID;

@Data
@JsonIgnoreProperties(ignoreUnknown = true)
public class AccountManagerDTO implements Serializable {

    private UUID id;

    private String name;
}
