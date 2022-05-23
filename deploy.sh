#!/bin/bash
TARGET_USER=temp
TARGET_HOST=temp.local
TARGET_DIR=/home/temp
ARM_VERSION=7
BINARY_NAME=temp

echo "Building for Raspberry Pi..."
env GOOS=linux GOARCH=arm GOARM=$ARM_VERSION go build -o out/"$BINARY_NAME"

echo "Stopping remote service on Raspberry Pi..."
ssh "$TARGET_USER"@"$TARGET_HOST" "sudo systemctl stop temp_hum.service"

echo "Uploading to Raspberry Pi..."
scp -i ~/.ssh/id_rsa out/"$BINARY_NAME" "$TARGET_USER"@"$TARGET_HOST":"$TARGET_DIR"/"$EXECUTABLE"

echo "Starting remote service on Raspberry Pi..."
ssh "$TARGET_USER"@"$TARGET_HOST" "sudo systemctl start temp_hum.service"

echo "Checking if service is running on Raspberry Pi..."
ssh "$TARGET_USER"@"$TARGET_HOST" "sudo systemctl status temp_hum.service"
