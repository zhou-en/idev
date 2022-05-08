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

- run: `idev --enUrl https://www.glassdoor.ca/Overview/Working-at-Vancouver-Public-Library-EI_IE468025.11,35.htm`
- output: `https:%2F%2Fwww.glassdoor.ca%2FOverview%2FWorking-at-Vancouver-Public-Library-EI_IE468025.11%2C35.htm`

### `parJson`

Parse stringfied JSON to JSON format and display in console.

- run: `idev --parJson \{\"a\":\"aa\"\}`
- output:

    ```json
    {
    "a": "aa"
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
