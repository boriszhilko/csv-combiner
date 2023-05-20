# CSV Combiner

The CSV Combiner is a command-line tool written in Go that combines names and descriptions for a company from two CSV files and writes the combined data into a new output CSV file.
It is assumed that the file can be loaded entirely into memory. For large files, see the [Extension for parallel processing of the large input file](#extension-for-parallel-processing-of-the-large-input-file) section.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/boriszhilko/csv-combiner.git
   cd csv-combiner
   ```

2. Build the project:

   ```bash
   go build ./cmd/csv-combiner
   ```

## Usage

To run the CSV Combiner tool, use the following command:

```bash
./csv-combiner names.csv descriptions.csv combined.csv
```

Replace `names.csv` with the path to the first input CSV file containing company names, `descriptions.csv` with the path to the second input CSV file containing company descriptions, and `combined.csv` with the desired output CSV file path where the combined data will be written.

The tool reads the names and descriptions from the input CSV files, combines them for each company, and writes the combined data into the output CSV file. The input CSV files should be properly formatted with a header row and data rows. The output CSV file will be created if it doesn't exist and overwritten if it already exists.

## Testing

To run the tests, use the following command:

   ```bash
   go test ./...
   ```

## Extension

The CSV Combiner tool is designed for easy extensibility, allowing you to add new input and output sources. By implementing the `io.Reader` interface for new input sources and the `company.Writer` interface for new output sources, you can seamlessly integrate them with the tool.

To add a new input source, create a parser that implements the `io.Reader` interface to read data from the desired source. Modify the `parseNames` and `parseDescriptions` functions in `main.go` to use the new parser.

For a new output source, implement the `company.Writer` interface with a writer that handles writing data to the desired destination. Update the `combineAndWrite` function in `main.go` to use the new writer.

This approach allows you to integrate the CSV Combiner tool with various sources, such as databases or APIs. The use of interfaces makes it easy to swap or add new implementations without affecting the core functionality of the tool.

## Extension for parallel processing of the large input file

To enable the CSV Combiner tool to process large input files that cannot be loaded entirely into memory, you would need to extend the application to implement the following approach:

1. Modify the `parseNames` and `parseDescriptions` functions in `main.go` to accept a buffered channel as an additional parameter. These functions will no longer read directly from the CSV reader but instead receive data from the buffered channel.

2. Create a new function, let's call it `streamData`, that takes the filename and a buffered channel as input. This function should be responsible for reading the input file in chunks and sending the data to the buffered channel.

3. Start multiple processor goroutines that read from the buffered channel, perform the necessary processing (such as combining names and descriptions), and write the processed data to the output source. Each processor goroutine should operate independently on a chunk of data received from the buffered channel.

4. Modify the `combineAndWrite` function to accept the buffered channel as an additional parameter. Instead of calling the `parseNames` and `parseDescriptions` functions directly, this function should start the `streamData` function in a separate goroutine, passing the buffered channel as an argument. Additionally, the `combineAndWrite` function should wait for the processor goroutines to complete before returning.

By implementing this approach, the CSV Combiner tool will be able to process large input files efficiently in a streaming fashion. The use of buffered channels and goroutines allows for concurrent processing of data, minimizing memory usage while maximizing processing throughput.

Note that additional error handling and synchronization mechanisms may be required to ensure proper coordination between the reader, processor goroutines, and output writer.
