## Godarena
  2D multiplayer game.
## Features
  Dinamic terrain
  <br/>
  <br/>
  Multiplayer
  <br/>
  <br/>
  AI
## How to install
  go get github.com/lukistar/Godarena
## Dependencies
  https://github.com/veandco/go-sdl2 for the draw part. If you want to use only the game logic you won't need it.
  
### For Windows and SDL2:
  You can just install mingw and daonwload SDL2-devel-2.0.3-mingw.tar.gz (MinGW 32/64-bit) from http://libsdl.org/download-2.0.php
  Extract. 
  If you are using 64bit mingw copy the files from `SDL2-2.0.3/x86_64-w64-mingw32` and merge them with your mingw folder.
  If you are using 32bit mingw copy the files from `SDL2-2.0.3/i686-w64-mingw32` and merge them with your mingw folder.
  In my case I had to put them in C:\mingw-w64\mingw64\x86_64-w64-mingw32 i.e bin to bin, include to include etc.
  Then type `go get -v github.com/veandco/go-sdl2/sdl` in the PowerShell(or whatever you are using).
  Same goes for the rest part of go-sdl2 installation.
## How does it work
  Where you should define your custom stuff(i.e to describe them following the rules, witch you can understand via my examples) -> data/sp/templates/default/
  <br/>
  <br/>
  Where you should place your custom models -> data/sp/models/default/
  <br/>
  <br/>
  Where you should place your own AI code -> data/sp/ai/default/
  <br/>
  <br/>
