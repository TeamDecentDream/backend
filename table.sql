CREATE TABLE member
(
    id       INT AUTO_INCREMENT,
    name     VARCHAR(100) NOT NULL,
    email    VARCHAR(100) NOT NULL,
    address  VARCHAR(100),
    reg_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE authorities
(
    member_id INT,
    authority VARCHAR(100),
    PRIMARY KEY (member_id, authority),
    FOREIGN KEY (member_id) REFERENCES member (id) ON DELETE CASCADE
);

create table notification
(
    id INT AUTO_INCREMENT,
    title VARCHAR(1000) NOT NULL ,
    contents VARCHAR(4000) NOT NULL,
    author_id int,
    reg_date timestamp default current_timestamp,
    update_date timestamp default current_timestamp,
    primary key (id),
    foreign key (author_id) references member (id) on delete cascade
);

create table significant
(
    id INT AUTO_INCREMENT,
    contents VARCHAR(4000) NOT NULL,
    author_id int,
    warn int default 0,
    reg_date timestamp default current_timestamp,
    update_date timestamp default current_timestamp,
    primary key (id),
    foreign key (author_id) references member (id) on delete cascade
);

create table member_request(
                               id INt auto_increment,
                               member_id int not null,
                               address varchar(400) not null,
                               reg_date timestamp default current_timestamp,
                               confirm_date timestamp,
                               state int default 0,
                               primary key (id),
                               foreign key (member_id) references member(id) on delete cascade
);

create table attendance (
                            id int auto_increment,
                            member_id int null,
                            enter_time timestamp default current_timestamp,
                            leave_time timestamp,
                            primary key (id),
                            foreign key (member_id) references member(id) on delete cascade
);

create table evaluate(
                         id int auto_increment,
                         member_id int null,
                         q1 int null,
                         q2 int null,
                         q3 int null,
                         q4 int null,
                         q5 int null,
                         q6 int null,
                         note varchar(4000),
                         reg_date datetime default current_timestamp,
                         primary key (id),
                         foreign key (member_id) references member(id) on delete cascade
);

create table transaction(
                            id int auto_increment,
                            title varchar(100),
                            amount int,
                            client varchar(100),
                            sell_buy int,
                            reg_date datetime default current_timestamp,
                            state int default 0,
                            primary key (id)
);

create table todo(
                     id int auto_increment,
                     author_id int not null ,
                     contents varchar(1000) not null,
                     state int default 0,
                     reg_date datetime default current_timestamp,
                     update_date datetime,
                     primary key (id),
                     foreign key (author_id) references member(id) ON DELETE CASCADE
);




# INSERT INTO member (name, email, address) VALUES
#                                               ('John Doe', 'john@example.com', '0x656710Bd0B06D5D6836816c961CF984BeCa4f554'),
#                                               ('Jane Smith', 'jane@example.com', '0x222710Bd0B06D5D2436816c961CF984BeCa4f554'),
#                                               ('Bob Johnson', 'bob@example.com', '0x999710Bd0B06D5D6836816c961CF984BeCa4f554');
#
# insert into authorities (member_id, authority) value ('3','ROLE_USER');
#
# drop table significant;
# drop table transaction;
# drop table attendance;
# drop table evaluate;
# drop table authorities;
# drop table notification;
# drop table member_request;
# drop table todo;
# drop table member;
