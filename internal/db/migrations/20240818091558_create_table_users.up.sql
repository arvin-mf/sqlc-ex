CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,            
    email VARCHAR(255) UNIQUE NOT NULL,   
    password VARCHAR(255),       
    oauth_id VARCHAR(255),                
    name VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL
) ENGINE = InnoDB;