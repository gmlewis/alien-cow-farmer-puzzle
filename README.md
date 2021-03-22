# Alien Cow Farmer Puzzle Solver

I came across this video by Chris Ramsay: https://youtu.be/9eKbIvgIdtU
and was amazed by the incredible detail that went into making the $30K puzzle,
so I started browsing on his channel and came across a title I couldn’t resist:

“This UFO puzzle might be my favorite!!”
https://youtu.be/y6gioQ_1FYU

I was thinking that the puzzle looked really challenging to solve, which made
me think, “wouldn’t it be fun to write a program to solve the puzzle and
animate the solution in [Blender](https://blender.org)?”

So I wrote a program in [Go](https://golang.org) (one of my favorite programming languages) to solve
the puzzle and write out a Python program that sets all the keyframes in Blender 2.92.

I also wrote another program to convert the puzzle data to an SVG file that
represents an outline of the two sliding panels. I pulled that file into
[Inkscape](https://inkscape.org) to merge the paths and imported that into
[Onshape](https://onshape.com) after converting it to DXF.
In Onshape I extruded the outline and created the horizontal sliding pins,
then exported the pieces as STL files. (I actually tried importing the SVG
directly into Blender, but the Onshape path was much easier since Blender
handles STL files way better than it does SVG or booleans.)

Onshape model: https://cad.onshape.com/documents/ced52173f7b40ffbba27529a/w/6e5048eae20934913941a56c/e/96c5b78421f18d9f264472d3

After pulling the models into Blender, I searched YouTube for some nice
procedural wood materials and came across this great video by Ducky 3D:
https://youtu.be/EbAdconaRJo

Then I wanted a nice metal material and came across another great video,
this time by Polyfjord:
https://youtu.be/O_spJmmST5I

I ran into troubles hiding the emissive light source and found this video
by Olav3D Tutorials which solved the problem:
https://youtu.be/_lsRQQLFVEY

Finally, I remembered that scripting Blender with the Python API is a bit
of a pain, and I came across this video by Curtis Holt about EasyBPY which
makes it easier:
https://youtu.be/ybnapDe4-Ts

I then wanted to make a Prezi Presenter -style presentation but discovered
that there is no longer a free version. It turns out that Curtis Holt had
used this style of presentation for EasyBPY, but created it entirely in
Blender which I thought was very clever!

I recorded the script for the video using OBS Studio (on Linux Mint
Cinnamon 19.2), then edited it with KdenLive and posted the results to
YouTube here:
https://youtu.be/13rdK3BD5A4

Enjoy!
