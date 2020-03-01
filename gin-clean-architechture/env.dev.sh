#!/bin/bash

export DB_DRIVER="postgres"
export DB_CONNECTION_STRING="host=localhost port=5432 user=demo dbname=demo password=password sslmode=disable"
export TEST_DB_CONNECTION_STRING="host=localhost port=5432 user=demo dbname=demo_test password=password sslmode=disable"


export JWT_SECRET="demo-jwt-secret"

