package com.diegorosa.kafka.manager;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.kafka.annotation.KafkaListener;

@SpringBootApplication
public class Application {

	public static void main(String[] args) {
		SpringApplication.run(Application.class, args);
	}

	@KafkaListener(id = "myId", topics = {"new_account"})
	public void listen(String record) {
		// TODO deserialize to customerDTO
		System.out.println("record: "+record);
	}
}
