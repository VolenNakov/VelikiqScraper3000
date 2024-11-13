#!/bin/bash
set -e

# Variables
APP_DIR="/opt/myapp"
SERVICE_FILE="/etc/systemd/system/myapp.service"
ENV_FILE="/opt/myapp/.env"

# Create application directory if it doesn't exist
sudo mkdir -p "$APP_DIR"

# Copy the binary to the application directory
sudo cp myapp "$APP_DIR/"
sudo chmod +x "$APP_DIR/myapp"

# Write environment variables to a secure file
echo "Writing environment variables to $ENV_FILE"
sudo tee "$ENV_FILE" > /dev/null <<EOL
PORT=${APP_PORT}
DATABASE_URL=${DATABASE_URL}
EOL

# Secure the environment file
sudo chmod 600 "$ENV_FILE"

# Create systemd service file
echo "Setting up systemd service at $SERVICE_FILE"
sudo tee "$SERVICE_FILE" > /dev/null <<EOL
[Unit]
Description=My Go Application
After=network.target

[Service]
Type=simple
User=scraper
WorkingDirectory=$APP_DIR
ExecStart=$APP_DIR/myapp
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
