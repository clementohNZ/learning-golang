package main

/*
Method Sets
https://golang.org/ref/spec#Method_sets

Method sets determine what methods attach to a TYPE. It is exactly what the name says: What is the set of
methods for a given type? That is its method set.

IMPORTANT: “The method set of a type determines the INTERFACES that the type implements.....”

-> a NON-POINTER RECEIVER works with values that are POINTERS or NON-POINTERS.
-> a POINTER RECEIVER only works with values that are POINTERS.

Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T

Details:
--------

A type may have a method set associated with it. The method set of an interface type is its interface.
The method set of any other type T consists of all methods declared with receiver type T. The method set of
the corresponding pointer type *T is the set of all methods declared with receiver *T or T (that is, it also
contains the method set of T). Further rules apply to structs containing embedded fields, as described in the
section on struct types. Any other type has an empty method set. In a method set, each method must have a
unique non-blank method name.

The method set of a type determines the interfaces that the type implements and the methods that can be called
using a receiver of that type.
 */
func main() {
}
