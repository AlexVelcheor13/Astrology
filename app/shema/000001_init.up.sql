CREATE TABLE apod
(
    date            date not null ,
    copyright          varchar(255) not null,
    explanation      varchar(2500) not null,
    hd_url      varchar(255) not null ,
    media_type varchar(20) not null,
    service_version      varchar(255) not null,
    title      varchar(255) not null,
    url      varchar(255) not null
);
CREATE UNIQUE INDEX pictures_date_idx ON apod ("date");