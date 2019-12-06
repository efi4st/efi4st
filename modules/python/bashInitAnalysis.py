import sys
import os

rootpath = str(sys.argv[1])
result = ""
print("# BashInitAnalysis")

print("Exists? /etc/bash.bashrc:"+str(os.path.isfile(rootpath+"etc/bash.bashrc")))

if(os.path.isfile(rootpath+"etc/bash.bashrc")):
    print("\n")
    print("_____________ /etc/bash.bashrc _________________")
    f=open(rootpath+"etc/bash.bashrc", "r")
    content=f.read()
    print(content)
    print("____________________________________________")
