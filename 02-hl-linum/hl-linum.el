(let ((overlay (make-overlay (point-max) (point-max) nil t t))
      (line-num (number-to-string (line-number-at-pos))))
  (overlay-put overlay 'intangible t)
  (overlay-put overlay 'before-string
               (concat
                (propertize " " 'display `(space :align-to (- right-fringe ,(+ 2 (length line-num)))))
                line-num))
  overlay)

