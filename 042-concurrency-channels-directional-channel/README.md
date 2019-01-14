# Directional Channels
When using channels as function parameters, you can specify if a channel is meant to only send or receive values. 
This specificity **increases the type-safety** of the program.

This ping function only accepts a channel for sending values. It would be a compile-time error to try to receive 
on this channel.

The pong function accepts one channel for receives (pings) and a second for sends (pongs).

https://gobyexample.com/channel-directions

## Notes
Assigning general (bidirectional) channel type to a specific (directional) type:

general -> specific ✅
specific -> general ❌ because you lose some data
specific -> specific ❌ they are different types