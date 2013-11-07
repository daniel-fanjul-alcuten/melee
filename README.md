Description
===========

Calculates all possible combinations of hits and misses.

It will be slow for large numbers (brute force).

Usage:

<pre>
$ melee -h
Usage of melee: <actions>
 Actions can be: letter:number, i.e. M:100, D:80, B:70, P:50.
 The order matters.
</pre>


Examples:

<pre>
$ melee m:100 d:50
3725 hits (74.500%), 1275 misses (25.500%), 5000 total
</pre>

<pre>
$ melee m:100 d:50 p:50
165425 hits (66.170%), 84575 misses (33.830%), 250000 total
</pre>

<pre>
$ melee m:100 d:100
4950 hits (49.500%), 5050 misses (50.500%), 10000 total
</pre>

<pre>
$ melee M:100 D:500 B:150 P:300
24502500 hits (1.089%), 2225497500 misses (98.911%), 2250000000 total
</pre>

Installation
============

<pre>
export GOPATH=...
go get github.com/daniel-fanjul-alcuten/melee
</pre>
