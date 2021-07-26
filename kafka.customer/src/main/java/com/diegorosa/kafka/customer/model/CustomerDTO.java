package com.diegorosa.kafka.customer.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;

import java.io.Serializable;
import java.util.UUID;

@Data
@JsonIgnoreProperties(ignoreUnknown = true)
public class CustomerDTO implements Serializable {

    private UUID id;

    private String name;

    private String email;

    private Double renda;
}
