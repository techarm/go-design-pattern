# go design pattern

## Factroy Pattern

- The simplest of all design patterns
- Crate an instance of an object with sensible default values

## Abstract Factory Pattern

- Anther common creational pattern
- Crate families of related objects without relying on theire concreate classes

## Repository Pattern

- Allows us to change database with ease
- Makes writing unit test much simpler
- An intermediary layer between an application's business logic and data storage

## Singleton Pattern

- Allow us to restrict the instantiation of something to a singular instance
- This pattern is useful when exactly one object is needed to coordinate actions across a system

## Builder Pattern

- Allow us to chain methods

## Adapter Pattern

- Allows us to have different programs (or parts of the same program) to communicate with one another
- Example: a handler gets information from something, and expects it to be in JSON format
- The data comes from two different remtoe sources. One is in JOSN, and one is in XML.
- The adapter allows us to get the XML data and convert it into waht the handler expects.

## Decorator Pattern

- Take an object, and decorate it with additional information
- Incredibly simple in Go (Embed a struct)
