(fn calc-fuel [mass]
  (- (math.floor (/ mass 3) ) 2))

(comment "yo")
(fn rec-fuel [mass]
  (let [fuel (calc-fuel mass)]
    (if (<= fuel 0)
        0
        (+ fuel (rec-fuel fuel)))))

(print (rec-fuel 14))
(print (rec-fuel 1234566343213213))
