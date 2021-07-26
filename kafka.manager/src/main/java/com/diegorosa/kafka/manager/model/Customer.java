package com.diegorosa.kafka.manager.model;


import lombok.Data;

import javax.persistence.*;
import java.util.UUID;

@Data
@Entity
public class Customer {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(columnDefinition = "BINARY(16)")
    private UUID id;

    private String name;

    private String email;

    private Double renda;

    @ManyToOne
    @JoinColumn(name="manager_id")
    private AccountManager manager;
}
