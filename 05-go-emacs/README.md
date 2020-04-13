# go + emacs

I've struggled for many years to get emacs and go playing nicely. I finally made
a lot of progress on it!

Go has a [known history of making changes that break text
editors](https://about.sourcegraph.com/go/gophercon-2019-go-pls-stop-breaking-my-editor).
Nevertheless, something else was still going on.

Part of the problem was that as tools were updated, I didn't pull those updates
into my environment.

This happened both from my having faulty
[`el-get`](https://github.com/dimitri/el-get) configs which werent pulling in
the latest `go-mode` versions and from not having updated my go tool chain.

These are some of the tools I realized I need to have updated:

```
golang.org/x/tools/cmd/...
github.com/rogpeppe/godef
github.com/nsf/gocode
honnef.co/go/tools/cmd/staticcheck
```

The main one there is `golang.org/x/tools/cmd/...` which contains things like
`guru`, `goimports`, and `gopls`.


```elisp
(require 'go-mode)

(add-to-list 'auto-mode-alist '("\\.go\\'" . go-mode))
(add-hook 'before-save-hook 'gofmt-before-save)

(add-hook 'go-mode-hook 'flycheck-mode)

(add-hook 'go-mode-hook 'lsp-deferred)
```

Getting rid of `auto-complete` in favor of `company` (and combining that with
`lsp`) has made the autocompletion experience much nicer.

My emacs config [lives here](https://github.com/r-medina/.emacs.d).
