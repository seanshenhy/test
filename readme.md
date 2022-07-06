# 简单编程题
## 问题 1：一个.go文件的顶级声明包含：

```
1.	 package，声明包名
2.	 import，导入其他包
3.	 type，定义类型
4.	 var / const，定义变量 / 常量
5.	 func，定义函数 / 方法
```
1.	 这些声明是否可以省略
2.	 这些声明是否有严格的先后顺序
3.	 其他你知道的，从这些声明衍生出的知识点，最多填写5条。例如：“使用‘_’导入一个包，则只会执行该包中的init()函数，不会导入其他任何内容”。

```
回答：
1 这些申明只有package是必须申明的，其他没用到的可以不需要出现在包文件里。
2 这些申明package需要严格放在最前面，import部分严格放在package申明后，且变量常量/type/函数定义前面，type定义和var定义以及func定义没有严格顺序，可以任意放置，但通常为了规范会按照上面1至5的顺序来放置。

3 
1）递归引入包中有init函数，会优先执行最里面的init函数，执行依次由内到外。
2）导入包的init函数会优先main函数执行。
3）若包里面有定义全局变量表达式，则全局变量表达式优先init函数执行。
4）首字母大写变量/常量/函数才能在其他包使用，小写的变量/常量/函数只能在包内部使用。
```


## 问题 2：使用代码举一个数组和切片声明的例子。
这里以int类型为例
数组
```go
var s [10]int
var s [10]int = [10]int{}
s := [10]int{}
```
切片
```go
var s []int
var s []int = []int{}
s := []int{}
```

## 问题 3：使用代码举一个匿名结构的例子。
```go
s := struct{
    name string
    age int
}{
    name:"zhangsan",
    age: 12,
}

```

## 问题 4：举一个defer语句延迟函数调用的例子。
读取文件内容
```go
func deferExample() {
    file, err := os.Open("./test.txt")
    if err != nil {
        fmt.Println("err = ", err)
        return
    }
    defer file.Close()

    rd := bufio.NewReader(file)
    for {
        buf, err := rd.ReadString('\n')
        if err == io.EOF {
            fmt.Printf("%s\n", buf)
            break
        } else if err != nil {
            fmt.Println("err = ", err)
        }
        fmt.Print(buf)
    }
}
```

## 问题 5：编写一个 Golang 程序，声明一个字符串变量，打印变量的地址，声明另一个 int 变量，以及指向它的指针。

```go
var s string = "this is a test string var"
var sp *string
sp = &s
fmt.Printf("%p %p", &s, sp)
```

## 问题 6：创建一个定义命名类型和该类型的方法（接收器函数）的 Go 程序。
```go
type student struct {
    Name string
    Class int
}
func (s *student)getName()string{
    return s.Name
}
```

## 问题 7：编写一个使用 goroutine 和 channel 的简单 Golang 程序。
```go
ch := make(chan int)
go func() {
    ch <- 100
}()
<-ch
fmt.Println("finish")
```

# 二、编程题：以下测试题二选一
```
1. 会员管理（您也可以选择以写一个区块链自动化合约方式来表达这个业务需求）
某电商，计划在其公众号上实现会员管理。
具体需求：
（1）老会员可以邀请新会员加入，形成上下级关系。每一个会员可以有多个下级，但是只有一个上级。
（2）每个会员，都有多个消费记录，记录其消费明细。
（3）每个月底，统计每个会员及其下级会员的消费总额
请完成下列任务：
（1）设计这个会员管理功能所需的数据库表, 并使用go语言struct描述
（2）编写计算某个会员，在某个月消费总额：其自身消费和所有下级会员的消费的总和
```

## 表设计
### 表结构
```sql
CREATE TABLE `member` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `parent_id` int(11) NOT NULL,
  `leaf_node` tinyint(1) DEFAULT 1,
  PRIMARY KEY (`id`),
  KEY `pidx` (`parent_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='会员表';

CREATE TABLE `month_stat` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `member_id` int(10) NOT NULL,
  `total` int(11) NOT NULL,
  `mon_date` char(8) NOT NULL,
  `created_at` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='会员每月月统计表';

CREATE TABLE `record` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `member_id` int(10) NOT NULL,
  `cost` int(11) NOT NULL,
  `created_at` int(11) NOT NULL,
  `mon_date` char(8) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='记录表';

```
### 数据
```sql
INSERT INTO `member` VALUES (1, 'zhangsan', 0, 0);
INSERT INTO `member` VALUES (2, 'lisi', 1, 0);
INSERT INTO `member` VALUES (3, 'wangwu', 1, 1);
INSERT INTO `member` VALUES (4, 'zhaoliu', 2, 0);
INSERT INTO `member` VALUES (5, 'zhuqi', 4, 1);
INSERT INTO `member` VALUES (6, 'jinba', 4, 1);


INSERT INTO `record` VALUES (1, 1, 100, 16000000, '2022-07');
INSERT INTO `record` VALUES (2, 2, 500, 16000000, '2022-07');
INSERT INTO `record` VALUES (3, 1, 200, 16000000, '2022-07');

```
使用方式

```sql
使用montask/main.go生成月统计记录，该记录是已经计算总消费。
使用member/main.go直接查month_stat表的月总消费总消费即可。
```







