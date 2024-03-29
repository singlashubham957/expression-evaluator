package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"regex/solve"
	"regexp"
	_ "strconv"
	"strings"
	"unicode"
)

var expr string
var set=make(map[string]interface{})
var arr []string
var match [][]string
var count int=0

func home(c *gin.Context){
	for k:= range  set{

delete(set,k)
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Type your Expression",
	})
}
func express(c *gin.Context){
	expr=c.PostForm("expr")
	strings.Replace(expr, "{", "(", -1)
	strings.Replace(expr, "[", "(", -1)
	strings.Replace(expr, "}", ")", -1)
	strings.Replace(expr, "]", ")", -1)

	myExp := regexp.MustCompile(`\{*\s*([a-zA-Z]+)\s*\}*\s*=\s*\{*\s*([a-zA-Z]+)\s*\}*`)
	match = myExp.FindAllStringSubmatch(expr,-1)

	for i := 0; i < len(match); i++ {
		set[match[i][1]]=-1
		set[match[i][2]]=-1
		fmt.Printf("%s %s\n", match[i][1], match[i][2])
	}
	fmt.Println(set)
	c.HTML(200, "new_form.html", gin.H{"set":set})
}
func eval(c *gin.Context) {

	for k := range set {
		set[k] = c.PostForm(k);
		fmt.Println(k, set[k]);
	}
	fmt.Println(set)
	for i := 0; i < len(match); i++ {
		if set[match[i][1]] == set[match[i][2]] {
			arr = append(arr, "1")
		} else {
			arr = append(arr, "0")
		}
	}

	ans := "";a:=len(expr);
	for i, _ := range expr {

		if expr[i] == '(' || expr[i] == ')' {
			ans += string(expr[i]);
		} else if expr[i] == '=' {
			ans += arr[count];
			count++;
		} else if i+2<a&&expr[i]=='o'&&expr[i+1]=='r'&&
			!unicode.IsLetter(rune(expr[i-1]))&&!unicode.IsDigit(rune(expr[i+2])){
			ans+="|";
		}else if i+3 < a && expr[i] == 'a' && expr[i+1] == 'n' && expr[i+2] == 'd' &&
			!unicode.IsLetter(rune(expr[i-1]))&&!unicode.IsDigit(rune(expr[i+3])) {
			ans+="&";
		}
	}

	ans1 := solve.Topostfix(ans)
	ans2 := solve.Eva(ans1);
	c.JSON(http.StatusOK, gin.H{"data": ans2})
}


func main(){
	router:=gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/", home)
	router.POST("/", express)
	router.POST("/eval", eval)
	router.Run(":8088")

}