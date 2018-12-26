#!/usr/bin/env bash

until mysql -h ${MARIADB_HOST} -uroot -p${MYSQL_ROOT_PASSWORD} -e "SELECT 1"; do
    >&2 echo "mariadb is unavailable - sleeping"
    sleep 2
done
>&2 echo "mariadb is up - connected"

echo "now init database..."

echo "create user skills..."
mysql -h ${MARIADB_HOST} -uroot -p${MYSQL_ROOT_PASSWORD} -e "CREATE USER 'skills'@'%' IDENTIFIED BY '${SKILLS_MARIADB_PWD}';"

echo "create database skills..."
mysql -h ${MARIADB_HOST} -uroot -p${MYSQL_ROOT_PASSWORD} -e "CREATE DATABASE IF NOT EXISTS skills;"

echo "grant all privilege..."
mysql -h ${MARIADB_HOST} -uroot -p${MYSQL_ROOT_PASSWORD} -e "GRANT ALL ON skills.* TO 'skills'@'%';"