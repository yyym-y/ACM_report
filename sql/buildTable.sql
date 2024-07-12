CREATE DATABASE acm_report;
use acm_report;

# problem_list(problem_id, problem_name, problem_url, type_id, solve_time, diff_id);
CREATE TABLE problem_list(
    problem_id INT PRIMARY KEY UNIQUE AUTO_INCREMENT,
    problem_name VARCHAR(50) NOT NULL,
    problem_url VARCHAR(150) NOT NULL,
    type_id INT NOT NULL,
    solve_time DATE,
    diff_id INT
);

# solution_list(problem_id, solution);
CREATE TABLE solution_list(
    problem_id int NOT NULL,
    solution TEXT NOT NULL
);

CREATE TABLE descripe_list(
    problem_id int NOT NULL,
    descripe_type VARCHAR(30) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE descripe_type(
    descripe_type_id int NOT NULL PRIMARY KEY UNIQUE AUTO_INCREMENT,
    descripe_name VARCHAR(20)
);

# type_list(type_id, type_name);
CREATE TABLE type_list(
    type_id INT PRIMARY KEY UNIQUE AUTO_INCREMENT,
    type_name VARCHAR(30) UNIQUE NOT NULL
);


# difficulty(diff_id, diff_name);
CREATE TABLE difficulty(
    diff_id INT AUTO_INCREMENT PRIMARY KEY,
    diff_name VARCHAR(30) NOT NULL UNIQUE
);

SHOW TABLES;

DROP TABLE descripe_type;
DROP TABLE descripe_list;