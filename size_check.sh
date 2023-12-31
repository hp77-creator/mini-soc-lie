#! /bin/sh





compile_and_measure() {
  file=$1
  flags=$2
  filename=$(basename "$file" .go)
  result_file_name="result"
  extension=".txt"
  result="${result_file_name}-${filename}${extension}"
  echo "Compiling $file with flags: $flags"

  go build $flags "$file"


  size=$(stat -f %z "$filename")


  echo "File: $filename with flags $flags: $size bytes" >> "$result"


}

go_file=$1

# Compile and measure with different flags
compile_and_measure "$go_file" "-ldflags=-w"
compile_and_measure "$go_file" "-ldflags=-w -ldflags=-s"
compile_and_measure "$go_file" "-ldflags=-w -ldflags=-s -gcflags=all=-l"
compile_and_measure "$go_file" "-ldflags=-w -ldflags=-s -gcflags=all=-l -gcflags=all=-B"
compile_and_measure "$go_file" "-ldflags=-w -ldflags=-s -gcflags=all=-l -gcflags=all=-B -gcflags=all=-wb=false"

# Add more compile_and_measure lines with different flags as needed

# Display the results
cat "$result"
if [ -e "$result" ]
then
  rm "$result"
fi

rm "$filename"