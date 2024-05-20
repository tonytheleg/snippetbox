# Setup

```bash
# Start DB
podman run -d --name lets_go_sql -e MYSQL_ROOT_PASSWORD=letsgo -p 3306:3306 mysql:latest

# Login
mysql -D snippetbox -h localhost -P 3306 --protocol=tcp -u root -p

# swl dump (ned to look at import)
podman exec lets_go_sql /usr/bin/mysqldump -u root -p=letsgo > backup.sql
```