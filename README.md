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
  You can just install mingw and download SDL2-devel-2.0.3-mingw.tar.gz (MinGW 32/64-bit) from http://libsdl.org/download-2.0.php
  <br>Extract.<br/> 
  <br>If you are using 64bit mingw copy the files from `SDL2-2.0.3/x86_64-w64-mingw32` and merge them with your mingw folder.<br/>
  <br>If you are using 32bit mingw copy the files from `SDL2-2.0.3/i686-w64-mingw32` and merge them with your mingw folder.<br/>
  <br>In my case I had to put them in C:\mingw-w64\mingw64\x86_64-w64-mingw32 i.e bin to bin, include to include etc.<br/>
  <br>Then type `go get -v github.com/veandco/go-sdl2/sdl` in the PowerShell(or whatever you are using).<br/>
  <br>Same goes for the rest part of go-sdl2 installation.<br/>
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
