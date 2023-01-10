# World Actor System for Go

WAS is a distributed message driven system based on actors, systems and a world state. 

# Concepts

WAS provides a dist
# Actor

Each actor provides states, can receive messages and emit new messages. 
An actor typical lives inside a module which acts as a namespace for actors. 


## State

The state is stored in each actors state store. Each actor has access to it's own state. To modify the state of other actors the actor need to send a change message to that actor and the actor will then modify the state. Messages are always send via the world. Each actor has access to the worls.

## System

A system is a state less actor which only provides behavior as a reaction to a message.

## World

The world knows all actors ans systems. It also runs the world tick. It can collect all actors state to a world state.

The world runs a tick which updates the actors in an interval. 
The interval can be freely adjusted by users.


## Distributed

