package com.diegorosa.kafka.customer.service;

import com.diegorosa.kafka.customer.model.Customer;
import com.diegorosa.kafka.customer.model.CustomerDTO;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@RequiredArgsConstructor
@Service
@Log4j2
public class CustomerService {
    private final CustomerRepository repository;

    private final KafkaTemplate<String, String> kafkaTemplate;

    private final ObjectMapper jsonMarshal;

    public CustomerDTO save(CustomerDTO customerDTO) {
        Customer entity = new Customer();
        entity.setEmail(customerDTO.getEmail());
        entity.setName(customerDTO.getName());
        entity.setSalary(customerDTO.getSalary());
        entity = repository.save(entity);
        customerDTO.setId(entity.getId());

        try {
            kafkaTemplate.send("new_account", jsonMarshal.writeValueAsString(customerDTO));
        } catch (JsonProcessingException e) {
            log.warn("new_account send error", e);
        }

        return customerDTO;
    }

    public Page<Customer> getPaginated() {
        return repository.findAll(PageRequest.of(0, 10));
    }
}
