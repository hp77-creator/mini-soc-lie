#! /bin/sh

compile_and_measure() {
  file=$1
  flags=$2

  echo "Compiling $file with flags: $flags"

  go build $flags "$file"

  filename=$(basename "$file" .go)
  size=$(stat -f %z "$filename")
  result_file_name="result"
  extension=".txt"
  result="${result_file_name}-${filename}${extension}"

  echo "File size with flags $flags: $size bytes" >> "$result"


}

go_file=$1

# Compile and measure with different flags
compile_and_measure "$go_file" "-ldflags='-w -s'"
compile_and_measure "$go_file" "-ldflags='-w -s' -gcflags=all=-l"

# Add more compile_and_measure lines with different flags as needed

# Display the results
cat results.txt