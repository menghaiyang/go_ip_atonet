package main

import (
	"net"
	"encoding/binary"
	"strings"
	"strconv"
	"fmt"
)


// Convert net.IP to int32 转换之后是网络字节序列(第一个返回值),本机字节序(第二个返回值)
func inet_aton(ipnr net.IP) (uint32, uint32) {
	var sum uint32
    l := len(ipnr)
	if l !=  4 {
		if l != 0 {
			fmt.Println("ip error, ip len:",l)
		}

		return 0, 0
	}
	sum += uint32(uint8(ipnr[0])) << 24
	sum += uint32(uint8(ipnr[1])) << 16
	sum += uint32(uint8(ipnr[2])) << 8
	sum += uint32(uint8(ipnr[3]))

	netSequence := sum

	s := int16(0x1234)
	b := int8(s)
	//fmt.Println("int16字节大小为", unsafe.Sizeof(s)) //结果为2
	if 0x34 == b {
		//system is  little endian
		//需要将数据转换为网络字节序列，即大端。
		var tmp []byte = make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, sum)
		netSequence = binary.BigEndian.Uint32(tmp)
	} else {
		//system edian is big endian
	}

	return netSequence, sum
}


// Convert net.IP to int32 转换之后是网络字节序列(第一个返回值),本机字节序(第二个返回值)
func inet_atonsss(ipnr net.IP) (uint32, uint32) {

	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum uint32

	sum += uint32(b0) << 24
	sum += uint32(b1) << 16
	sum += uint32(b2) << 8
	sum += uint32(b3)

	netSequence := sum

	s := int16(0x1234)
	b := int8(s)
	//fmt.Println("int16字节大小为", unsafe.Sizeof(s)) //结果为2
	if 0x34 == b {
		//system is  little endian
		//需要将数据转换为网络字节序列，即大端。
		var tmp []byte = make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, sum)
		netSequence = binary.BigEndian.Uint32(tmp)
	} else {
		//system edian is big endian
	}

	return netSequence, sum
}

func main(){

	//b0, _ := strconv.Atoi("0")
	//	//fmt.Println("b0 : ",b0)

	//var ipnr = net.IP{1,0,0,1}
	//var ipnr2 = net.IP{0,0,0,0}

	//fmt.Println("ipnr: ",ipnr)
	//tmm, _ := inet_aton(ipnr)
	//fmt.Println("tmm: ",tmm)

	//tmm2, _ := inet_atonsss(ipnr2)
	//fmt.Println("tmm2: ",tmm2)
	//
	//kk := uint32(0) << 16
	//fmt.Println("kk: ",kk)

	var ipnr = net.IP{1,0,255,0}

	tmm, _ := inet_aton(ipnr)
	    fmt.Println("tmm: ",tmm)
	}














//import (
//	_ "github.com/go-sql-driver/mysql"
//	"database/sql"
//	"fmt"
//	//"time"
//)
//func main() {
//	db, err := sql.Open("mysql", "astaxie:astaxie@/test?charset=utf8")
//	checkErr(err)
//
//	//插入数据
//	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
//	checkErr(err)
//	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
//	checkErr(err)
//
//	id, err := res.LastInsertId()
//	checkErr(err)
//	fmt.Println(id)
//
//	//更新数据
//	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
//	checkErr(err)
//	res, err = stmt.Exec("astaxieupdate", id)
//	checkErr(err)
//
//
//	affect, err := res.RowsAffected()
//	checkErr(err)
//	fmt.Println(affect)
//
//
//	//查询数据
//	rows, err := db.Query("SELECT * FROM userinfo")
//	checkErr(err)
//	for rows.Next() {
//		var uid int
//		var username string
//		var department string
//		var created string
//		err = rows.Scan(&uid, &username, &department, &created)
//		checkErr(err)
//		fmt.Println(uid)
//		fmt.Println(username)
//		fmt.Println(department)
//		fmt.Println(created)
//	}
//	//删除数据
//	stmt, err = db.Prepare("delete from userinfo where uid=?")
//	checkErr(err)
//	res, err = stmt.Exec(id)
//	checkErr(err)
//	affect, err = res.RowsAffected()
//	checkErr(err)
//	fmt.Println(affect)
//	db.Close()
//}
//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

