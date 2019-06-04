#!/usr/bin/fish
#GNU All-Permissive License
#Copying and distribution of this file, with or without modification,
#are permitted in any medium without royalty provided the copyright
#notice and this notice are preserved.  This file is offered as-is,
#without any warranty.

# variables
set app_id moe.jade.oct2pus.listen
set -x GO111MODULE on
set translations po/*.po


# create_translations ...creates translations for the app
function create_translations
    for po in (ls po/*.po)
        set name (basename {$po})
        set lang {$name}%.*
        set po/locale/{$lang}/LC_MESSAGES/{$app_id}.mo
        mkdir -p po/locale/{$lang}/LC_MESSAGES/
        msgfmt -c -v -o {$mo} {$po}
    end
end

# Build the app
cd src/; go build
if count {$translations} > /dev/null
    create_translations
end