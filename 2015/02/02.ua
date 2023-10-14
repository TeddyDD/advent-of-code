Input ← &fras "input"

Parse ← ↯ ¯1_3 ⊜parse∵(/×≠"\n\rx").

Slack ← ⊢⊏⍏.≡/×
Paper ← +Slack∶×2/+≡/×.⍜⍉'↘¯1[⍥(↻1.)2]
Ribbon ← +⊙/×/++.↙2⊏⍏..

First ← /+≡Paper Parse
Second ← /+≡Ribbon Parse

First Input
Second Input

~~~
⍤ . =58 Paper 2_3
⍤ . =+58 43 Paper Parse $ 2x3x4
                        $ 1x1x10

⍤ . =34 Ribbon 2_3_4
⍤ . =14 Ribbon 1x1x10
~~~
