## VK Callback API to Telegram

This project provides a solution for integrating VK (formerly known as VKontakte) Callback API with Telegram using the Go programming language.

### Features

- Receive VK events in real-time and forward them to Telegram.
- Support for various VK events, including new messages, wall posts, comments on photos and videos.
- Customizable message formatting for Telegram notifications.
- Easy setup and configuration.

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/ad/vk-callbackapi-to-telegram.git
   ```

2. Install the required dependencies:

   ```
   go get -d -v ./...
   ```

3. Set up the configuration file:

   - Rename the `config.example.json` file to `config.json`.
   - Open `config.json` and enter your VK and Telegram API credentials.

4. Build the application:

   ```
   go build .
   ```

5. Start the application:

   ```
   ./vk-callbackapi-to-telegram
   ```

### Usage

1. Register a VK Callback API server and configure the required events to be forwarded to the application's endpoint.

2. Set up a Telegram bot and obtain the bot token.

3. Configure the `config.json` file with the necessary credentials and settings.

4. Customize the message formatting in the code according to your preferences.

5. Build and start the application.

### Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

### Acknowledgements

- This project is inspired by the need to bridge VK Callback API with Telegram.
- Special thanks to the contributors and maintainers of the project.

### Support

If you find this project helpful, consider giving it a star on GitHub and supporting the maintainer. Thank you!
