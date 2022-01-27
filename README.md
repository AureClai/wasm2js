# Web Assembly bytes to JavaScript variable converter

Converts an Web Assembly file to a Javascript variable `data` in raw hexadecimal. Actually clone every byte from an input file to an output `data` Javascript variable with plain hexadecimal representation.

## Requirement

- Go v1.17

## Installation

1. Clone
```
git clone https://github.com/AureClai/wasm2js
```

2. Build
```
cd wasm2js
go build .
```

## Usage

### Without output path

```
wasm2js_executable <papth/to/input_file.wasm>
```

### With output path

```
wasm2js_executable papth/to/input_file.wasm -o path/to/output_file.js
```

## To do
- Testing