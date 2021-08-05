package com.diegorosa.kafka.customer.controller;

import com.diegorosa.kafka.customer.model.Customer;
import com.diegorosa.kafka.customer.model.CustomerDTO;
import com.diegorosa.kafka.customer.service.CustomerService;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RequiredArgsConstructor
@RestController
public class CustomerController {

    private final CustomerService customerService;

    @PostMapping("/customer")
    public CustomerDTO openNewAccount(@RequestBody CustomerDTO customerDTO) {
        return customerService.save(customerDTO);
    }

    @GetMapping("/customer")
    public Page<Customer> page() {
        return customerService.getPaginated();
    }
}
