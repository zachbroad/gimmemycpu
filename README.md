# GimmeMyCPU

A Go-based application that monitors various retailers for AMD Ryzen CPU availability and sends notifications when stock is detected.

## Features

- Monitors multiple retailers including:
  - Best Buy
  - B&H Photo
  - Newegg
  - AMD Direct
  - TigerDirect
- Tracks AMD Ryzen 5900X and 5950X CPUs
- Notifications via:
  - Windows Toast notifications
  - SMS (via Twilio)

## How It Works

The application continuously checks each retailer's product pages for stock availability. When a CPU is found in stock:
1. Sends a Windows toast notification with a clickable link
2. Sends an SMS notification via Twilio
3. Displays console output with availability information

## Configuration

The application uses the following main configurations:
- Proxy list source: Raw proxy list from GitHub
- Retailer-specific delays to avoid detection
- Custom User-Agent and headers to simulate browser requests
- Configurable out-of-stock detection strings

## Dependencies

- github.com/kevinburke/twilio-go - For SMS notifications
- gopkg.in/go-toast/toast.v1 - For Windows toast notifications

## Setup

1. Clone the repository
2. Update the Twilio credentials in the code:
   - AccountSID
   - AuthToken
   - Set the environment variables:
    - FROM_PHONE_NUMBER
    - TO_PHONE_NUMBER
3. Run the application:
```go
go run main.go
```

## Note

This is a monitoring tool designed for personal use. Please be mindful of retailers' terms of service and rate limiting policies when using this application.
