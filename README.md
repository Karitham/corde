# corde

[![report card](https://goreportcard.com/badge/github.com/Karitham/corde)](https://goreportcard.com/report/github.com/Karitham/corde)
[![badge](https://pkg.go.dev/badge/github.com/Karitham/corde)](https://pkg.go.dev/github.com/Karitham/corde)

---

corde is a discord webhook callback API wrapper. It aims to have an high level API to simplify the user's life as much as possible.

It contains many functions to help write as clean and concise code as possible, such as `f` suffixed functions which act like `Sprintf`.

It also has builders of some kinds, such as embed and response builders, which make error handling and responding a breeze.

All those don't hinder the fact that you can use corde with as much control as you want.

Be aware that breaking changes will happen as of now, at least until v1.

## Examples

The most basic one is [**bongo**](_example/bongo/main.go)

Then comes [**todo**](_example/todo/) that shows off subcommands routing and options

And then we have [**moderate-myself**](_example/moderate-myself/main.go) which is able to show and delete currently available commands.
It demonstrates the usage of components such as buttons.

Corde is also actively used to rewrite and develop another discord bot of mine called [**WaifuBot**](https://github.com/Karitham/WaifuBot/) (for now see the corde branch)

## Install

```sh
go get github.com/Karitham/corde
```

## Why

Having used most go discord libs available, I wanted something very lightweight yet high-level, that could just run on a lambda without state, caching or long start-up times.

Corde has a single dependency as of yet, and it's just a codegen utility.

The common libs out there didn't really fit that purpose, and neither did they implement the webhook callback API available.

It's from that usecase that corde is born.

Corde means *rope* in french, because discord's API inspires exactly that /s
