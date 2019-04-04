#!/usr/bin/fish
#GNU All-Permissive License
#Copying and distribution of this file, with or without modification,
#are permitted in any medium without royalty provided the copyright
#notice and this notice are preserved.  This file is offered as-is,
#without any warranty.

# This is a Simple fish script to create a bunch of size variants for svg files,
# requires rsvg-convert and fish,
# fish is provided in the 'fish' package on Debian GNU/Linux and rsvg-convert
# is provided under the 'librsvg2-dev' package in Debian GNU/Linux

# variables
set icon_path data/icons
set icon_name minidisc
set project moe.jade.oct2pus.listen

# png creates a PNG file from an svg file 
# argv[1] is the height and width
function png
  for svg in (ls {$icon_path}/raw/{$icon_name}.svg)
    # first sed removes directory, second sed removes file extension
    # this could probably be writen better
    rsvg-convert -w {$argv[1]} -h {$argv[1]} {$svg} > \
    {$icon_path}/{$argv[1]}/{$project}.png
  end
end 

# the program 'loop'
# theres probably a more beautiful way to write this
for x in 16 24 32 48 64 128
  mkdir -p {$icon_path}/{$x}
  png {$x}
end

