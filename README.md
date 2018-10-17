# vue-pd-build-test-with-gorice
用go-rice快速查看、发布webpack打包的产品。

## 说明
本仓库只有一个代码文件。
此项目衍生于我在读baidu-pcsgo代码的时候。go-rice可用于将静态web文件（html、js、css等）打包进单一二进制可执行文件，比较适合开源发布及交作业（逃）。

vue项目在build之后的文件一般不能直接通过文件形式在浏览器查看，一般是通过将代码转移到nginx目录下进行部署然后查看。

此代码可以直接构建带有html等静态web文件的超小型没用server。

然后，说白了，本项目其实没什么用，我自己应该也不会用，我只是想知道go-rice这东西怎么用而已。

## 依赖
go-rice。

## 用法
将dist目录替换为自己的。

## 协议
MIT.