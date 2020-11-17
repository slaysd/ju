# Jeeseung's Favorite Command Cli (a.k.a ju)
자주 사용하는 커맨드를 cli 툴로 만듦

# Installation

## Required
- Golang

``` bash
go get github.com/slaysd/ju
go install github.com/slaysd/ju
```

# Feature
- [Notify](#notify)
- [Git](#git)


## Notify
특정 커맨드의 실행 여부를 이메일로 report해주는 명령
``` bash
ju notify -- echo hello world
```

## Git

### Git Open Web
현재 git repository에 해당하는 웹페이지를 열기
``` bash
ju git open
```
