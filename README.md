## 301

A simple service to track URL request

Visit `301.yourdomain/link?url=https%3A%2F%2Fgithub.com&a=refScopeA&e=refIdE` on your browser will bring you to `https://github.com` and save a log in a DB.

You can keep track of five references in the link (a,b,c,d,e query vars) which are not required.


### Build a container

```bash
# build container
docker build -t 301 .

# run container
docker run -d -p 8080:8080 -e APP_NETWORK=":8080" -e APP_DB_DRIVER="mysql" -e APP_DB_SOURCE="redirect:redirect@tcp(192.168.1.19:3306)/redirect" -t 301

```


### SETUP DB

Create a table to store the logs:

```sql
CREATE TABLE `redirect` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `url` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `a` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci,
  `b` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci,
  `c` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci,
  `d` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci,
  `e` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci,
  PRIMARY KEY (`id`)
);
```

#### Compile and Run

```go
go build 301.go

# Not required, default to ":8080"
export APP_NETWORK=":8080"

# Required, no default
export APP_DB_DRIVER="mysql"

# Required, no default
export APP_DB_SOURCE="user:password@tcp(localhost)/dbname?charset=utf8"

./301
```


#### Usage

Call **301** on the shell
```bash

curl "localhost:8080/link?url=https%3A%2F%2Fgithub.com&a=RefA&e=RefE"
```

or visit
[http://localhost:8080/link?url=https%3A%2F%2Fgithub.com&a=RefA&e=RefE](localhost:8080/link?url=https%3A%2F%2Fgithub.com&a=RefA&e=RefE) in your browser

Then query the DB to get the result

```sql
SELECT * FROM redirect
```

#### Disclaimer

This is my first approach to GOLANG, 301 is a little piece of code but I do not assure the reliability.
