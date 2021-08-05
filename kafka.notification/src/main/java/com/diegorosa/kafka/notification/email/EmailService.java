package com.diegorosa.kafka.notification.email;

import com.diegorosa.kafka.notification.model.CustomerDTO;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@RequiredArgsConstructor
@Service
public class EmailService {
    private final ObjectMapper jsonMarshal;

    @KafkaListener(id = "email-service", topics = {"new_account"})
    public void send(String record) throws JsonProcessingException {
        var customer = jsonMarshal.readValue(record, CustomerDTO.class);
        System.out.println("sending welcome email to :: "+customer.getEmail());
    }
}
