## 参考链接
https://mp.weixin.qq.com/s?__biz=MzkxMjQzMjA0OQ==&mid=2247483677&idx=1&sn=d1c0e52b1fd31932867ec9b1d00f4ec2&poc_token=HBN3B2mju3G-A44GR7wO6PO_rSldk-dzycDvgOrZ


## 前言
context 是 golang 中的经典工具，主要在异步场景中用于实现并发协调以及对 goroutine 的生命周期控制。除此之外，context 还兼有一定的数据存储能力. 本着知其然知其所以然的精神，本文和大家一起深入 context 源码一探究竟，较为细节地对其实现原理进行梳理.


## 核心数据结构
1.1 context.Context
