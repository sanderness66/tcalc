# tcalc(1)

Kozmix Go, 14-FEB-2021

tcalc


<a name="synopsis"></a>

# Synopsis

```
tcalc expr ...
```


<a name="description"></a>

# Description

**tcalc**
is a simple command line calculator that happens to understand numbers
in time stamp format. In other words, you can use
**tcalc**
to add, substract, multiply or divide ordinary numbers and/or time stamps.
It prints the result of the calculation both as a number (or as seconds, if you
will), and in "hms" notation.

Time durations can be in any of these formats (where
_n_
denotes any integer):

.TS
tab(|);
l l.
colon notation|
_n_|seconds
n**:n**|minutes and seconds
n**:n:n**|hours, minutes and seconds

hms notation|
n**s**|seconds
n**mns**|minutes and seconds
n**hnmns**|hours, minutes and seconds
.TE

Operators are the usual +, -, *, / and % (and maybe others, I didn't
write the arithmetic evaluator).


<a name="examples"></a>

# Examples


1. How long is 30 minutes plus 65 seconds?

.EX
$ tcalc 30m + 65s
1865 	 31m05s
.EE

2. If you slice an hour and twenty minutes into 30 equal parts, how
long is each part?

.EX
$ tcalc 1h20m / 30
160 	 02m40s
.EE

3. How long does
_Sgt. Patate's_
first 7 inch vinyl record run?

.EX
$ tcalc 03:17 + 01:47 + 01:47 + 02:45
576 	 09m36s
.EE

4. If I have a video clip that runs for 1h, 23m and 45s and split it
up in 10m chunks, how many 10m chunks will I get?

.EX
$ tcalc 1h23m45s / 10m
8 	 08s
.EE

Note that the
_08s_
in this case doesn't mean anything, but
**tcalc**
doesn't know that. It's up to you to pick the result you need. Also note that
**tcalc**
rounds the answer down to the nearest integer.

5. And how much will be left over at the end?

.EX
$ tcalc 1h23m45s % 10m
225 	 03m45s
.EE


<a name="bugs"></a>

# Bugs

We make very little effort indeed checking for correct arguments, and
it's very easy to cause runtime panics by feeding
**tcalc**
with incorrect arguments.

**tcalc**
doesn't understand fractions of a second.


<a name="author"></a>

# Author

svm

