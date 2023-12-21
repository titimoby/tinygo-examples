# gba/01-Font

This repo contains an example for TinyGo Game Boy Advance.

We use:
* [TinyGo](https://tinygo.org/)
* [TinyFont](https://github.com/tinygo-org/tinyfont)

## General

This simple app run on Game Boy Advance portable console and display text messages.

## Pre-requisites

Install Go in 1.16 version minimum.

Install [TinyGo](https://tinygo.org/getting-started/install/).

Install [mGBA](https://tinygo.org/getting-started/install/macos/) emulator.

## Run the app (during development)

```
$ tinygo run -target=gameboy-advance main.go
tinygo:ld.lld: warning: lld uses blx instruction, no object with architecture supporting feature detected
```

## Result

![App](doc/font.png)

## Build the app

* For mGBA, VisualBoyAdvance emulator or real GBA console:

`$ GOFLAGS=-mod=mod tinygo build -o font.gba -target=gameboy-advance main.go`

## Test the app/game

Let's run our app on mGBA emulator:

`$ mgba font.gba`

### mGBA Controls

Controls are configurable in the **settings** menu of **mGBA**. Many game controllers should be automatically mapped by default. 
The default keyboard controls are as follows:

```
A: X
B: Z
L: A
R: S
Start: Enter
Select: Backspace
```