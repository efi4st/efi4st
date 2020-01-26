import sys
import os

rootpath = str(sys.argv[1])
result = ""
print("# InitSystemAnalysis")
print(" --> SystemD ?")
print("  ------> Exists? /etc/systemd/:"+str(os.path.isdir(rootpath+"etc/systemd")))
print("  ------> Exists? /etc/systemd/system/:"+str(os.path.isdir(rootpath+"etc/systemd/system")))
print("\n")
print(" --> SysVInit ?")
print("  ------> Exists? /etc/init.d/:"+str(os.path.isdir(rootpath+"etc/init.d")))
print("  ------> Exists? /etc/inittab:"+str(os.path.isfile(rootpath+"etc/inittab")))

if(os.path.isdir(rootpath+"etc/systemd/system")):
    print("\n")
    print("_____________ systemd services _________________")
    listOfFiles = os.listdir(rootpath+"etc/systemd/system/")
    for entry in listOfFiles:
        print (entry)
    print("____________________________________________")

if(os.path.isdir(rootpath+"etc/init.d")):
    print("\n")
    print("_____________ init.d scripts _________________")
    listOfFiles = os.listdir(rootpath+"etc/init.d/")
    for entry in listOfFiles:
        print ("+"+"/etc/init.d/"+entry)
    print("____________________________________________")

print("\n")
print("_____________ rc.d infrastructure _________________")
listOfFiles = os.listdir(rootpath+"etc/")
for entry in listOfFiles:
    if(entry.startswith('rc') and entry.endswith('.d')):
        print (entry+":")
        listOfFiles2 = os.listdir(rootpath+"etc/"+entry+"/")
        for entry2 in listOfFiles2:
            print("  --> "+entry2+"   "+"/etc/"+entry+"/"+entry2)
print("____________________________________________")

if(os.path.isfile(rootpath+"etc/bash.bashrc")):
    print("\n")
    print("_____________ /etc/inittab _________________")
    f=open(rootpath+"etc/inittab", "r")
    content=f.read()
    print(content)
    print("____________________________________________")
