package com.diegorosa.kafka.customer.service;

import com.diegorosa.kafka.customer.model.Customer;
import com.diegorosa.kafka.customer.model.CustomerDTO;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.stereotype.Service;

@RequiredArgsConstructor
@Service
public class CustomerService {
    private final CustomerRepository repository;

    public CustomerDTO save(CustomerDTO customerDTO) {
        Customer entity = new Customer();
        entity.setEmail(customerDTO.getEmail());
        entity.setName(customerDTO.getName());
        entity.setRenda(customerDTO.getRenda());
        entity = repository.save(entity);
        customerDTO.setId(entity.getId());
        return customerDTO;
    }

    public Page<Customer> getPaginated() {
        return repository.findAll(PageRequest.of(0, 10));
    }
}
