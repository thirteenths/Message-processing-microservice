#!/bin/bash
sleep 10
goose postgres "host=$DB_HOST user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME" up