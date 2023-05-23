<p align="center">
  <a href="https://github.com/kevincobain2000/re">
    <img alt="re" src="https://imgur.com/Jmrdvjp.png" width="360">
  </a>
</p>
<p align="center">
  Stop going back and forth to the README for instructions. <br>
  Command Line Tool to execute commands in README.md file. <br>
</p>

Select multiple and execute commands in README.md.

*Commands parsed by `re` from `README.md` file.*

![re](https://imgur.com/zFiYhgO.png)


**Hassle Free:** Simple command to get all the commands from `README.md` file. Works with Github URLs.

**About:** By executing `re` command, you will get a list of commands to scroll through.

**How it works:** The tool parses the `README.md` file's markdown in current dir you are on. Analyzes code-blocks and filters `sh`, `bash`, `powershell`, `zsh` etc. commands as selectable prompts.

**Platforms:** Supports (arm64, arch64, Mac, Mac M1, Ubuntu and Windows).

**Supports:** Github.


## Usage

READ the README.md file in the current dir.

```sh
# By default it looks into sh, bash, powershell etc.. code blocks
re
re -t sh # filter by language i.e only commands inside ` ` ` sh blocks are parsed
```

or from URL

```sh
# automatically picks readme from main, master or develop branch
re https://github.com/kevincobain2000/re

# or direct link to the readme
re https://github.com/kevincobain2000/re/blob/master/README.md

# tags work as usual
re -t sh <url>
```

## Install

Using curl

```sh
curl -sLk https://raw.githubusercontent.com/kevincobain2000/re/master/git.io.sh | sh
# add to path in .bashrc or .zshrc
export PATH="$HOME/.re/bin:$PATH"
```

or via go

```sh
go install github.com/kevincobain2000/re@latest
```


# Change Log

- v1.0 - Initial release
- v1.1 - Add support for `re github.com/...repo`
- v1.3 - Add support for enterprise and direct tree/blob paths of `README.md`. Multi-line parsing. And Tags.
