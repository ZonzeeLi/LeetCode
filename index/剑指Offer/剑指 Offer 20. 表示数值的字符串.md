## 剑指 Offer 20. 表示数值的字符串

### 题目传送门

[点击这里](https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/)

### 解题思路

这道题我觉得算是很难的一道题，首先这个方法，有限状态自动机，很难想到，再说明一下这道题意，这道题题意很好理解，就是让我们处理一个字符串做表示的数字是否合法，按照题目给找的规则来处理即可，但是情况众多，会出一某一种情况变成另一种情况后再变回去的情况，这里使用的有限状态机自动机简单说明一下，它说明的某一种状态(可以称作初始状态)，可以跳转到下一个接收状态，或者是被拒绝，这样多次的转换，就形成了一个集合(表)，来记录出现过的状态，并且可以表示出其转换的规则，如果接下来的状态不是可接收状态，就会被拒绝，自动机的思路和正则表达式有很大的关系，类似于字符串匹配的KMP算法等。如果没有接触过编译原理或者看过很多CS相关的书籍，可能对状态机不会有一定的了解，可以用这道题对其有个简单的认识，如果是电信的同学可以类比成一些数字电路中的信号跳转。

### 代码

```go
type State int
type CharType int

// 有限状态自动机
// 记录状态(当前状态/接收状态)，状态自动机，就是表明可以从某一个初始状态转移到另一个接收状态，否则就是被拒绝
const (
	STATE_INITIAL State = iota
	STATE_INT_SIGN
	STATE_INTEGER
	STATE_POINT
	STATE_POINT_WITHOUT_INT
	STATE_FRACTION
	STATE_EXP
	STATE_EXP_SIGN
	STATE_EXP_NUMBER
	STATE_END
)

const (
	CHAR_NUMBER CharType = iota
	CHAR_EXP
	CHAR_POINT
	CHAR_SIGN
	CHAR_SPACE
	CHAR_ILLEGAL
)

// 将字符串中的byte类型传换成定义好的类型
func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case 'e', 'E':
		return CHAR_EXP
	case '.':
		return CHAR_POINT
	case '+', '-':
		return CHAR_SIGN
	case ' ':
		return CHAR_SPACE
	default:
		return CHAR_ILLEGAL
	}
}

func isNumber(s string) bool {
    // 我们根据所有情况统计出可以从某一个当前状态根据出现的字符转换为其可接收的状态
	transfer := map[State]map[CharType]State{
		STATE_INITIAL: map[CharType]State{ // 初始状态
			CHAR_SPACE:  STATE_INITIAL, // 空格字符->初始状态 (表明当前字符后面转换成可接收的状态，下面同理)
			CHAR_NUMBER: STATE_INTEGER, // 数字字符->数字
			CHAR_POINT:  STATE_POINT_WITHOUT_INT, // 标点符->标点前无数字
			CHAR_SIGN:   STATE_INT_SIGN, // 符号字符->符号后数字
		},
		STATE_INT_SIGN: map[CharType]State{ // 符号后数字
			CHAR_NUMBER: STATE_INTEGER, // 数字字符->数字
			CHAR_POINT:  STATE_POINT_WITHOUT_INT, // 标点符->标点前无数字
		},
		STATE_INTEGER: map[CharType]State{ // 数字状态
			CHAR_NUMBER: STATE_INTEGER, // 数字字符->数字
			CHAR_EXP:    STATE_EXP, // e/E字符->e/E状态
			CHAR_POINT:  STATE_POINT, // 标点字符->标点前有数字
			CHAR_SPACE:  STATE_END, // 空格字符->结束
		},
		STATE_POINT: map[CharType]State{ // 标点前有数字
			CHAR_NUMBER: STATE_FRACTION, // 数字字符->小数位
			CHAR_EXP:    STATE_EXP, // e/E字符->e/E状态
			CHAR_SPACE:  STATE_END, // 空格字符->结束状态
		},
		STATE_POINT_WITHOUT_INT: map[CharType]State{ // 标点前无数字
			CHAR_NUMBER: STATE_FRACTION, // 数字字符->小数位
		},
		STATE_FRACTION: map[CharType]State{ // 小数位
			CHAR_NUMBER: STATE_FRACTION, // 数字字符->小数位
			CHAR_EXP:    STATE_EXP, // e/E字符->e/E状态
			CHAR_SPACE:  STATE_END, // 空格字符->结束状态
		},
		STATE_EXP: map[CharType]State{ // e/E状态
			CHAR_NUMBER: STATE_EXP_NUMBER, // 数字字符->e/E后数字
			CHAR_SIGN:   STATE_EXP_SIGN, // 符号字符->e/E后符号
		},
		STATE_EXP_SIGN: map[CharType]State{ // e/E后符号
			CHAR_NUMBER: STATE_EXP_NUMBER, // 数字字符->e/E后数字
		},
		STATE_EXP_NUMBER: map[CharType]State{ // e/E后数字
			CHAR_NUMBER: STATE_EXP_NUMBER, // 符号字符->e/E后符号
			CHAR_SPACE:  STATE_END, // 空格字符->结束状态
		},
		STATE_END: map[CharType]State{ // 结束状态
			CHAR_SPACE: STATE_END, // 空格字符->结束状态
		},
	}
	state := STATE_INITIAL // 统计完后从最开始的初始化状态开始遍历字符
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i]) // 转换
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ] // 状态自动机，转换状态
		}
	}
    // 可作为结束状态的几种情况
	return state == STATE_INTEGER || state == STATE_POINT || state == STATE_FRACTION || state == STATE_EXP_NUMBER || state == STATE_END
}
```