# idev

<!-- toc -->

- [`enUrl`](#enUrl)
- [`parJson`](#parJson)
- [`parJsonf`](#parJsonf)

<!-- tocstop -->
Collection of command line tools for development.

## Dependencies

- Cobra:
    - https://github.com/spf13/cobra
- Task:
    - https://github.com/go-task/task

## List of tools to create

### `enUrl`

Encode the give http url.

-

run: `idev --enUrl https://www.glassdoor.ca/Overview/Working-at-Vancouver-Public-Library-EI_IE468025.11,35.htm`
-
output: `https:%2F%2Fwww.glassdoor.ca%2FOverview%2FWorking-at-Vancouver-Public-Library-EI_IE468025.11%2C35.htm`

### `parJson`

Parse stringfied JSON to JSON format and display in console.

- run: `idev --parJson '{\"a\":\"aa\", \"b\": {\"c\":\"cc\"}, \"d\": [1,2,3]}'`
    - output:,
        ```json
        {
          "a": "aa",
          "b": {
            "c": "cc"
          },
          "d": [
            1,
            2,
            3
          ]
      }
      ```

### `parJsonf`

Read the give text file and parse stringfied content to JSON.

- run: `idev --parJsonf fileName.txt`
- output:

    ```json
    {
      "a": "aa"
    }
    ```

## Development

Reference: https://github.com/spf13/cobra-cli/blob/main/README.md

- Install `cobra-cli`

    ```shell
    go install github.com/spf13/cobra-cli@latest
    ```

- Initialize a project
    - cd `my-app/`
    - Run `cobra-cli init` in your root of your project directory

- Add a command to the project
    - Let's say we want to create a command
      called `server`: `cobra-cli add serve`
        - The command will be: `my-app serve`
    - If we want to add a flag

- Use config file for `cobra-cli`
    - There is an existing config file in this repo: `.cobra.yaml`
    - Option 1:
        - Run `cobra-cli --config .cobra.yaml add <command>` in the root
          directory
    - Option 2:
        - Create an alias in your `~/.profile`
            - `alias cobra=cobra-cli --config ./.cobra.yaml`
        - it'll always use the config at where you run the `cobra` command
