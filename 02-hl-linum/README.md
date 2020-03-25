# `hl-linum`

In moving over to zshell/iTerm2 and getting my emacs environment more up to
date, I realized that I would like more information about what line I'm on while
I'm coding.

Learned about [`progn` in
elisp](https://www.gnu.org/software/emacs/manual/html_node/eintr/progn.html).
This makes the sequence of expressions following the call just happen for side
effects except for last one which becomes return value.

I've spent a lot of time reading through emacs documentation trying to figure
out how to do this. One idea I had was to see how `display-line-numbers` does
its thing, but I couldn't find the underlying code that actually does the
displaying.

[I asked a question on reddit - hopefully someone
answers](https://www.reddit.com/r/emacs/comments/fna78w/how_do_i_display_text_in_my_current_buffer/).

My new lead is to use [the idea behind
linum](https://github.com/emacs-mirror/emacs/blob/master/lisp/linum.el) to just
draw it at the end of the line.

Learned about [`let*`](https://emacs.stackexchange.com/questions/42449/what-does-let-let-asterisk-mean-in-elisp)
which lets sequential bindins depend on the ones prior.

emacs overlays! Thank you reddit! https://emacs.stackexchange.com/questions/2200/what-are-overlays-for-and-how-do-they-differ-from-text-properties

https://www.gnu.org/software/emacs/manual/html_node/elisp/Managing-Overlays.html

https://emacs.stackexchange.com/questions/15078/inserting-before-an-after-string-overlay

https://www.gnu.org/software/emacs/manual/html_node/eintr/save_002dexcursion.html

perhaps want to use `post-command-hook` like in [hl-line](https://github.com/emacs-mirror/emacs/blob/master/lisp/hl-line.el)
