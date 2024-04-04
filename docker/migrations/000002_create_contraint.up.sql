SET search_path TO vending_machine_service;

ALTER TABLE USER_LOGIN ADD CONSTRAINT unique_username UNIQUE(username);
