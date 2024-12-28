
# Procex - Application Management Discord Integration

**Procex** is an innovative solution that transforms your Discord server into a powerful environment for controlling your applications in real-time. With Procex, you can manage your system's processes, interact with applications, and collaborate directly through Discord bot commands.

**Warning ⚠⚠: STILL IN DEVELOPMENT SO ISSUES OR DYSFUNCTIONALITIES MIGHT BE PRESENT**



## Features

- **Real-Time Process Control:** Start, stop, restart, and list processes directly from Discord.
- **Workspace Management:** Create and manage workspaces for your processes.
- **Authentication:** Securely authenticate users to access and manage workspaces.
- **Supports PM2 Environments:** Compatible with environments supported by PM2 (e.g., Node.js, JavaScript, and other compatible runtimes).

## Supported PM2 Environments
- **Node.js**
- **JavaScript**
- **Python**
- **Go**
- **Other PM2 Supported Runtimes**

## Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/yourusername/procex.git
   cd procex
   ```

2. **Build the Project**
   Ensure Go is installed, then run:
   ```bash
   go build
   ```

3. **Set Up Environment Variables**
   Create a `config.env` file in the `bot` directory with the following variables:
   ```env
   TOKEN="your token here"
   ```

4. **Run the Bot**
   ```bash
   ./procex
   ```

5. **Invite the Bot to Your Server**
   Use the OAuth2 URL provided in your Discord developer portal to invite the bot to your server.

## Usage

### Available Commands

- **`/start`** - Starts a specific process using a project directory and command.
- **`/stop`** - Stops a specific process.
- **`/restart`** - Restarts a specific process.
- **`/list`** - Lists all running processes.
- **`/auth`** - Authenticates a user for workspace access.
- **`/workspaces`** - Lists existing workspaces.
- **`/workspace_create`** - Creates a new workspace with specific permissions.
- **`/pull`** - Pulls a repository and runs code.
- **`/project_create`** - Creates a new project.


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [DiscordGo](https://github.com/bwmarrin/discordgo) for providing the framework to interact with Discord.
- The open-source community for inspiration and support.

## Support

For issues or feature requests, please open an [issue](https://github.com/Betryx/procex/issues) on GitHub or contact the project maintainer.
