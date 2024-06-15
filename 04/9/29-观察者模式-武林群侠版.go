// @Author huzejun 2024/6/16 2:04:00
package main

import "fmt"

type Event struct {
	Noti    Notifier //被知晓的通知者
	One     Listener //事件的主动发出者
	Another Listener //事件的被动接收者
	Msg     string   //具体的事务描述
}

// ----- 抽象层 -------
type Listener interface {
	//当同伴被揍了，该怎么办
	OnFriendBeFight(event *Event)

	//返回一个英雄的帮派 加 名称
	Title() string

	GetName() string
	GetParty() string
}

type Notifier interface {
	//添加观察者
	AddListener(listener Listener)
	//删除观察者
	RemoveListener(listener Listener)
	//通知广播
	Notify(event *Event)
}

// ----- 实现层 -------
// 英雄（观察者）
type Hero struct {
	Name  string //姓名
	Party string //所属的帮派
}

func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) GetParty() string {
	return hero.Party
}

func (hero *Hero) OnFriendBeFight(event *Event) {
	// 判断是否为当事人，忽略消息
	if hero.Name == event.One.GetName() || hero.Name == event.Another.GetName() {
		return
	}

	// 本帮派的同伴将其他帮派的揍了，要拍手叫好！！！
	if hero.Party == event.One.GetParty() {
		fmt.Println(hero.Title(), "得知消息， 拍手叫好！！！")
	}

	// 本帮派同伴被其他帮派的揍了，要“维护公平正义” 报仇反击 ！！！
	if hero.Party == event.Another.GetParty() {
		fmt.Println(hero.Title(), ", 得知消息，发起报仇反击！！！")
		hero.Fight(event.One, event.Noti)
		return
	}

}

func (hero *Hero) Title() string {
	return fmt.Sprintf("[%s]%s", hero.Party, hero.Name)
}

// 英雄主动攻击别人的方法
func (hero *Hero) Fight(another Listener, baixiao Notifier) {
	//1.生成一个武林事件
	event := new(Event)
	event.Msg = fmt.Sprintf("%s 将 %s 揍了...", hero.Title(), another.Title())
	event.Noti = baixiao
	event.One = hero
	event.Another = another

	//2.让百晓生知道，并且进行广播
	baixiao.Notify(event)
}

// 百晓生（通知者）
type Baixiao struct {
	heroList []Listener
}

// 添加观察者
func (b *Baixiao) AddListener(listener Listener) {
	b.heroList = append(b.heroList, listener)
}

// 删除观察者
func (b *Baixiao) RemoveListener(listener Listener) {
	for index, l := range b.heroList {
		if listener == l {
			b.heroList = append(b.heroList[:index], b.heroList[index+1:]...)
		}
	}
}

// 通知广播
func (b *Baixiao) Notify(event *Event) {
	fmt.Println("[世界消息] 百晓生广播消息：", event.Msg)

	//通知江湖大侠
	for _, listener := range b.heroList {
		listener.OnFriendBeFight(event)
	}
}

// ------ 业务逻辑层 ------
func main() {

	hero1 := Hero{
		"黄蓉",
		"丐帮",
	}

	hero2 := Hero{
		"洪七公",
		"丐帮",
	}

	hero3 := Hero{
		"乔峰",
		"丐帮",
	}

	hero4 := Hero{
		"张无忌",
		"明教",
	}

	hero5 := Hero{
		"灭绝师太",
		"明教",
	}

	hero6 := Hero{
		"金毛狮王-谢逊",
		"明教",
	}

	baixiao := Baixiao{}

	baixiao.AddListener(&hero1)
	baixiao.AddListener(&hero2)
	baixiao.AddListener(&hero3)
	baixiao.AddListener(&hero4)
	baixiao.AddListener(&hero5)
	baixiao.AddListener(&hero6)

	fmt.Println("武林一片平静.....")

	hero1.Fight(&hero4, &baixiao)

}
