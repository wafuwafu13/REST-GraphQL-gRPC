# REST

http://localhost:1323/users  
```json
{"1":{"id":1,"name":"foo"},"2":{"id":2,"name":"bar"}}
```  

http://localhost:1323/user/1  
```json
{"id":1,"name":"foo"}
```

http://localhost:1323/user/2
```json
{"id":2,"name":"bar"}
```
  
http://localhost:1323/posts  
```json
{"1":{"id":1,"userId":1,"title":"title1","content":"content1"},"2":{"id":2,"userId":1,"title":"title2","content":"content2"},"3":{"id":3,"userId":2,"title":"title2","content":"content2"}}
```
  
http://localhost:1323/post/1  
```json
{"1":{"id":1,"userId":1,"title":"title1","content":"content1"},"2":{"id":2,"userId":1,"title":"title2","content":"content2"}}
```
  
http://localhost:1323/post/2  
```json
{"3":{"id":3,"userId":2,"title":"title2","content":"content2"}}
```
