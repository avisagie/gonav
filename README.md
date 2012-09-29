
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


Emacs
=====

It works particularly well with emacs when you add the this to your
.emacs file:  
`(add-hook 'shell-mode-hook 'compilation-shell-minor-mode)`

To use, just type `M-x shell`, followed by Enter. Run gonav in that
shell, with  
`~/projects/gonav/gonav -root=/home/me/go/src`

Type the name of the function or type you're looking for, and the
files and line numbers will become clickable links. Epic.
