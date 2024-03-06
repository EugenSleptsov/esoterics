# Esoterics

The Esoterics module provides access to multiple interpreters of Esoteric languages.

## Available Interpreters

- PlusMinusEquals: A simple interpreter for the PlusMinusEquals esoteric programming language.

## Installation

To use the Esoterics module in your Go project, you can import it as follows:

```go
import "github.com/EugenSleptsov/esoterics"
```

## Usage

To use a specific interpreter, import its corresponding package and call its functions. For example:

```go
import (
    "github.com/EugenSleptsov/esoterics"
)

func main() {
    input := "+-=-=-="
    output := esoterics.PlusMinusEquals(input)
    fmt.Println(output)
}
```

## Contributing

Contributions to the Esoterics module are welcome! If you have suggestions, feature requests, or bug reports, please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the [MIT License](LICENSE).
