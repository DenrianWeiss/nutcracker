# NutCracker

NutCracker is a Proof-of-Concept [profanity](https://github.com/johguse/profanity/blob/master/Dispatcher.cpp) address cracker.

## How

profanity uses worker id and a base entropy to generate address. However, profanity only get 32 bit entropy and then use meson twister PRNG to expand the entropy to private key.  
This makes an open meet-in-the middle attack possible.

## Other things

This version of program is deliberately slow to prevent abuse. DO NOT ASK FOR FASTER VERSION since the profanity drama is still ongoing.