//
//// 处理/upload 逻辑
//func upload(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("method:", r.Method) //获取请求的方法
//	if r.Method == "GET" {
//		crutime := time.Now().Unix()
//		h := md5.New()
//		io.WriteString(h, strconv.FormatInt(crutime, 10))
//		token := fmt.Sprintf("%x", h.Sum(nil))
//		t, _ := template.ParseFiles("upload.gtpl")
//		t.Execute(w, token)
//	} else {
//		r.ParseMultipartForm(32 << 20)
//		file, handler, err := r.FormFile("uploadfile")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		fmt.Fprintf(w, "%v", handler.Header)
//		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer f.Close()
//		io.Copy(f, file)
//	}
//}
//
//func sayhelloName(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（ request body）
//	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
//	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	fmt.Println(r.Form["url_long"])
//	//for k, v := range r.Form {
//	//	fmt.Println("key:", k)
//	//	fmt.Println("val:", strings.Join(v, ""))
//	//}
//
//	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
//}
//func login(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	fmt.Println("method:", r.Method) //获取请求的方法
//	if r.Method == "GET" {
//		t, _ := template.ParseFiles("login.gtpl")
//		t.Execute(w, nil)
//	} else {
//		//请求的是登陆数据，那么执行登陆的逻辑判断
//		fmt.Println("username:", r.Form["username"])
//		fmt.Println("password:", r.Form["password"])
//	}
//}
//func main() {
//	http.HandleFunc("/", sayhelloName) //设置访问的路由
//	http.HandleFunc("/login", login) //设置访问的路由
//	err := http.ListenAndServe("127.0.0.1:9090", nil) //设置监听的端口
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//}

////一个简易的路由
//import (
//    "fmt"
//    "net/http"
//)
//
//type MyMux struct {
//}
//
//func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path == "/" {
//		sayhelloName(w, r)
//		return
//	}
//	http.NotFound(w, r)
//	return
//}
//
//func sayhelloName(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello myroute!")
//}
//
//func main() {
//	mux := &MyMux{}
//	http.ListenAndServe(":9090", mux)
//}


//import (
//	"fmt"
//)
//
//type Handler interface {//接口
//	Do(k, v interface{})
//}
//
//
//
//type HandlerFunc func(k, v interface{})
//
//func (f HandlerFunc) Do(k, v interface{}) {
//	f(k, v)
//}
//
//func Each(m map[interface{}]interface{}, h Handler) {
//	if m != nil && len(m) > 0 {
//		for k, v := range m {
//			h.Do(k, v)
//		}
//	}
//}
//
//func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
//	Each(m, HandlerFunc(f))
//}
//
//func selfInfo(k, v interface{}) {
//	fmt.Printf("大家好,我叫%s,今年%d岁\n", k, v)
//}
//
//func main() {
//	persons := make(map[interface{}]interface{})
//	persons["张三"] = 20
//	persons["李四"] = 23
//	persons["王五"] = 26
//
//	EachFunc(persons, selfInfo)
//
//}


//import "fmt"
//
////原始接口实现
////
//type Handler interface {
//	Do(k, v interface{})
//}
//
//func Each(m map[interface{}]interface{}, h Handler) {
//	if m != nil && len(m) > 0 {
//		for k, v := range m {
//			h.Do(k, v)
//		}
//	}
//}
//
//type welcome string
//
//func (w welcome) Do(k, v interface{}) {
//	fmt.Printf("%s,我叫%s,今年%d岁\n", w,k, v)
//}
//
//func main() {
//	persons := make(map[interface{}]interface{})
//	persons["张三"] = 20
//	persons["李四"] = 23
//	persons["王五"] = 26
//
//	var w welcome = "大家好"
//
//	Each(persons, w)
//}
//




//
//import (
//	"fmt"
//	"log"
//	"net/http"
//	"strings"
//)
//
//func sayHelloName(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm() //解析参数，默认是不会解析的
//	fmt.Println(r.Form)
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	fmt.Println(r.Form["url long"])
//	fmt.Println("---------------------------")
//
//	for k, v := range r.Form {
//		fmt.Println("key:", k)
//		fmt.Println("val:", strings.Join(v, ""))
//	}
//	fmt.Fprintf(w, "Hello shengsheng!")
//	fmt.Println("************************************")
//}
//func main() {
//
//	http.HandleFunc("/", sayHelloName)
//	err := http.ListenAndServe(":9090", nil)
//	if err != nil {
//		log.Fatal("ListenAdnserve: ", err)
//	}
//}
//


//
//
//import (
//	"fmt"
//	"net/http"
//)
//
//const (
//	SERVER_PORT       = 8081
//	SERVER_DOMAIN     = "localhost"
//	RESPONSE_TEMPLATE = "hellolllllworldddd"
//)
//
//func rootHandler(w http.ResponseWriter, req *http.Request) {
//
//	w.Header().Set("Content-Type", "text/html")
//	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
//	w.Write([]byte(RESPONSE_TEMPLATE))
//
//}
//
//func main() {
//	fmt.Println("hello worldd gogo")
//	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
//	http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "rui.crt", "rui.key", nil)
//}
