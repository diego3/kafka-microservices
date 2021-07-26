package com.diegorosa.kafka.manager.service;

import com.diegorosa.kafka.manager.model.AccountManager;
import com.diegorosa.kafka.manager.model.AccountManagerDTO;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@RequiredArgsConstructor
@Service
public class AccountManagerService {

    private final AccountManagerRepository repository;

    public AccountManagerDTO save(AccountManagerDTO accountManagerDTO) {
        AccountManager manager = new AccountManager();
        manager.setName(accountManagerDTO.getName());
        manager = repository.save(manager);
        accountManagerDTO.setId(manager.getId());
        return accountManagerDTO;
    }
}
