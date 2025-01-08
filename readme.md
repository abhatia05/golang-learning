# golang-learning
Repo to learn golang

## 02 - Echo server
This folder has a basic echo server build using golang and a package based upon that. 
It also has docker file in it. 

```
git clone git@github.com:abhatia05/golang-learning.git
cd 02-echoServer
docker build . -t echo-server
docker run -p 1323:1323 echo-server
```
