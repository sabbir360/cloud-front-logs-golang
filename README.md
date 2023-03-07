# CloudFront Log
This project takes input from CloudFront Logs and can filter and discard any unwanted info you want to skip. And generate a filtered **csv** file separately.

Currently supported CSV file headers are (but customizable)
|Date|Time|Size|ClientIp|Host|Endpoint|Status|UserAgent|ResponseTime|
|----|----|----|---------|----|--------|------|----------|-------------|
|date|time|size|client_ip|host|endpoint|status|user_agent|response_time|

## Prerequisites
 - Golang 1.8+

 ## How to run this project
 - Rename `.env.sample` to `.env`.
 - Put your log file to `logs` directory and update the file name to `LOG_FILE` variable.
 - Do same for your desired final output file using `CSV_FILE` variable.
 - Install required packages using `go mod tidy && go get github.com/joho/godotenv`. Assuming you've golang installed into your system
 - Run `go build` and then `main.exe`

 ## Tips for MySQL

If you want to export the CSV directly to your mysql DB do the followings (As insert using python can take immersive time)

- Create database and run below schema
```sql
CREATE TABLE `cflogs` (
  `id` INT AUTO_INCREMENT NOT NULL,
  `date` DATE NOT NULL,
  `time` TIME NOT NULL,
  `size` INT NOT NULL,
  `client_ip` VARCHAR(15) NOT NULL,
  `host` VARCHAR(100) NOT NULL,
  `endpoint` VARCHAR(250) NOT NULL,
  `status` INT NOT NULL,
  `user_agent` TEXT NOT NULL,
  `response_time` FLOAT NOT NULL,
  CONSTRAINT `PRIMARY` PRIMARY KEY (`id`)
);
CREATE INDEX `index_date`
ON `cflogs` (
  `date` ASC
);
CREATE INDEX `index_ep`
ON `cflogs` (
  `endpoint` ASC
);
CREATE INDEX `index_host`
ON `cflogs` (
  `host` ASC
);
CREATE INDEX `index_rst`
ON `cflogs` (
  `response_time` ASC
);
CREATE INDEX `index_status`
ON `cflogs` (
  `status` ASC
);
CREATE INDEX `index_time`
ON `cflogs` (
  `time` ASC
);
```

- Then keep your CSV to `/var/lib/mysql-files/` as MySQL by default only allow secured location.
- Logon to MySQL console `mysql -uuser -ppassword db_name`
- Run this
```sql
LOAD DATA INFILE '/var/lib/mysql-files/your_file.csv' INTO TABLE cflogs FIELDS TERMINATED BY ',' ENCLOSED BY '"' LINES TERMINATED BY '\n' IGNORE 1 ROWS (date, time, size, client_ip, host, endpoint, status, user_agent, response_time);
```
