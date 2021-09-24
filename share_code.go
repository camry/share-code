package share_code

import (
    "strings"
)

const (
    Base       = "YCNH4RL8GJMTXFE12BWVKZSDIU6O9Q73A5P" // 把0剔除，当做补位符号，比如小于四位的邀请码在高位补0，这样36进制就变成了35进制。
    SuffixChar = "0"                                   // 补位字符，不能与自定义重复
    BinLen     = 35                                    // 进制长度：len(Base)
    CodeLen    = 6                                     // 生成邀请码最小长度
)

func strPadLeft(input string, padLength int, padString string) string {
    output := padString

    for padLength > len(output) {
        output += output
    }

    if len(input) >= padLength {
        return input
    }

    return output[:padLength-len(input)] + input
}

func strRev(s string) string {
    n := len(s)
    runes := make([]rune, n)

    for _, v := range s {
        n--
        runes[n] = v
    }

    return string(runes[n:])
}

func pow(a, b int64) int64 {
    var p int64

    p = 1

    for b > 0 {
        if b&1 != 0 {
            p *= a
        }

        b >>= 1
        a *= a
    }

    return p
}

func Encode(id int64) string {
    code := ""

    for id > 0 {
        mod := id % BinLen
        id = (id - mod) / BinLen

        code = string(Base[mod]) + code
    }

    if len(code) < 4 {
        code = strPadLeft(code, CodeLen, SuffixChar)
    }

    return code
}

func Decode(code string) int64 {
    var num int64

    index := strings.LastIndex(code, "0")

    if index != -1 {
        code = code[index+1:]
    }

    code = strRev(code)

    for i := 0; i < len(code); i++ {
        num += int64(strings.Index(Base, string(code[i]))) * pow(BinLen, int64(i))
    }

    return num
}
