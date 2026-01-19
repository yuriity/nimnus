import os
import sublime
from pathlib import PurePath

FONT = "Iosevka Nerd Font"

#   …/nimnus   main ⇣1    v20.19.6

# ABCDEFGHIJKLMNOPQRSTUVWXYZ
# abcdefghijklmnopqrstuvwxyz
# 1234567890
ABCDEFGHIJKLMNOPQRSTUVWXYZ = ""
abcdefghijklmnopqrstuvwxyz = ""
NUMERALS = 1234567890
SIMILAR = "oO08 iIlL1 g9qCGQ 8%& <([{}])> .,;: ~-_="
DIACRITICS_ETC = "â é ù ï ø ç Ã Ē Æ œ"


class SideBarDuplicateCommand(SideBarCommand):

    def run(self, paths, **kwargs):
        source = self.get_path(paths, **kwargs)
        base, leaf = os.path.split(source)

        # find the file extension
        name, ext = os.path.splitext(leaf)
        if ext != '':
            while '.' in name:
                name, _ext = os.path.splitext(name)
                ext = _ext + ext
                if _ext == '':
                    break

        source = self.get_path(paths, **kwargs)

        input_panel = self.window.show_input_panel(
            'Duplicate As:', source, partial(self.on_done, source), None, None)

        input_panel.sel().clear()
        input_panel.sel().add(
            sublime.Region(len(base) + 1, len(source) - len(ext)))
