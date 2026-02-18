# SOOL Simple Object Oriented Language

The SOM has it's language SOL, so called Simple Object Language.
That would have been the perfect name for this layer, alas SOL is not
simple. At least not in the sense of few constructs and easy to process.

Hence this SOOL layer, where we try to keep things to a minimum and also
keep code generation in mind.

## The OO part

The OO part means we still use classes and objects with instances and
inheritance, blocks and methods.

For unification we use a meta-class object to encode class instances and methods.

As mentioned in the architecture, we type objects and types are immutable.
A type implements a class.

## Statements (no expressions)

Also to simplify, we get rid of recursively defined structures. This means a
method (and block for that matter) is a list of statements.
There are just a few kinds of statement:

*Assignment* assigns a Literal, block or method/block call.
*send* calls on an explicit variable, all arguments are variables
*return* returns an explicit variable, returns are explicit

## SSA

The above makes this a sort of SSA, Single Assignment, which means we can
annotate inferred type info to the variables.

As a side note the locals (method and block) are very very similar to instance
variables and thus all will end up with similar forms. Also the arguments
actually.

These should all get typed representation, ie be objects that have type.
And by objects i mean som/sool objects.

## Blocks

I am still learning go, so this is a bit open. Atm i am thinking that a block
is like a literal, ie we assign it to a variable. Then we can pass it around
or call value (with args) on it. A block is a base type like int/float etc.

The main issue that i haven't fully worked out is no-local returns. SOL, like
SmallTalk and ruby uses return to return from the method that calls the block.
This is not possible in go and thus the block calling has to be different from
method calling and check for non-local return and return explicitly from the
method.
