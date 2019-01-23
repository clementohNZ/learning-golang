# gofmt
`gofmt` formats code

# golint
`golint` prints out style mistakes

The suggestions made by golint are exactly that: suggestions. Golint is not perfect, and has both false positives 
and false negatives. Do not treat its output as a gold standard. We will not be adding pragmas or other knobs to 
suppress specific warnings, so do not expect or require code to be completely "lint-free". In short, this tool is not, 
and will never be, trustworthy enough for its suggestions to be enforced automatically, for example as part of a build 
process. Golint makes suggestions for many of the mechanically checkable items listed in Effective Go and the 
CodeReviewComments wiki page.

# govet
`govet` is concerned with correctness