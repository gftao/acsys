DROP TABLE tbl_role_backenduser_rel;
CREATE TABLE tbl_role_backenduser_rel (id int NOT NULL AUTO_INCREMENT, role_id int NOT NULL, backend_user_id int NOT NULL, created datetime NOT NULL, PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8 DEFAULT COLLATE=utf8_general_ci;
INSERT INTO tbl_role_backenduser_rel (id, role_id, backend_user_id, created) VALUES (62, 24, 2, '2017-12-18 07:35:58');
INSERT INTO tbl_role_backenduser_rel (id, role_id, backend_user_id, created) VALUES (67, 22, 1, '2019-01-03 02:26:49');
INSERT INTO tbl_role_backenduser_rel (id, role_id, backend_user_id, created) VALUES (68, 25, 3, '2019-01-03 02:28:31');
INSERT INTO tbl_role_backenduser_rel (id, role_id, backend_user_id, created) VALUES (69, 26, 4, '2019-01-03 02:28:31');
