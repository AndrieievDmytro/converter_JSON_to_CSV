# converter_JSON_to_CSV
Custom converter from JSON to CSV written in Golang

Run command "go build ." to build the program.

In order to run program use command line with the next arguments :

-p string

        Path to file. Format : input/file_name.json/csv

-t string

        Type of a file to convert(json/csv). 
        
-f string

        Name of the file to convert

Example

        Example: .\parcer -t json -p input/sessions.json
