# 一、说明

本代码库使用 Go 语言以可运行的测试用例形式，演示 22 种常用设计模式的思想与实现。每个设计模式对应一个 `*_test.go`
文件，文件头部注释说明该模式的思想及其与示例的联系，测试函数即为可执行的使用示例。

# 二、设计模式

点击图片进入源代码文件。

## 2.1 创建型模式（Creational）

<div style="display:flex;flex-wrap:wrap;gap:12px;">

<a href="factory_method_test.go" title="工厂方法 Factory Method" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:247px;height:351px;background-image:url('./design_pattern.png');background-position:0px 0px;background-repeat:no-repeat;"></div>
  <p style="width:247px;margin:4px 0 0;font-size:14px;">把对象的创建集中在工厂方法中，客户端面向接口编程，实现创建与使用解耦。</p>
</a>

<a href="abstract_factory_test.go" title="抽象工厂 Abstract Factory" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:247px;height:351px;background-image:url('./design_pattern.png');background-position:-247px 0px;background-repeat:no-repeat;"></div>
  <p style="width:247px;margin:4px 0 0;font-size:14px;">创建一系列相互配套的相关产品（产品族），而无需指定其具体类。</p>
</a>

<a href="builder_test.go" title="生成器 Builder" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:247px;height:347px;background-image:url('./design_pattern.png');background-position:0px -351px;background-repeat:no-repeat;"></div>
  <p style="width:247px;margin:4px 0 0;font-size:14px;">分步骤构建复杂对象，同一构建流程搭配不同建造者即可产出不同表示。</p>
</a>

<a href="prototype_test.go" title="原型 Prototype" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:247px;height:349px;background-image:url('./design_pattern.png');background-position:-247px -351px;background-repeat:no-repeat;"></div>
  <p style="width:247px;margin:4px 0 0;font-size:14px;">通过克隆已有对象来创建新对象，使代码不依赖对象所属的具体类。</p>
</a>

<a href="singleton_test.go" title="单例 Singleton" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:247px;height:351px;background-image:url('./design_pattern.png');background-position:0px -698px;background-repeat:no-repeat;"></div>
  <p style="width:247px;margin:4px 0 0;font-size:14px;">保证一个类全局只有一个实例，并提供一个访问该实例的统一入口。</p>
</a>

</div>

## 2.2 结构型模式（Structural）

<div style="display:flex;flex-wrap:wrap;gap:12px;">

<a href="adapter_test.go" title="适配器 Adapter" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:243px;height:351px;background-image:url('./design_pattern.png');background-position:-535px 0px;background-repeat:no-repeat;"></div>
  <p style="width:243px;margin:4px 0 0;font-size:14px;">将不兼容的接口转换为目标接口，使原本无法协同工作的类可以合作。</p>
</a>

<a href="bridge_test.go" title="桥接 Bridge" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:247px;height:351px;background-image:url('./design_pattern.png');background-position:-778px 0px;background-repeat:no-repeat;"></div>
  <p style="width:247px;margin:4px 0 0;font-size:14px;">把抽象与实现拆分为两个独立层次并通过组合连接，避免多维度继承的类爆炸。</p>
</a>

<a href="composite_test.go" title="组合 Composite" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:243px;height:347px;background-image:url('./design_pattern.png');background-position:-535px -351px;background-repeat:no-repeat;"></div>
  <p style="width:243px;margin:4px 0 0;font-size:14px;">以统一接口把对象组合成树状结构，使客户端一致地对待单个对象与整体。</p>
</a>

<a href="decorator_test.go" title="装饰 Decorator" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:247px;height:347px;background-image:url('./design_pattern.png');background-position:-778px -351px;background-repeat:no-repeat;"></div>
  <p style="width:247px;margin:4px 0 0;font-size:14px;">用同接口的包装对象层层包裹原对象，在运行时动态地为其叠加新行为。</p>
</a>

<a href="facade_test.go" title="外观 Facade" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:243px;height:347px;background-image:url('./design_pattern.png');background-position:-535px -698px;background-repeat:no-repeat;"></div>
  <p style="width:243px;margin:4px 0 0;font-size:14px;">为复杂子系统提供一个简单的统一入口，隐藏内部细节与调用顺序。</p>
</a>

<a href="flyweight_test.go" title="享元 Flyweight" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:247px;height:351px;background-image:url('./design_pattern.png');background-position:-778px -698px;background-repeat:no-repeat;"></div>
  <p style="width:247px;margin:4px 0 0;font-size:14px;">分离内在状态与外在状态，通过工厂共享内在状态，在有限内存中载入更多对象。</p>
