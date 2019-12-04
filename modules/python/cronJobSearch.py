import sys
import os

rootpath = str(sys.argv[1])
result = ""
print("# CronJobSearch")

print("Exists? /etc/crontab:"+str(os.path.isfile(rootpath+"etc/crontab")))
print("Exists? /var/spool/cron/:"+str(os.path.isdir(rootpath+"var/spool/cron")))
print("Exists? /etc/cron.d/:"+str(os.path.isdir(rootpath+"etc/cron.d")))

if(os.path.isfile(rootpath+"etc/crontab")):
    print("\n")
    print("_____________ /etc/crontab _________________")
    f=open(rootpath+"etc/crontab", "r")
    content=f.read()
    print(content)
    print("____________________________________________")

if(os.path.isdir(rootpath+"var/spool/cron")):
    for filename in os.listdir(rootpath+"var/spool/cron/"):
        print("\n")
        print("_____________ /var/spool/cron/"+filename+"_________________")
        f=open(rootpath+"var/spool/cron/"+filename, "r")
        content=f.read()
        print(content)
        print("____________________________________________")

if(os.path.isdir(rootpath+"/etc/cron.d/")):
    for filename in os.listdir(rootpath+"etc/cron.d/"):
        print("\n")
        print("_____________ /etc/cron.d/"+filename+"_________________")
        f=open(rootpath+"etc/cron.d/"+filename, "r")
        content=f.read()
        print(content)
        print("____________________________________________")