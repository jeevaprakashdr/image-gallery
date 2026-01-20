#!/bin/bash

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="host=localhost user=postgres password=postgres dbname=gallery sslmode=disable"
export GOOSE_MIGRATION_DIR=./backend/infrastructure/postgres/migrations

export GALLERY_BUCKET_NAME="images"
export GALLERY_SCALED_BUCKET_NAME="scaled-images"

export MINIO_URL="localhost:9000"
export MINIO_ACCESS_KEY_ID="minioadmin"
export MINIO_ACCESS_KEY_SECRETE="minioadmin"
export MINIO_GALLERY_BUCKET="images"

export WS_ADDRESS=":8081"