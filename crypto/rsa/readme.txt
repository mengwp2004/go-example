
#参考以下
http://blog.studygolang.com/2013/01/go%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86%E4%B9%8Brsa/


#
1）创建私钥：
openssl genrsa -out private.pem 1024 //密钥长度，1024觉得不够安全的话可以用2048，但是代价也相应增大
2）创建公钥：
openssl rsa -in private.pem -pubout -out public.pem



