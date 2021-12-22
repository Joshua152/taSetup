# About TA Setup

Program to help with grading programs for Java

The program automatically puts each file in a folder with the name of the student and renames the file to get rid of the "(#)" within the file name

## Download

Click on the latest release on the releases panel (should be on the right) and download the correct binary for your machine

Windows: taSetup.exe
Mac (with Intel): taSetup-macIntel
Mac (with Apple Silicon): taSetup-m1

## Setup

**If getting message about insufficient permissions run this command inside folder with the taSetup file:**

On mac:
`chmod +x taSetup-macIntel` OR `chmod +x taSetup-m1`

On windows:
`icacls taSetup.exe /grant Users:m`

**If getting popup about unknown developer on mac:**

Open via file explorer GUI (right click open -> open anyway) then run via terminal again

## Usage

1) Download progs from email
2) Unzip
3) Open terminal to that folder location
4) run `taSetup` there (ex if taSetup is in the same folder `./taSetup` OR `./taSetup-m1` OR `./taSetup-macIntel`; if `taSetup` is not in the same folder, write out the path to it)

If you want to just be able to call it from anywhere without having to write out the entire path location, add the program to the system PATH:

Windows:
1) Open start menu, search for **Edit the system environment variables**
2) Click **Environment Variables**
3) Click on **PATH** under **User variables**
4) Click **Edit** 
5) Click **New**
6) Enter the location to the `taSetup` program

IDK how to do it for mac (feel free to submit a pull request updating the README if you do know)

## Output

Should look something like this:

```
1| ~prog (10).java | ~name/prog.java
...
```


Tested by George Gino
