;;; hl-linum.el --- A minor mode for displaying line number in line

(defun hl-linum-do ()
  (save-excursion
  (end-of-line)
  (let ((eol-floating-column (+ (current-column) 10))
        (ov (make-overlay (point) (point) nil t t)))
    (overlay-put ov
                 'after-string (what-line)))))
