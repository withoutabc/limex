package main

import "fmt"

type Metal struct {
	State        string //状态
	AtomicNumber int    //原子序数
	Color        string
	Density      string //密度
	MeltingPoint int    //熔点
	BoilingPoint int    //沸点
}

func main() {
	var Temperature = 26
	var iron Metal //初始化结构体
	iron.State = "solid"
	iron.AtomicNumber = 26
	iron.Color = "银白色"
	iron.Density = "7.86g／cm^3"
	iron.MeltingPoint = 1535
	iron.BoilingPoint = 2750.
	fmt.Printf("铁的原子序数为%d\n铁的颜色为%v\n铁的密度为%v\n铁的熔点为%d°C\n铁的沸点为%d°C\n",
		iron.AtomicNumber, iron.Color, iron.Density, iron.MeltingPoint, iron.BoilingPoint)
	fmt.Println("请输入要加热到的温度（单位省略）")
	fmt.Scan(&Temperature)
	iron.Burn(Temperature)                       //使用方法
	fmt.Printf("现在铁的状态为%v\n", iron.State) //打印目前铁的状态
	fmt.Printf("铁的温度为%d°C", Temperature)
}
func (iron *Metal) Burn(a int) { //建立方法
	if a < 1535 {
		iron.State = "固态"
	} else if a >= 1535 && a < 2750 {
		iron.State = "液态"
	} else {
		iron.State = "气态"
	}

}
