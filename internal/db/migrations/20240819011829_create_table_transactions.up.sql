CREATE TABLE transactions (
    id CHAR(36) PRIMARY KEY,                   
    user_id CHAR(36) NOT NULL,                
    descript VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE = InnoDB;