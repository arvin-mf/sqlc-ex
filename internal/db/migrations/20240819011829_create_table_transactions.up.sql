CREATE TABLE transactions (
    id CHAR(36) PRIMARY KEY,                   
    user_id CHAR(36) NOT NULL,                
    descript VARCHAR(255)
) ENGINE = InnoDB;