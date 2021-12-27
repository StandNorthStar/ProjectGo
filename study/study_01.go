package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	/*
	1. 有 1、2、3、4 这四个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
	 */
	fmt.Println("1. ------------------------")
	totalcount := 0
	for i:=1;i<5;i++ {
		for j:=1;j<5;j++ {
			for k:=1;k<5;k++ {
				if i!=k && i!=j && k!=j {
					totalcount+=1
					//fmt.Println("第",totalcount,"方式",i,j,k)
				}
			}
		}
	}
	fmt.Println("共：",totalcount,"种不同")

	/*
	企业发放的奖金根据利润提成。利润(I)低于或等于 10 万元时，奖金可提成 10%；利润高于 10 万元，低于 20 万元，低于 10 万元的部分按 10% 提成，高于 10 万元的部分，可提成 7.5%。
	20 万到 40 万之间时，高于 20 万元的部分，可提成 5%；40 万到 60 万之间时高于 40 万元的部分，可提成 3%；60 万到 100 万之间时，高于 60 万元的部分，可提成 1.5%，高于 100 万元时，超过 100 万元的部分按 1% 提成。
	从键盘输入当月利润 I，求应发放奖金总数？
	 */
	// 方法一：
	var profit float64
	var bonus float64
	fmt.Println("2. ------------------------")
	fmt.Println("输入利润：")
	fmt.Scanf("%f \n", &profit)
	if profit <= 10 {
		bonus = profit * 0.1
	} else if (profit > 10 && profit <= 20){
		bonus = 10 * 0.1 + (profit - 10) * 0.075
	} else if (profit > 20 && profit <= 40) {
		bonus = (profit - 20)*0.05 + 10 * 0.075 + 10 * 0.1
	} else if (profit > 40 && profit <= 60) {
		bonus = (profit-40)*0.03 + 20*0.05 + 10 * 0.075 + 10 * 0.1
	} else if (profit > 60 && profit <= 100) {
		bonus = (profit - 60) *0.015 + 20 * 0.03 + 20*0.05 + 10 * 0.075 + 10 * 0.1
	} else if (profit > 100) {
		bonus = (profit - 100) * 0.01 + 40 * 0.015 + 20 * 0.03 + 20*0.05 + 10 * 0.075 + 10 * 0.1
	}
	fmt.Println("方法一结果 -- 利润：",profit,"万元", "奖金：",bonus,"万元")
	// 方法二：
	bonus = 0
	fmt.Print("方法二结果 -- 利润：",profit,"万元")
	switch  {
	case profit>100:
		bonus = (profit - 100) * 0.01
		profit = 100
		fallthrough
	case profit>60:
		bonus += (profit-60)*0.015
		profit=60
		fallthrough
	case profit>40:
		bonus += (profit-40)*0.03
		profit=40
		fallthrough
	case profit>20:
		bonus += (profit-20)*0.05
		profit=20
		fallthrough
	case profit>10:
		bonus += (profit-10)*0.075
		profit=10
		fallthrough
	default:
		bonus += profit * 0.1
	}
	fmt.Println(" 奖金：",bonus,"万元")

	/*
	3. 输入某年某月某日，判断这一天是这一年的第几天？
	 */
	var yy,m,d int
	var days int
	fmt.Println("3. ------------------------")
	fmt.Print("请输入年月日：")
	fmt.Scanf("%d%d%d\n",&yy,&m,&d)
	fmt.Printf("%d 年 %d 月 %d 日", yy, m, d)
	switch m {
	case 12:
		days += d
		d = 30   // 十一月会用到这个变量，所以是30
		fallthrough
	case 11:
		days +=d
		d=31     // 十月会用到此变量，所以是31
		fallthrough
	case 10:
		days +=d
		d=30
		fallthrough
	case 9:
		days +=d
		d=31
		fallthrough
	case 8:
		days +=d
		d=31
		fallthrough
	case 7:
		days +=d
		d=30
		fallthrough
	case 6:
		days +=d
		d=31
		fallthrough
	case 5:
		days +=d
		d=30
		fallthrough
	case 4:
		days+=d
		d=31
		fallthrough
	case 3:
		days+=d
		d=28
		if (yy%400 ==0 ) || (yy%4 ==0 && yy%100 == 0) {
			d+=1
		}
		fallthrough
	case 2:
		days+=d
		d=31
		fallthrough
	case 1:
		days += d
	}
	fmt.Printf("是今年的第%d天 \n",days)


	/*
	4. 输入三个 整数 x，y，z，请把这三个数由小到大输出。
	 */
	fmt.Println("4. ------------------------")
	var x,y,z int
	fmt.Print("请输入数字：")
	fmt.Scanf("%d%d%d\n",&x,&y,&z)
	if x > y {
		x, y = y, x
	}
	if x > z {
		x, z = z, x
	}
	if y > z {
		y, z = z, y
	}
	fmt.Printf("由小到大: %d < %d < %d \n", x,y,z)

	/*
	5. 输出 9*9 乘法口诀表。
	 */
	fmt.Println("5. ------------------------")
	// 横排
	for i:=1;i<10;i++ {
		for j:=i;j<10;j++ {
			sum := i*j
			fmt.Printf("%d * %d = %-5d",i,j,sum)
		}
		fmt.Println("")
	}
	// 竖排
	for i:=1;i<10;i++{
		for j:=1;j<=i;j++{
			sum := i*j
			fmt.Printf("%d * %d = %-5d",j,i,sum)
		}
		fmt.Println("")
	}

	/*
	6. 输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数。
	 */
	fmt.Println("6. ------------------------")
	var zm,kg,num,oth int
	fmt.Print("请输入一串字符：")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Sprintf(input)
	for _, ch := range input {
		switch {
		case ch >= 'A' && ch <= 'Z':
			zm++
		case ch >= 'a' && ch <= 'z':
			zm++
		case ch >= '0' && ch <= '9':
			num++
		case ch == ' ':
			kg++
		default:
			oth++
		}
	}
	fmt.Printf("字母：%d 数字：%d 空格:%d 其他:%d \n", zm,num,kg,oth)

	/*
	7. 猴子吃桃问题：猴子第一天摘下若干个桃子，当即吃了一半，还不瘾，又多吃了一个。第二天早上又将剩下的桃子吃掉一半，又多吃了一个。
	以后每天早上都吃了前一天剩下的一半零一个。到第 10 天早上想再吃时，见只剩下一个桃子了。求第一天共摘了多少。
	 */
	fmt.Println("7. ------------------------")
	var total int = 1
	for d:=9;d>=1;d-- {
		total = (total+1)*2
		fmt.Printf("d = %d , total=%d \n",d,total)
	}

	/*
	8. 利用递归方法求 5!。 5!=1*2*3*4*5
	 */
	fmt.Println("8. ------------------------")
	jc := matchJC(10)
	fmt.Printf("jc = %d \n", jc)

	/*
	9. 利用递归 函数 调用方式，将所输入的 5 个字符，以相反顺序打印出来。
	 */
	fmt.Println("9. ------------------------")
	fmt.Printf("请输入字符串：")
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	fmt.Printf("%s \n",input1)
	fmt.Printf("转换后：%s \n",Reverse(input1))

	/*
	10. 有 5 个人坐在一起，问第五个人多少岁？他说比第 4 个人大 2 岁。问第 4 个人岁数，他说比第 3 个人大 2 岁。问第三个人，
	又说比第 2 人大两岁。问第 2 个人，说比第一个人大两岁。最后问第一个人，他说是 10 岁。请问第五个人多大？
	 */
	fmt.Println("10. ------------------------")
	age := deduceAge(10)
	fmt.Println(age)


}

// 8 利用递归方法求阶层
func matchJC(p int) int {
	r := 1
	if p == 1 {
		return r
	}
	r = p * matchJC(p-1)
	return r
}

// 9 倒叙输出字符
func Reverse(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// 10. 根据五个人的对话推断年龄。
func deduceAge(n int) int {
	age := 10
	if n == 1 {
		return age
	}
	age = deduceAge(n-1) + 2
	return age
}