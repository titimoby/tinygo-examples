# gba/04-Pixel

This repo contains an example for TinyGo Game Boy Advance.

We use:
* [TinyGo](https://tinygo.org/)
* [Tinyheart](https://github.com/tinygo-org/tinyheart)

## General

This simple app run on Game Boy Advance portable console an heart in pixels.

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

![App](doc/heart.png)

## Build the app

* For mGBA:

`$ GOFLAGS=-mod=mod tinygo build -size short -o bin/heart.elf -target=gameboy-advance main.go ; mv bin/heart.elf bin/heart.gba`

* For VisualBoyAdvance emulator or real GBA console:

`$ tinygo build -size short -o bin/heart.gba -target=gameboy-advance main.go`

## Result

![App](doc/heart.png)

## Build the app

`$ tinygo build -size short -o bin/heart.gba -target=gameboy-advance main.go`

## Test the app/game

Let's run our app on mGBA emulator:

`$ mgba bin/heart.gba`

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