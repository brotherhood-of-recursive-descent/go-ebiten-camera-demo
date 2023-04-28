# go-ebiten-camera-demo

A detailed explanation on how to implement a 2d camera using the ebitenengine in go

The camera should support the following 
- zoom
- rotation
- position

## first commit

> d82626f1493cf90f5903179bbc0f9517c204a47b

This explains the a standard [ebitengine](https://ebitengine.org) game. 

You have your game struct in [game/game.go](game/game.go) satisfying the [ebiten.Game interface](https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Game) and hand it over to the ebitengine in [main.go](main.go). Voilá, you have yourself a game :smile:

![demo animation](docs/demo.gif)

## A word of translation

Ebitengine uses a left-handed coordinate system. That is that the origin (0,0) is in the upper-left corner. When you place an image on the screen, the image's origin, also the upper-left corner, is placed at the coordinate (0,0).

To move an image you use the `ebiten.DrawImageOptions{}` GeoM.Translate() method. With passing new x,y coordinates to the method ebitengine moves the origin of the image to the provided coordinate. If you know vector math, this is the equivalent of adding a vector to another one.

> todo, add image explaining translation
