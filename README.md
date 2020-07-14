### 以应用为中心获取应用使用到的kubernetes资源对象
#### pod
#### service
#### configMap
#### secret
#### pvc
#### pv
#### stroageClass
#### serviceAccount
#### ...

### 改进点
- 获取kubernetes客户端依赖./kube/config文件
- 有状态应用和无状态应用的逻辑处理可以分装为一个方法,减少代码冗余
- 没有制作docker镜像

### 测试结果

### 已知问题


