package com.diegorosa.kafka.manager.service;

import com.diegorosa.kafka.manager.model.AccountManager;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public interface AccountManagerRepository extends JpaRepository<AccountManager, UUID> {

}
