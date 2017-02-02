# base64img
Encodes/Decodes images to/from base64 format

## Usage
```
base64img action filename

action can be either encode or decode
filename, in case of encode, is a jpg/png image file
filename, in case of decode, is a text file containing "data:image/png;base64,..."

Output is written to stdout
```

## Example 1 : Decoding
```
$ base64img decode examples/redsquare.txt > examples/redsquare.png
```

## Example 2: Encoding
```
$ base64img encode examples/redsquare.png
```

Output:
```
data:image/png;base64,iVBORw0K...
```
