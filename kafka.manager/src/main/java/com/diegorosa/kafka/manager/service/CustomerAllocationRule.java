package com.diegorosa.kafka.manager.service;

import com.diegorosa.kafka.manager.model.CustomerDTO;
import org.springframework.stereotype.Service;

@Service
public class CustomerAllocationRule {
    private static final double MANAGER_PLATINIUM = 1000.0;
    private static final double MANAGER_GOLD = 15000.0;

    public boolean isValid(CustomerDTO customerDTO) {
        if (customerDTO.getSalary() >= MANAGER_GOLD) {
            return true;
        }

        if (customerDTO.getSalary() >= MANAGER_PLATINIUM) {
            return true;
        }

        return false;
    }
}
