import sys
import os
import json
import re

def findPathInLine(line, relPath):
        #line is path absolute
        cwd = os.getcwd()
        if(os.path.exists(cwd + "/../../working/filesystem"+line)):
                return line
        #line is path relative
        elif((os.path.exists(cwd + "/../../working/filesystem" + relPath + line)) or (os.path.exists(cwd + "/../../working/filesystem" + relPath + "/" + line))):
                return line
        else:
                parts = re.split(' \"=', line)
                for part in parts:
                        words = part.split()
                        for word in words:
                                if(len(word) > 0):
                                        if(word[0]=='/'):
                                                if(os.path.exists(cwd + "/../../working/filesystem"+word)):
                                                        return word
                                                #line is path relative
                                                elif((os.path.exists(cwd + "/../../working/filesystem" + relPath + word)) or (os.path.exists(cwd + "/../../working/filesystem" + relPath + "/" + word))):
                                                        return word
                return ""    


print("# AppChainAnalysis")
input = ""
pathOrgApp = str(sys.argv[1])
relPath = pathOrgApp.rsplit('/', 1)[0]
firstLevelPathList = []
for line in sys.stdin:
        appPath = findPathInLine(line.strip(), relPath)
        if(len(appPath) > 3):
                firstLevelPathList.append(appPath)

result = json.dumps(firstLevelPathList)
print(result)


