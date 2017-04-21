## 301

A simple service to track URL request

Visit `301.yourdomain/link?url=https%3A%2F%2Fgithub.com&a=refScopeA&e=refIdE` on your browser will bring you to `https://github.com` and save a log in a DB.

You can keep track of five references in the link a,b,c,d,e as a query vars (which are not required).

#### SETUP DB

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

# Required, no default value
export APP_DB_DRIVER="mysql"

# Required, no default value
export APP_DB_SOURCE="user:password@tcp(ip:port)/dbname?charset=utf8"

./301
```

#### Use with docker

Build your own **301**

```bash
## 1. Modify the source code in `301.go`

# 2. build container
docker build -t 301 .

# 3. run container (do not forget to modify -e params)
docker run -d -p 8080:8080 -e APP_NETWORK=":8080" -e APP_DB_DRIVER="mysql" -e APP_DB_SOURCE="redirect:redirect@tcp(192.168.1.19:3306)/redirect" -t 301
```

OR pull the image directly from [hub.docker.com](https://hub.docker.com/u/pnicorelli/)

```shell
# 1. pull the image
docker pull pnicorelli/301

# 2. run the container (do not forget to modify -e params)
docker run -d -p 8080:8080 -e APP_NETWORK=":8080" -e APP_DB_DRIVER="mysql" -e APP_DB_SOURCE="redirect:redirect@tcp(192.168.1.19:3306)/redirect" pnicorelli/301
```

To test the container
```shell
# 3. Test
curl localhost:8080
#  if everithings goes 301 should reply with:
#  
#  301 Tracker [v0.1]
#  db_status: ok
```
#### Usage

Call **301** from the shell

```bash
curl "localhost:8080/link?url=https%3A%2F%2Fgithub.com&a=RefA&e=RefE"
```

or visit
[http://localhost:8080/link?url=https%3A%2F%2Fgithub.com&a=RefA&e=RefE](localhost:8080/link?url=https%3A%2F%2Fgithub.com&a=RefA&e=RefE) in your browser

Then query the DB to get the result

```sql
SELECT * FROM redirect
```


#### APIs

###### GET /

Do not accept parameters, reply in `text/plain` with HTTP/1.1 200 OK

```bash
301 Tracker [v0.1]   # APPName [AppVersion]
db_status: ok        # db connection status
```

I use this to check the **301** healt

###### GET /link

Parameters in query string

| name | required | description                         |
|------|----------|-------------------------------------|
| url  | yes      | the encoded url to be redirected to |
| a    | no       | #1 generic reference                |
| b    | no       | #2 generic reference                |
| c    | no       | #3 generic reference                |
| d    | no       | #4 generic reference                |
| e    | no       | #5 generic reference                |

This endpoint reply with

```bash
> HTTP/1.1 301 Moved Permanently                               ##HTTP Header
> Location: https://thewebsiteonthe.url                        ##HTTP Header
<a href="https://thewebsiteonthe.url">Moved Permanently</a>.   ##HTTP Body
```

#### Disclaimer

This is my first approach to GOLANG, 301 is a little piece of code but I do not assure the reliability.
