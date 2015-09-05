## Godarena
  2D arena game.
## Features
  Coming soon
## How to install
  go get github.com/lukistar/Godarena
## Dependencies
  https://github.com/veandco/go-sdl2 for the draw part. If you want to use only the game logic you won't need it.
#### For Windows with mingw:
  Download SDL2-devel-2.0.3-mingw.tar.gz (MinGW 32/64-bit) from http://libsdl.org/download-2.0.php<br/>
  <br>Extract.<br/> 
  <br>If you are using 64bit copy the files from `SDL2-2.0.3/x86_64-w64-mingw32` and merge them with your mingw folder.<br/>
  <br>If you are using 32bit copy the files from `SDL2-2.0.3/i686-w64-mingw32` and merge them with your mingw folder.<br/>
  <br>In my case I had to put them in C:\mingw-w64\mingw64\x86_64-w64-mingw32 i.e bin to bin, include to include etc.<br/>
  <br>Then type `go get -v github.com/veandco/go-sdl2/sdl` in the PowerShell(or whatever you are using).<br/>
  <br>Same goes for the rest part of go-sdl2 installation.<br/>
  <br>Make sure that you have all files from SDL2_image-devel-2.0.0-mingw's bin(64 bit if you are using 64 bit, 32 bit if you are using 32 bit) in the folder of main.go

