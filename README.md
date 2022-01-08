#

<span align="center">

```txt
                        _       
                       | |      
  ___   ___   _ __   __| |  ___ 
 / __| / _ \ | '__| / _` | / _ \
| (__ | (_) || |   | (_| ||  __/
 \___| \___/ |_|    \__,_| \___|
                                
```

</span>

<span align="center">

[![report card](https://goreportcard.com/badge/github.com/Karitham/corde)](https://goreportcard.com/report/github.com/Karitham/corde)
[![badge](https://pkg.go.dev/badge/github.com/Karitham/corde)](https://pkg.go.dev/github.com/Karitham/corde)

</span>

---

corde is a discord webhook callback API wrapper. It aims to have an high level API to simplify the user's life as much as possible.

It contains many functions to help write as clean and concise code as possible, such as `f` suffixed functions which act like `Sprintf`.

It also has builders of some kinds, such as embed and response builders, which make error handling and responding a breeze.

All those don't hinder the fact that you can use corde with as much control as you want.

Be aware that breaking changes will happen as of now, at least until v1.

## Install

```sh
go get github.com/Karitham/corde
```

### Usage

The easiest way to run the examples, or even to run your own is to use [ngrok](https://ngrok.com/) no login or domain required, and the product itself is great.

## Examples

The most basic one is [**bongo**](0_example/bongo/main.go)

Then comes [**todo**](0_example/todo/) that shows off subcommands routing and options

And then we have [**moderate-myself**](0_example/moderate-myself/main.go) which is able to show and delete currently available commands.
It demonstrates the usage of components such as buttons.

Finally, we have [**nft**](0_example/nft/main.go) which is a simple example using user commands & message commands.

Corde is also actively used to rewrite and develop another discord bot of mine called [**WaifuBot**](https://github.com/Karitham/WaifuBot/) (for now see the corde branch)

## Why

Having used most go discord libs available, I wanted something very lightweight yet high-level, that could just run on a lambda without state, caching or long start-up times.

Corde has a single dependency as of yet, a radix tree for routing purposes.

The common libs out there didn't really fit that purpose, and neither did they implement the webhook callback API available.

Corde means *rope* in french, because discord's API inspires exactly that */s*
