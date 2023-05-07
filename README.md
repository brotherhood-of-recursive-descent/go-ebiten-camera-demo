go-ebiten-camera-demo
=====================

A detailed explanation on how to implement a 2d camera using the ebitenengine in the Golang language.

A basic game
-------------

Writing games is fun your friends tell you. So one day you sit down and start to write your first couple lines of code that resembles a game. Since Golang is your favourite language you checkout game engines and find ebitenengine.org. You skimmed over the documentation, learn about the Update() #TODO(link) and Draw()#TODO(link) method and end up with the following:

> TODO screenshot of a bug 

Not bad for a start. In your code you loaded an image, and drew it on the screen specifying the translation using the DrawImageOptions #TODO(link). 

> TODO draw image of coordinate system with the translation explained (no vectors yet!)

small refactoring
-----------------

Pumped to have something drawn on the screen, you read more about structuring your code and decide to go down the game-object route. While at it you implement input handling and are pleased with a moving bug. 

> TODO screenshot of a bug on background

But you soon realize that the bug moves out of the screen :frowning:. 


screen != world
---------------

 - mouse -> screen
 - objects -> world


enter the camera
-----------------

- center the player (lots of lines in Draw())
- introduce the camera

change position
---------------

 - introduce the concept of position
 - vector math

zooming
------
 - allow zooming 

smoothing
---------
 - adds asymptotic math to make the camera following more smooth

rotation
---------
- adds rotation to the camera

translate between world and screen
-----------------------------------
- utility functions for conversion between the world



References
==========
 * https://ebitengine.org/en/documents/ - useful documentation on how to use the amazing ebitengine library written by Hajime Hoshi.
 * https://github.com/MelonFunction/ebiten-camera - a generic implementation of a camera based on another implementation. It served as a basis for this implementation.
 * [Math for Game Programmers: Juicing Your Cameras With Math](https://www.youtube.com/watch?v=tu-Qe66AvtY) - a detailed explanation on how to make the camera feel nice. We used this a reference for the smoothing example.
