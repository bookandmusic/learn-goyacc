# learn-goyacc

这是一个使用 goyacc 实现的语法分析器练习项目。

## 简介

本项目包含两个子项目，分别是一个计算器和一个简单的SQL解析器，均使用 goyacc 编写。该项目旨在学习和实践编译原理中的词法分析和语法分析。

## 安装

在使用本项目之前，请确保您的系统已安装 Go 语言环境。您可以从 Go 官方网站 下载并安装最新版本的 Go 语言。

### 克隆项目

首先，将本项目克隆到本地：

```bash
git clone git@github.com:bookandmusic/learn-goyacc.git
cd learn-goyacc
```

### 安装 goyacc
然后，您需要安装 goyacc 工具。您可以通过以下命令安装：

```bash
go install golang.org/x/tools/cmd/goyacc@latest
```

## 使用

### 生成语法分析器

#### 计算器
在生成计算器的语法分析器之前，确保您在 calc 目录下，然后运行以下命令：

```bash
cd calc
goyacc -o parser.go calc.y
```
该命令将根据 calc.y 文件生成 parser.go 文件。

### SQL 解析器
在生成SQL解析器的语法分析器之前，确保您在 sql 目录下，然后运行以下命令：

```bash
cd ../sql
goyacc -o parse.go sql.y
```

该命令将根据 sql.y 文件生成 parse.go 文件。


## 项目结构

项目结构如下：

```bash
.
├── calc
│   ├── calc.go          # 计算器的主要逻辑
│   ├── calc.y           # 计算器的语法文件
│   ├── lexer.go         # 计算器的词法分析器
│   ├── main             # 计算器的可执行文件
│   ├── parser.go        # 由 goyacc 生成的计算器语法分析器
│   └── parse_test.go    # 计算器的测试文件
├── go.mod               # Go 模块文件
└── sql
    ├── lexer.go         # SQL 解析器的词法分析器
    ├── node.go          # SQL 解析器的语法树节点定义
    ├── parse.go         # 由 goyacc 生成的 SQL 解析器语法分析器
    ├── parse_test.go    # SQL 解析器的测试文件
    ├── sql.go           # SQL 解析器的主要逻辑
    └── sql.y            # SQL 解析器的语法文件
```
