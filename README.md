
Intro
=====

A very simple navigation/exploration tool for go code. Lets you find
all the functions by name, and find all method defined explicitly on
their types. I.e. it does not follow embedding.


Installation
============

Get readline from github.com/igoralmeida/readline-go

Get it to build, somehow. I had to copy libreadline.so from
/usr/lib/i386-linux-gnu/libreadline.so on my debian wheezy system.

go install gonav_cli
bin/gonav -help

Enjoy.
