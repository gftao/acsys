DROP TABLE pcpmp.tbl_role;
CREATE TABLE pcpmp.tbl_role (id int NOT NULL AUTO_INCREMENT, name varchar(255) NOT NULL, seq int DEFAULT 0 NOT NULL, PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8 DEFAULT COLLATE=utf8_general_ci;
INSERT INTO pcpmp.tbl_role (id, name, seq) VALUES (22, '超级管理员', 20);
INSERT INTO pcpmp.tbl_role (id, name, seq) VALUES (24, '角色管理员', 10);
INSERT INTO pcpmp.tbl_role (id, name, seq) VALUES (25, '课程资源管理员', 5);
