# mediumup

CLI-tool for linux that uploads a markdown document to [medium](https://medium.com/) as a draft.


## Installation

Build, or download, the executable and place it somewhere within your $PATH

## Usage

Before the program can be used, you will have to generate an integration token from your medium settings page and have it handy for copy-pasting on the first time the program runs. **IMPORTANT:** the supplied token is saved in PLAIN TEXT within the config file under `~/.config/mediumup/config.json`.

```
Usage: mediaup [OPTION]... TITLE FILE

Available options:
  -t string
    	A comma-separated list of tags
```
