# Bamboo
采用分段存储的分布式数据链

![Image text](https://github.com/luonannet/Bamboo/blob/master/docs/design.fw.png)

数据被切分成若干块进行分段存储，整体使用
如果把整个数据链的数据当成一根竹子的话，那么每个人只需要存储一个竹节的数据即可。当本机需要使用其他竹节的数据的时候会通过路由规则找到。

为了避免数据的损失和恶意串改，每根竹子都会有一万份备份。他们同样分散在竹节之中。这样形成一个竹林。竹林包括一万根竹子，每根竹子都是一个完整的数据链数据。但每个人只需要保存一个竹节即可。
  

<iframe src="https://www.163.com" width="700px" height="500px" frameborder="0" scrolling="no"  frameborder=0  
 allowfullscreen>    </iframe>
 
