package com.diegorosa.kafka.manager.service;

import com.diegorosa.kafka.manager.model.AccountManager;
import com.diegorosa.kafka.manager.model.AccountManagerDTO;
import com.diegorosa.kafka.manager.model.CustomerDTO;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@RequiredArgsConstructor
@Service
public class AccountManagerService {

    private final AccountManagerRepository repository;

    private final ObjectMapper jsonMarshal;

    private final CustomerAllocationRule allocationRule;

    public AccountManagerDTO save(AccountManagerDTO accountManagerDTO) {
        AccountManager manager = new AccountManager();
        manager.setName(accountManagerDTO.getName());
        manager = repository.save(manager);
        accountManagerDTO.setId(manager.getId());
        return accountManagerDTO;
    }

    @KafkaListener(id = "manager-service", topics = {"new_account"})
    public void allocate(String record) throws JsonProcessingException {
        var customer = jsonMarshal.readValue(record, CustomerDTO.class);
        if (allocationRule.isValid(customer)) {
            System.out.println("allocating a manager for "+customer.getName() + " with salary "+customer.getSalary());
        } else {
            System.out.println("Customer "+customer.getName()+" validated.");
        }
    }
}
