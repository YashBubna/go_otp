# OTP Verification Module Using Twilio

This project is a secure OTP (One-Time Password) verification module created in Golang using Twilio services. The module allows users to register their phone numbers and verifies their identity through an OTP sent via SMS, enhancing application security and user authentication.

## Features

- **OTP Generation and Verification:** Automatically generates and verifies OTPs sent to the user's phone number via Twilio.
- **Twilio Integration:** Leverages Twilio's SMS services to send OTPs securely.
- **User Registration:** Registers and verifies users by their phone number before adding them to the database.
- **REST API with Gin Framework:** Provides an easy-to-use API for sending and verifying OTPs.

## Tech Stack

- **Language:** Golang
- **API Framework:** Gin
- **SMS Service:** Twilio
- **Environment Variables:** `godotenv` for environment variable management
