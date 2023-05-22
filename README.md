<p align="center">
  <a href="https://github.com/kevincobain2000/re">
    <img alt="re" src="https://imgur.com/Jmrdvjp.png" width="360">
  </a>
</p>
<p align="center">
  Stop going back and forth to the README for instructions. <br>
</p>

*Commands parsed by `re` from `README.md` file.*

![re](https://imgur.com/DgrXIVs.png)

*Command execution after selection*

![re](https://imgur.com/Y9HUHO0.png)


**Hassle Free:** Simple command to get all the commands from `README.md` file.

**About:** By executing `re` command, you will get a list of commands to scroll through.

**How it works:** The tool parses the `README.md` file's markdown in current dir you are on. Analyzes code-blocks and filters `sh`, `bash`, `powershell`, `zsh` etc. commands as selectable prompts.

**Platforms:** Supports (arm64, arch64, Mac, Mac M1, Ubuntu and Windows).

**Supports:** Github.


## Usage

READ the README.md file in current dir

```
re
```

or from URL

```
re https://github.com/kevincobain2000/re
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
- v1.2 - Add support for `re github.com/...repo`
