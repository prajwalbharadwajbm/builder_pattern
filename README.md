# builder_pattern
This project is a simple implementation of the Builder design pattern in Go. The Builder pattern is used to construct complex objects step by step. It separates the construction of a complex object from its representation, allowing the same construction process to create different representations.

## Files
### builder_interface.go
This file defines the Builder interface, which specifies the methods required to build a house.

### normalBuilder.go
This file contains the implementation of the NormalBuilder, which constructs a normal house with wooden windows and doors and two floors.

### tentBuilder.go
This file contains the implementation of the TentBuilder, which constructs an tent with cloth windows and doors and one floor.

### house.go
This file defines the House struct, which represents the product built by the builders.

### mestri.go
This file defines the Director struct, which orchestrates the building process by directing a builder to construct a house.

### main.go
This file contains the client code that demonstrates how to use the Builder pattern to construct different types of houses.

## Usage
To run the example:
```bash
go run main.go
```

## Conclusion
The Builder pattern provides a flexible solution for constructing complex objects. By encapsulating the construction process within separate builder objects, the pattern allows for the creation of different variations of products without altering the client code.

## Reference
Inspired By: [Refactoring Guru - Builder in Go](https://refactoring.guru/design-patterns/builder/go/example#)

