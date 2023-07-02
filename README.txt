#slon

slon - CLI utility for analyze system storage

slon scanning files in given directory and output information about storage usage

#Instalation:
1. Ensure you have Go 1.20 or above installed
2. Clone this repository to your local machine
3. Navigate to the project directory
4. Run following command:
go build memory-cli-utility/cmd/slon
5. After building, you need to add the path to the executable file to the PATH variable
   in order to use it from the terminal without specifying the full path each time.

#Usage:

to start using slon you have to type following command to terminal:
slon <flags> <path-to-directory>

The flags are:
s - sorting files and directories
p - output percent of parent directory size
e - write results of analyse into file (example of usage: -e=log.txt)
n - set maximal nesting level of directories to output