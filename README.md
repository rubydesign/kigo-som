## An experiment to run OO code fast

This time using SOM as input and Go as output. With the hope to reach a decent
level in are-we-fast-yet, a benchmark comparing many languages.

- Ki - short for Kide, a failed attempt to compiler ruby to binary
- Go - the output language used for this transpiler
- SOM - A Simple Object Machine with a good definition and many implementations

The key components being your usual suspects, a parser, one or two intermediate
representations and a runtime (on top of the go runtime)

Obviously, with the transpiling, this is a c-with-classes approach (the c++
predecessor that transpiled to c). I hope to get so many freebies from this
(like garbage collection), while not 100% sure if this is even doable (hence
experiment).

## Key ideas

It's really just the one: Objects have type, not class.

Types "implement" a class at one point in the classes lifetime.

Classes change, types do not.

Types define a memory layout and (callable) functions. Classes define source
level methods. The job of this transpiler being to translate from class source,
to type method, and keep track of it all.

Object may change type (not class) but this may be quite the hefty operation,
if the object has to be moved and references updated, in case the memory
layout changes. So we give memory layout some slack to avoid moving.

## Implications (and experience from before)

Since some years of experience with kide exist, i have a good idea of direction,
but not details. Since i designed a calling convention for kide, some things like
exceptions and blocks are a little blurry. But i also noticed i was writing
kind of typed (c/go/rust..) code in an IR language, so that should be
transferable.

Two very specific points will make it straight here, the object layout (with type)
and the calling, with messages.

### Types

(this is where it gets a bit circular, hold tight)
Every object has a type, or to be more precise, every object instance has an
instance of Type. The type instance, being an object, also. But the type
also has a class (immutable link) and a type from which it was formed by adding
a single instance variables.

Types define the memory layout of an object and thus can be used to compile code
with certainty. Types form a directed graph, with the object type as the root,
and building up in complexity (no of instance variables).

Objects have instance variables, but also indexed data. How to represent this
(a union, several splices..) is not 100% clear yet.

## Parser

The first step was to build a parser, and test it against
SOM. I prefer to use a generator and ideally one that generates go code and
doesn't construct it on the fly, so i use antlr4, as there is a g4
file in the SOM repository.

Now the parser is done, i am writing cst and converting the antlr nodes.

## General intermediate representation (gir)

Since the ast is generic, an intermediate representation is necessary to get the
semantics set up and provide an adequate structure to output code.
Since the go code this outputs is quite high level, i am currently thinking
one level should be enough, we transpile and compile and run.

## Compiler layer

This is just everything needed to get the transpiling an go compiling done.

## Runtime

I'll call the runtime any code necessary to make the generated code work.
This will include something i have called Builtin before, which is functionality
the transpiled code relies on. The *magic* that makes system calls, integer
and string operations work, because those can never be expressed in sol.
