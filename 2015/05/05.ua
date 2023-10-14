Input ← &fras $ input
Parse ← ⊜∘∵(/×≠"\n\r").

# First part
Vov ← ≥3/+∊∶"aeiou"
Rep ← ≥1/+≡/=◫2
NotBadStr ← =0/+≠4⊗∶["ab" "cd" "pq" "xy"] ◫2

ChkFirst ← /×⊂⊂⊃⊃Vov Rep NotBadStr
First ← /+≡ChkFirst

# Second Part
Pairs ← ≥2/↥-⊛ ∶⇡⧻.◫2
Pal ← >0/+≡/×=≡⇌.◫3

ChkSecond ← /↧⊂⊃Pairs Pal
Second ← /+≡ChkSecond

# Debug
# ≡{/×.⊂⊂⊃⊃Vov Rep NotBadStr .} Parse Input
# /+≡(⊢).≡{/↧.⊂⊃Pairs Pal .} Parse Input

# Run solutions
First Parse Input
Second Parse Input

~~~
⍤ . ChkFirst "ugknbfddgicrmopn"
⍤ . ChkFirst "aaa"
⍤ . ¬ ChkFirst "jchzalrnumimnmhp"
⍤ . ¬ ChkFirst "haegwjzuvuyypxyu"
⍤ . ¬ ChkFirst "dvszwmarrgswjxmb"
⍤ . ¬ ChkFirst "a"
⍤ . ¬ ChkFirst "nvrslorjpqaasudn"
⍤ . ChkSecond "qjhvhtzxzqqjkmpb"
⍤ . ChkSecond "xxyxx"
⍤ . ¬ ChkSecond "uurcxstgmygtbstg"
⍤ . ¬ ChkSecond "ieodomkazucvgmuy"
⍤ . ChkSecond "skerkykytazvbupg"
⍤ . =258 First Parse Input
⍤ . =53 Second Parse Input
~~~
