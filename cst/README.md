
# Concrete Syntax Tree

A concrete tree has typed nodes, ie typed in the language, here go, ie structs.
The more well known Abstract Syntax Tree (ast) has abstract nodes (usually just
one type, that has the necessary type encoded as a field).

I like concrete trees more  because they are less abstract (pun intended) and
one doesn't always have to check for the type. In other words one can write
concrete function implementations and doesn't need switches. Granted, one can
see this as 6 vs half a dozen, but for me it's more 4 vs 8.

## Structs mirror grammar

So we build this level by basically having a struct per grammar rule. I've copied
the rules into the files.  Maybe someday it will be a file per rule, but for now
i've banged a reasonable amount together.

The Methods creating the structure are very uniformly named, as we are working
off the antlr api. Ie for a grammar rule *result* inside a BlockBody, there will
be a method call *MakeBlockBody*, and it will use an antlr method *Result*
and in turn call a *MakeResult*
