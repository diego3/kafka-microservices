package com.diegorosa.kafka.customer;

import org.apache.kafka.clients.admin.NewTopic;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.kafka.config.TopicBuilder;

@SpringBootApplication
public class ApplicationCustomerService {

	public static void main(String[] args) {
		SpringApplication.run(ApplicationCustomerService.class, args);
	}

	@Bean
	public NewTopic topic() {
		return TopicBuilder.name("new_account")
				.partitions(2)
				.replicas(1)
				.build();
	}

}
