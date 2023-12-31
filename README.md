# Skraper

Welcome to the Skraper repository! Skraper is a versatile bot that helps you scrape data from various sources, primarily focusing on Telegram channels. This README will guide you through the setup process and explain how to run the bot effectively.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Configuration](#configuration)
- [Running the Bot](#running-the-bot)
- [Supported Databases](#supported-databases)
- [Docker Support](#docker-support)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

Before running the Skraper bot, please ensure you have the following prerequisites installed on your system:

1. Go programming language (Golang): Make sure you have Go installed on your machine. You can find installation instructions at [https://golang.org/](https://golang.org/).

2. `.env` File: You will need to create an `.env` file in the root directory of the project and fill in the required environment variables as follows:

```
DB_URI=         # Database connection URI
DB_TYPE=        # Database type (sqlite or mysql)
PHONE_NUMBER=   # Your phone number for Telegram authentication
APP_ID=         # Your Telegram application ID
APP_HASH=       # Your Telegram application hash
2FA=            # Your 2FA (2-Factor Authentication) setting
CODE=           # Code to login to your telegram account 
```

## Configuration

Before running the Skraper bot, you need to set up the `.env` file as described in the "Prerequisites" section. One of the variables in the `.env` file is `CODE`, which plays a crucial role in enabling the bot to work properly.

### Obtaining the `CODE`

1. Start the Bot: When you run the bot for the first time, the Telegram authentication process will be initiated.

2. 2-Factor Authentication (2FA): If you have set up 2-Factor Authentication (2FA) on your Telegram account, you will receive a verification code.

3. Stop the Bot: After starting the bot, but before it begins to execute its main functions, stop the bot manually.

4. Retrieve the `CODE`: Look for the verification code you received during the 2FA process. This code is essential for the bot to authenticate itself with your Telegram account.

### Setting `CODE` in the `.env` File

Once you have obtained the `CODE`, you must add it to the `.env` file in the project's root directory.

```
CODE=YOUR_VERIFICATION_CODE
```

Replace `YOUR_VERIFICATION_CODE` with the actual code you received.

### Restarting the Bot

After setting the `CODE` in the `.env` file, restart the bot using the steps mentioned in the "Running the Bot" section. With the `CODE` correctly configured, the bot will be able to authenticate itself and function correctly, providing you with a seamless scraping experience.

Remember, if you reset your Telegram 2FA settings or encounter any issues related to authentication, you may need to repeat the process of obtaining the `CODE` and updating the `.env` file accordingly.

### `config.yaml` file

The Skraper bot requires a `config.yaml` file to be placed in the root repository. This file helps define the Telegram channels from which data will be scraped. The `config.yaml` file should follow the given scheme:

```yaml
channels:
  - channel_id1
  - channel_id2
  # Add more channel IDs here as needed
```

Simply list the Telegram channel IDs under the `channels` section to start scraping data from those channels.

## Running the Bot

To run the Skraper bot, follow these steps:

1. Edit the `.env` file and fill in the required environment variables with appropriate values.

2. Run the bot using Go:
   ```
   go run ./cmd/main.go
   ```
   Alternatively, you can build the bot first and then execute it:
   ```
   go build ./cmd/main.go
   ./main
   ```

3. The bot will save session files from Telegram in a folder called `session`.

4. If the bot requires a Telegram code for 2FA, start the bot without setting the `CODE` variable in the `.env` file. After receiving the code, stop the bot and enter the code in the `.env` file.

## Supported Databases

Skraper currently supports two databases: SQLite and MySQL. Choose the appropriate database type and set the `DB_TYPE` environment variable accordingly in the `.env` file.

## Docker Support

Docker support for the Skraper bot is currently under development and will be added in the future. Stay tuned for updates and announcements!

## Contributing

Contributions to Skraper are welcome! If you have any bug fixes, feature enhancements, or suggestions, feel free to open an issue or submit a pull request. Please adhere to the existing code style and guidelines.

## License

Skraper is licensed under the [MIT License](LICENSE). Feel free to use and modify the code as per the terms of the license.

Thank you for using Skraper! If you have any questions or need assistance, don't hesitate to reach out to the maintainers of this repository. Happy scraping!