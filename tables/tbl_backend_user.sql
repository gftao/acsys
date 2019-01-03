DROP TABLE pcpmp.tbl_backend_user;
CREATE TABLE pcpmp.tbl_backend_user (id int NOT NULL AUTO_INCREMENT, real_name varchar(255) NOT NULL, user_name varchar(255) NOT NULL, user_pwd varchar(255) NOT NULL, is_super tinyint(1) DEFAULT 0 NOT NULL, status int DEFAULT 0 NOT NULL, mobile varchar(16) NOT NULL, email varchar(256) NOT NULL, avatar varchar(256) NOT NULL, PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8 DEFAULT COLLATE=utf8_general_ci;
INSERT INTO pcpmp.tbl_backend_user (id, real_name, user_name, user_pwd, is_super, status, mobile, email, avatar) VALUES (1, '��������Ա', 'admin', 'e10adc3949ba59abbe56e057f20f883e', true, 1, '18612348765', 'huiyjr@126.com', '/static/upload/1.jpg');
INSERT INTO pcpmp.tbl_backend_user (id, real_name, user_name, user_pwd, is_super, status, mobile, email, avatar) VALUES (2, '����', 'luolan', 'e10adc3949ba59abbe56e057f20f883e', false, 1, '18612348765', 'huiyjr@126.com', '/static/upload/1.jpg');
INSERT INTO pcpmp.tbl_backend_user (id, real_name, user_name, user_pwd, is_super, status, mobile, email, avatar) VALUES (3, '��ͨ�û�', 'nxy', 'e10adc3949ba59abbe56e057f20f883e', false, 1, '', '', '/static/upload/1.jpg');
INSERT INTO pcpmp.tbl_backend_user (id, real_name, user_name, user_pwd, is_super, status, mobile, email, avatar) VALUES (4, '��ͨ�û�', 'pufa', 'e10adc3949ba59abbe56e057f20f883e', false, 1, '', '', '/static/upload/20190103105143.jpg');
INSERT INTO pcpmp.tbl_backend_user (id, real_name, user_name, user_pwd, is_super, status, mobile, email, avatar) VALUES (5, '����', 'lisi', 'e10adc3949ba59abbe56e057f20f883e', false, 1, '', '', '');