# Go File Compressor

This Repo provides a simple Go program to compress files using gzip.

## How to Use

1. **Enter Source File:** When prompted, type the path to the file you want to compress. For example: `my_file.txt`
2. **Enter Destination File:** When prompted, type the path and name you want to use for the compressed file. For example: `compressed_file`.  The program will automatically add the `.gz` extension to the filename.
3. **Run the program:** The program will compress the source file and save it as a `.gz` file in the same directory.

## Example Usage
Enter source file: my_file.txt
Enter destination file: compressed_file
File compressed successfully

This will create a file named `compressed_file.gz` in the same directory as `my_file.txt`.

## Notes

- The program currently only supports compressing files in the same directory. 
- You may need to install the `gzip` package if it's not already included.

## Contributing

Feel free to fork this Repo and add features or improvements.