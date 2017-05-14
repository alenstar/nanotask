# nanoweb
nano web framework
一个小型轻量级的RESETFul框架
# 测试
添加一个对象：

curl -X POST -d '{"Score":1337,"PlayerName":"Sean Plott"}' http://127.0.0.1:8080/object

返回一个相应的objectID:hello

查询一个对象

curl -X GET http://127.0.0.1:8888/object/hello

查询全部的对象

curl -X GET http://127.0.0.1:8888/object

更新一个对象

curl -X PUT -d '{"Score":10000}'http://127.0.0.1:8888/object/hello

删除一个对象

curl -X DELETE http://127.0.0.1:8888/object/hello

