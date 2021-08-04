package com.diegorosa.kafka.notification.email;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@Service
public class EmailService {

    public void send() {

    }

    @KafkaListener(id = "myId", topics = {"new_account"})
    public void listen(String record) {
        System.out.println("sending email to "+record);
    }
}
