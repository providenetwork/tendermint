import fileinput
import re

# This script goes through the provided file, and replaces any " \#<number>",
# with the valid mark down formatted link to it. e.g.
# " [\#number](https://github.com/providenetwork/tendermint/pull/<number>)
# Note that if the number is for a an issue, github will auto-redirect you when you click the link.
# It is safe to run the script multiple times in succession. 
#
# Example usage $ python3 linkify_changelog.py ../CHANGELOG_PENDING.md
for line in fileinput.input(inplace=1):
    line = re.sub(r"\s\\#([0-9]*)", r" [\\#\1](https://github.com/providenetwork/tendermint/pull/\1)", line.rstrip())
    print(line)