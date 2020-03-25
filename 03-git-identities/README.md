# Git Identities

Not so much, a program, but just something I've been meaning to do fr a while:
figure out my Git identity story.

The other day I saw [this blog post on Hacker
News](https://www.micah.soy/posts/setting-up-git-identities/) - these are my
notes as I go through it.


Steps:
- get rid of existing global config for identity
- require config to be set per project
- generate pgp keys
- save identities to git config
- make alias
- set for individual projects
- git config --global commit.gpgsign true
- git config --global credential.helper = osxkeychain
