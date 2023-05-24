<p align="center">
  <a href="https://github.com/kevincobain2000/re">
    <img alt="re" src="https://imgur.com/Jmrdvjp.png" width="360">
  </a>
</p>
<p align="center">
  Stop going back and forth to README. <br>
  CLI to execute commands from local or Github. <br>
  <br>
  <a href="https://kevincobain2000-x.medium.com/automating-command-execution-straight-from-readme-md-5880f4a7f8f1">Read on Medium</a>
</p>


![re](https://imgur.com/krlHmBZ.png)
![re](https://imgur.com/BCRgrh5.png)
![re](https://imgur.com/EKJUokU.png)


**Hassle Free:** Easy install on any arch. No dependencies.

**About:** Hit `re` command, select multiple commands and execute.

**Platforms:** Supports (arm64, arch64, Mac, Mac M1, Ubuntu and Windows).

**Supports:** Any README.md on local or from Github and Github Enterprise URLs.

**Supports:** Multi-line commands. And colorful UI based on previous execution status.

## Install

Using curl

```sh
curl -sLk https://raw.githubusercontent.com/kevincobain2000/re/master/git.io.sh | sh
echo '\nexport PATH="$HOME/.re/bin:$PATH"\n' >> ~/.bashrc
```

or via go

```sh
go install github.com/kevincobain2000/re@latest
```

## Usage

```sh
re

# By default it looks for sh, bash, powershell etc.. code blocks
# filter by language i.e only commands inside ` ` ` sh blocks are parsed
re sh
```

OR from URL

```sh
# automatically picks README.md from main, master or develop branch
re https://github.com/kevincobain2000/re

# or direct link any other MD file, even provide custom tags
re custom_tag https://github.com/kevincobain2000/re/blob/master/EXAMPLE.md
```

## Colors

`re` displays commands upon previous execution status as RED or GREEN.

You can clear the colors by using `re clear` flag or `re clear <url>`.

This will clear the colors for the given URL, when no URL is provided then cleared based on the commands in `README.md` in the current dir.

# Change Log

- v1.0 - Initial release
- v1.1 - Add support for `re github.com/...repo`
- v1.3 - Add support for enterprise and direct tree/blob paths of `README.md`. Multi-line parsing. And Tags.
- v1.4 - Colors, commands history of exit codes.