</a>

<a href="proxy_test.go" title="代理 Proxy" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:247px;height:351px;background-image:url('./design_pattern.png');background-position:-535px -1045px;background-repeat:no-repeat;"></div>
  <p style="width:247px;margin:4px 0 0;font-size:14px;">提供与真实服务同接口的替代品，在转发请求前后附加访问控制、缓存等处理。</p>
</a>

</div>

## 2.3 行为型模式（Behavioral）

<div style="display:flex;flex-wrap:wrap;gap:12px;">

<a href="chain_of_responsibility_test.go" title="责任链 Chain of Responsibility" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:246px;height:351px;background-image:url('./design_pattern.png');background-position:-1070px 0px;background-repeat:no-repeat;"></div>
  <p style="width:246px;margin:4px 0 0;font-size:14px;">让请求沿着处理者链依次传递，每个处理者决定处理请求或转发给下一个。</p>
</a>

<a href="command_test.go" title="命令 Command" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:244px;height:351px;background-image:url('./design_pattern.png');background-position:-1316px 0px;background-repeat:no-repeat;"></div>
  <p style="width:244px;margin:4px 0 0;font-size:14px;">把请求封装为对象，使调用者与接收者解耦，支持排队、替换与回滚。</p>
</a>

<a href="iterator_test.go" title="迭代器 Iterator" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:244px;height:351px;background-image:url('./design_pattern.png');background-position:-1560px 0px;background-repeat:no-repeat;"></div>
  <p style="width:244px;margin:4px 0 0;font-size:14px;">在不暴露集合内部结构的前提下，为遍历集合元素提供统一的接口。</p>
</a>

<a href="mediator_test.go" title="中介者 Mediator" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:247px;height:351px;background-image:url('./design_pattern.png');background-position:-1804px 0px;background-repeat:no-repeat;"></div>
  <p style="width:247px;margin:4px 0 0;font-size:14px;">用中介对象集中封装一组对象间的交互，把多对多依赖收敛为一对一。</p>
</a>

<a href="memento_test.go" title="备忘录 Memento" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:246px;height:347px;background-image:url('./design_pattern.png');background-position:-1070px -351px;background-repeat:no-repeat;"></div>
  <p style="width:246px;margin:4px 0 0;font-size:14px;">在不破坏封装的前提下，保存对象状态快照并在以后将其还原。</p>
</a>

<a href="observer_test.go" title="观察者 Observer" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:244px;height:347px;background-image:url('./design_pattern.png');background-position:-1316px -351px;background-repeat:no-repeat;"></div>
  <p style="width:244px;margin:4px 0 0;font-size:14px;">对象状态变化时自动通知所有订阅者，实现主题与观察者之间的松耦合。</p>
</a>

<a href="state_test.go" title="状态 State" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:244px;height:351px;background-image:url('./design_pattern.png');background-position:-1560px -351px;background-repeat:no-repeat;"></div>
  <p style="width:244px;margin:4px 0 0;font-size:14px;">把每种状态封装成独立的类，使对象随内部状态改变行为，消除大量条件分支。</p>
</a>

<a href="strategy_test.go" title="策略 Strategy" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:247px;height:351px;background-image:url('./design_pattern.png');background-position:-1804px -351px;background-repeat:no-repeat;"></div>
  <p style="width:247px;margin:4px 0 0;font-size:14px;">把算法封装成可互相替换的策略对象，使其能在运行时动态切换。</p>
</a>

<a href="template_method_test.go" title="模板方法 Template Method" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:246px;height:351px;background-image:url('./design_pattern.png');background-position:-1070px -698px;background-repeat:no-repeat;"></div>
  <p style="width:246px;margin:4px 0 0;font-size:14px;">在基类中定义算法骨架，允许子类重写特定步骤而不改变整体结构。</p>
</a>

<a href="visitor_test.go" title="访问者 Visitor" style="text-decoration:none;color:inherit;text-align:center;">
  <div style="width:247px;height:351px;background-image:url('./design_pattern.png');background-position:-1316px -698px;background-repeat:no-repeat;"></div>
  <p style="width:247px;margin:4px 0 0;font-size:14px;">把作用于对象结构的操作封装为访问者，新增操作无需修改各元素类。</p>
</a>

</div>
