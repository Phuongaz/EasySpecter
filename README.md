# EasySpecter

A simple specter for Minecraft Bedrock Edition based on [gophertunnel](https://github.com/Sandertv/gophertunnel)

# How to use

- Step 1
    + Install [EasySpecter](https://github.com/Phuongaz/EasySpecter/releases)
- Step 2
    + Run `EasySpecter.exe`

# Commands
- `xbox login` login xbox account
- `xbox join <ip:port>` Spawn specter to server with xbox account
- `join <name> <ip:port>` Spawn specter with `<name>` and non xbox account (must disable `xbox-auth` in server)
- Specter commands:
    + `specter <specter name> chat <message>`
    + `specter <specter name> move <x> <y> <z>`
    + `specter quit`
    + `specter list`

# Using as library

- Get a package

    ```
    go get github.com/phuongaz/easyspecter
    ```

- Create a specter
    + [Wiki](https://github.com/Phuongaz/EasySpecter/wiki)
