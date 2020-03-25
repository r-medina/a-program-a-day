;; from https://stackoverflow.com/questions/24724913/best-way-to-add-per-line-information-visually-in-emacs

(save-excursion
  (let* (
      (window-start (window-start))
      (window-end (window-end)))
    (goto-char window-end)
    (while (re-search-backward "\n" window-start t)
      (let* (
          (pbol (point-at-bol))
          (peol (point-at-eol))
          (raw-char-count (abs (- peol pbol)))
          (starting-column
            (propertize (char-to-string ?\uE001)
              'display
              `((space :align-to 1)
                (space :width 0))))
          (colored-char-count
            (propertize (number-to-string raw-char-count)
              'face '(:background "gray50" :foreground "black")
              'cursor t))
          (one-spacer
            (propertize (char-to-string ?\uE001)
              'display
              `((space :width 1))))
          (two-spacers
            (propertize (char-to-string ?\uE001)
              'display
              `((space :width 2))))
          (final-char-count
            (cond
              ((and
                  (< raw-char-count 100)
                  (> raw-char-count 9))
                (concat one-spacer colored-char-count))
              ((< raw-char-count 10)
                (concat two-spacers colored-char-count))
              (t colored-char-count))) )
        (overlay-put (make-overlay pbol pbol)
          'before-string
          (concat starting-column final-char-count two-spacers) )))))
