package com.diegorosa.kafka.manager.controller;

import com.diegorosa.kafka.manager.model.AccountManagerDTO;
import com.diegorosa.kafka.manager.service.AccountManagerService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RequiredArgsConstructor
@RestController
public class ManagerController {
    private final AccountManagerService accountManagerService;

    @PostMapping("/manager")
    public AccountManagerDTO registerNewManager(@RequestBody AccountManagerDTO accountManagerDTO) {
        return accountManagerService.save(accountManagerDTO);
    }

}
