import sys
import requests

result = ""
for line in sys.stdin:
    result = result + line

project = str(sys.argv[1])
r = requests.post('http://127.0.0.1:8144/testResults/addResultSet/'+project, json={'result':result, 'source': sys.argv[2]})
print(r.text)