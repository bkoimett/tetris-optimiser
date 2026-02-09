# tetris-optimiser
Develop a program that receives only one argument, a path to a text file which will contain a list of tetrominoes and assemble them in order to create the smallest square possible


tetris-optimizer/
├── main.go                 # Entry point
├── go.mod                  # Go module file
├── internal/
│   ├── parser/            # File parsing logic
│   │   ├── parser.go
│   │   └── parser_test.go
│   ├── solver/            # Tetris solving algorithms
│   │   ├── solver.go
│   │   └── solver_test.go
│   └── models/            # Data structures
│       └── tetromino.go
├── cmd/
│   └── tetris-optimizer/
│       └── main.go        # Alternative entry point (optional)
└── test_files/            # Test input files
    ├── sample.txt
    ├── valid.txt
    └── invalid.txt