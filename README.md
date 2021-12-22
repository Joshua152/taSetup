# About TA Setup

Program to help with grading programs for Java

The program automatically puts each file in a folder with the name of the student and renames the file to get rid of the "(#)" within the file name

## Setup

If getting message about insufficient permissions run this command inside folder with the taSetup file:

On mac:
`chmod +x taSetup-macIntel` OR `chmod +x taSetup-m1`

On windows:
`icacls taSetup.exe /grant Users:m`

## Usage

Run file in the directory with the loose programs

## Output

Should look something like this:

```
1| ~prog (10).java | ~name/prog.java
...
```


Tested by George Gino
