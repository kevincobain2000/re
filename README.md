<p align="center">
  <a href="https://github.com/kevincobain2000/re">
    <img alt="re" src="https://imgur.com/Jmrdvjp.png" width="360">
  </a>
</p>
<p align="center">
  Stop going back and forth to the README for instructions. <br>
  CLI to execute commands in README.md on local or Github. <br>
</p>


![re](https://imgur.com/zFiYhgO.png)


**Hassle Free:** Easy install.

**About:** Hit `re` command, select multiple commands and execute.

**Platforms:** Supports (arm64, arch64, Mac, Mac M1, Ubuntu and Windows).

**Supports:** Any README.md on local or from Github and Github Enterprise URLs.


## Usage

```sh
# By default it looks README.md file in directory you are executing from
re

# By default it looks for sh, bash, powershell etc.. code blocks
re -t sh # filter by language i.e only commands inside ` ` ` sh blocks are parsed
```

or from URL

```sh
# automatically picks README.md from main, master or develop branch
re https://github.com/kevincobain2000/re

# or direct link to the README.md
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
