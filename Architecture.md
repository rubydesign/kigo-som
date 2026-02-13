The parsing is going well, the mind looks to the future and starts to see
shapes in the mist

# The big four

At the moment it looks like the work falls into four main categories
- HIR , high level ir
- Memory model
- Calling Convention
- blocks / exceptions

Most seems adaption from kide, but some will have to be worked out when we
get there. Still feeling clunky with the lack of proper oo, so modelling in go
will not be as smooth

## HIR

So it is clear already that what we have from the parser is not workable. An
extra layer is needed for two main reasons.

The first is the total (oh so lovely) recursion of sol. Like any good oo
one can just keep calling on results, in arguments and all, it's a mess.
One would need to implement a runtime stack, which i won't, or get rid of
expressions and turn everything into statements by using temporary variables.
Which is what i will do.

The other is that while we now have a "clean" description, ie no more parse
artifacts, we also don't have the type information we need. Because of the
first reason it's better to introduce this later.

## Memory Model

While in Kide i was concerned with raw memory, now we have Go supplied slices and
base types. How to exactly make an Object out of these is not 100% sure.
but i strongly feel that a *universal* layout will be the way. Ie one object to
rule them all. (Which was off course not the case in Kide, which was more like mri)

I think this is at least a good start because of Go's string typing and slow
casting. I read that to be typesafe at runtime, cast are checked in go at runtime
(unlike c), resulting in substantial slowdown when commonly used. Will have to
experiment, but for now i can easily live with wasting some space.

Current thinking is that we will need the type, and each of the base types, ie
int, float, array and string.

In terms of management we start with almost nothing. Nothing means to use built
in, almost nothing means a thin layer above that for statistics (which i saw
the vm is supposed to collect)
**Maybe in the future pre-allocating chunks could be a thing. Not sure about weak
refs in go yet, but i think  i read that finalizers exist

## Calling convention

Just like the object, calling has to be general. So only one signature for
all and every generated function. Off course a little bit to be determined, but
some is clear.

Every method
- returns an object.
- receives an array of objects
- receives a Message (context), possibly double linked

At least atm i think the linked list will help to resolve the oh so popular
blocks relatively painlessly. Dynamically in the baseline, but at least without
walking the stack. Apparently go comes with stack walking utilities, so maybe
that is a worthwhile experiment one day, but for now keeping it simple.

## Blocks + exceptions

So it looks a bit like som skipped exceptions, making life a lot easier. Also
have to check how much passing around of blocks the basic code (tests and
benchmarks) includes.

First version will be "by the book". This means a block dynamically resolves
variables in the scope it executes. This is off course not ideal, but should
work. Since blocks (like methods) store all variable names, code to access them
"just" has to go to the current scope (Message) and it's type info to get indexes
for the variables. Same as methods, just at runtime.

Since a very very common use of blocks is not to capture them but to execute
them in the scope of the method being called, the resolution can actually be done
at compile-time, since the scope is known (ie the scope the block is defined in).

Also, in som (as in smalltalk and a degree ruby), many of the block execution
happens in "well known" methods. In som/smalltalk especially those are ifTrue,
ifFalse and the various loops. These can "easily" be inlined and thus blocks can
be sidestepped / inlined away at compile time.
