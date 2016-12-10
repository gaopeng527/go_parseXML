// ParseXML project doc.go

/*
ParseXML document
*/
package main

/*
解析XML到struct的时候遵循如下的规则：

如果struct的一个字段是string或者[]byte类型且它的tag含有",innerxml"，Unmarshal将会将此字段所对应的元素内所有内嵌的原始xml累加到此字段上，如上面例子Description定义。最后的输出是

<server>
    <serverName>Shanghai_VPN</serverName>
    <serverIP>127.0.0.1</serverIP>
</server>
<server>
    <serverName>Beijing_VPN</serverName>
    <serverIP>127.0.0.2</serverIP>
</server>
如果struct中有一个叫做XMLName，且类型为xml.Name字段，那么在解析的时候就会保存这个element的名字到该字段,如上面例子中的servers。

如果某个struct字段的tag定义中含有XML结构中element的名称，那么解析的时候就会把相应的element值赋值给该字段，如上servername和serverip定义。
如果某个struct字段的tag定义了中含有",attr"，那么解析的时候就会将该结构所对应的element的与字段同名的属性的值赋值给该字段，如上version定义。
如果某个struct字段的tag定义 型如"a>b>c",则解析的时候，会将xml结构a下面的b下面的c元素的值赋值给该字段。
如果某个struct字段的tag定义了"-",那么不会为该字段解析匹配任何xml数据。
如果struct字段后面的tag定义了",any"，如果他的子元素在不满足其他的规则的时候就会匹配到这个字段。
如果某个XML元素包含一条或者多条注释，那么这些注释将被累加到第一个tag含有",comments"的字段上，这个字段的类型可能是[]byte或string,如果没有这样的字段存在，那么注释将会被抛弃。
上面详细讲述了如何定义struct的tag。 只要设置对了tag，那么XML解析就如上面示例般简单，tag和XML的element是一一对应的关系，如上所示，我们还可以通过slice来表示多个同级元素。

注意： 为了正确解析，go语言的xml包要求struct定义中的所有字段必须是可导出的（即首字母大写）
*/

/*
和我们之前定义的文件的格式一模一样，之所以会有os.Stdout.Write([]byte(xml.Header)) 这句代码的出现，是因为xml.MarshalIndent或者xml.Marshal输出的信息都是不带XML头的，为了生成正确的xml文件，我们使用了xml包预定义的Header变量。

我们看到Marshal函数接收的参数v是interface{}类型的，即它可以接受任意类型的参数，那么xml包，根据什么规则来生成相应的XML文件呢？

如果v是 array或者slice，那么输出每一个元素，类似value
如果v是指针，那么会Marshal指针指向的内容，如果指针为空，什么都不输出
如果v是interface，那么就处理interface所包含的数据
如果v是其他数据类型，就会输出这个数据类型所拥有的字段信息
生成的XML文件中的element的名字又是根据什么决定的呢？元素名按照如下优先级从struct中获取：

如果v是struct，XMLName的tag中定义的名称
类型为xml.Name的名叫XMLName的字段的值
通过struct中字段的tag来获取
通过struct的字段名用来获取
marshall的类型名称
我们应如何设置struct 中字段的tag信息以控制最终xml文件的生成呢？

XMLName不会被输出
tag中含有"-"的字段不会输出
tag中含有"name,attr"，会以name作为属性名，字段值作为值输出为这个XML元素的属性，如上version字段所描述
tag中含有",attr"，会以这个struct的字段名作为属性名输出为XML元素的属性，类似上一条，只是这个name默认是字段名了。
tag中含有",chardata"，输出为xml的 character data而非element。
tag中含有",innerxml"，将会被原样输出，而不会进行常规的编码过程
tag中含有",comment"，将被当作xml注释来输出，而不会进行常规的编码过程，字段值中不能含有"--"字符串
tag中含有"omitempty",如果该字段的值为空值那么该字段就不会被输出到XML，空值包括：false、0、nil指针或nil接口，任何长度为0的array, slice, map或者string
tag中含有"a>b>c"，那么就会循环输出三个元素a包含b，b包含c，例如如下代码就会输出

FirstName string   `xml:"name>first"`
LastName  string   `xml:"name>last"`

<name>
<first>Asta</first>
<last>Xie</last>
</name>
上面我们介绍了如何使用Go语言的xml包来编/解码XML文件，重要的一点是对XML的所有操作都是通过struct tag来实现的，所以学会对struct tag的运用变得非常重要，在文章中我们简要的列举了如何定义tag。
*/
