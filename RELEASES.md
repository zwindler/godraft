# godraft RELEASES

## 2.1.0

* Lands slider also dynamically enforces limits (min/max cards in deck)
* Add gitpod file
* Add a release file to track improvements

## 2.0.0

* Major rework, with 2 screens now:
  * First screen asks you the format (draft, standard, commander) and game style (aggro, midrange, control)
  * Second screen is the same page as before, but adapted according to the previous choices
* Fixed issues with land suggestions (edges cases with sliders)

## 1.1.0

* Added rules to better deal with splashed colors
  * 2-3 colors, 3 lands min per color
  * 4-5 colors, 2 lands min per color
* Dynamic sliders, with lands number suggestion according to the number of non land cards

## 1.0.0

* First working version
* Only for drafts (40 cards)
* Basic logging server-side
* Basic CSS
* Basic field validations
* Add a makefile and a Dockerfile to ease build/deployment
