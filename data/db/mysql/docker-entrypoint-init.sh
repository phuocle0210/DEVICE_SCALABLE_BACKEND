envsubs < ./init.sql > ./processed-init.sql
exec mysqld --init-file=./processed-init.sql