#!/bin/bash
set -e

# Variables
APP_DIR="$HOME/scraper"
SERVICE_FILE="/etc/systemd/system/scraper.service"
ENV_FILE="$APP_DIR/.env"

# Create application directory if it doesn't exist
mkdir -p "$APP_DIR"

# Copy the binary to the  application directory
cp server "$APP_DIR/"
chmod +x "$APP_DIR/server"

# Write environment variables to a secure file
echo "Writing environment variables to $ENV_FILE"
tee "$ENV_FILE" > /dev/null <<EOL
PORT=${APP_PORT}
DB_FILE=${DB_FILE}
EOL

# Secure the environment file
sudo chmod 600 "$ENV_FILE"

# Create systemd service file
echo "Setting up systemd service at $SERVICE_FILE"
sudo tee "$SERVICE_FILE" > /dev/null <<EOL
[Unit]
Description=Scraper3000
After=network.target

[Service]
Type=simple
User=volen
WorkingDirectory=$APP_DIR
ExecStart=$APP_DIR/server
Restart=always
RestartSec=5
EnvironmentFile=$ENV_FILE

[Install]
WantedBy=multi-user.target
EOL

# Reload systemd and start the service
sudo systemctl daemon-reload
sudo systemctl enable myapp.service
sudo systemctl restart myapp.service

echo "Deployment complete. Service is up and running."
