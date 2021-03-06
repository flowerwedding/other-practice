package main
//角色有血量、蓝量、攻击力、技能四个属性
//A角色：血量 100 蓝量 0 攻击力10（每次攻击减少对方10血量，初始蓝量为0）
//​              技能 无消耗、每次攻击有50%概率触发连击（又一次攻击），每回合理论上可以无限连击
//B角色：血量 300 蓝量 50 攻击力 20 （每次攻击回复 10 蓝量 ，初始蓝量为0）
//​              技能  蓝量达到50自动触发、降低对方10%当前攻击力
//回合制游戏，不考虑攻击速度
//比试十次，分别由每回合都是A先手5次，和每回合都是B先手5次组成
//请你计算出他们各自的胜率

//因为A血量是100，B每回合的攻击是20
//所以如果前五回合B不死亡则A必死亡
//故B的技能（A的攻击变为5）和改变攻击顺序不用考虑
//即问题等价于B能撑过第五回合的概率
//1.若B在第一回合死亡，即A触发了29次连击，概率（1/2)^29
//2.若B在第二回合死亡，即A总共触发了28次连击，共29种可能，概率29*（1/2）^28
//3.若B在第三回合死亡，即A总共触发了27次连击，共28*27种可能，概率28*27*（1/2)^27
//4.若B在第四回合死亡，即A总共触发了26次连击，共27*26*25种可能，概率27*26*25*（1/2)^26
//5.若B在第五回合死亡，即A总共触发了25次连击，共26*25*24*23种可能，概率26*25*24*23*（1/2)^25
//综上所述，A获胜的概率为p=（1/2)^29+29*（1/2）^28+28*27*（1/2)^27+27*26*25*（1/2)^26+26*25*24*23*（1/2)^25
//         B获胜的概率为1-p

import "fmt"

func main(){
	fmt.Println((1/2)^29+29*(1/2)^28+28*27*(1/2)^27+27*26*25*(1/2)^26+26*25*24*23*(1/2)^25)
}