CREATE TABLE users (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    email_verification_key VARCHAR(255),
    email_verification_time TIMESTAMP,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    cell_phone VARCHAR(20) NOT NULL,
    city VARCHAR(255),
    state VARCHAR(255),
    zip_code VARCHAR(20),
    district VARCHAR(255),
    street VARCHAR(255),
    street_number VARCHAR(20),
    complement VARCHAR(255),
    birthdate DATE,
    gender ENUM('Male', 'Female', 'Other'),
    cpf VARCHAR(14),
    source TINYINT,
    record_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    level INTEGER
);

CREATE TABLE user_sub_accounts (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    key_api VARCHAR(255) UNIQUE
);

CREATE TABLE related_sub_accounts_in_master_user (
    sub_account_id INT,
    user_id INT,
    PRIMARY KEY (sub_account_id, user_id),
    FOREIGN KEY (sub_account_id) REFERENCES user_sub_accounts(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
