package main

import (
    "encoding/json"
    "fmt"
    "time"
    "github.com/lxn/walk"
	"github.com/anvie/port-scanner"
    "net"
	"strings"
	"sort"
	"golang.org/x/text/encoding/korean"
	"bytes"
	"syscall"
	"os/exec"
	"github.com/faith/color"
	"golang.org/x/sys/windows/registry"
	"strconv"
    "os"
    "bufio"
    "html/template"
)

type Requ struct{
 Yes bool
}

type Peer struct {
	Index   int
	ip     string
	checked bool
}

type PeerModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Peer
}

func NewPeerModel() *PeerModel {
        empty:=[]string{"empty"}
	m := new(PeerModel)
	m.ResetRows(1,empty)
	return m
}

func (m *PeerModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *PeerModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index

	case 1:
		return item.ip
	}

	panic("unexpected col")
}

// Called by the TableView to retrieve if a given row is checked.
func (m *PeerModel) Checked(row int) bool {
	return m.items[row].checked
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *PeerModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Called by the TableView to sort the model.
func (m *PeerModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.SliceStable(m.items, func(i, j int) bool {
		a, b := m.items[i], m.items[j]

		c := func(ls bool) bool {
			if m.sortOrder == walk.SortAscending {
				return ls
			}

			return !ls
		}

		switch m.sortColumn {
		case 0:
			return c(a.Index < b.Index)

		case 1:
			return c(a.ip < b.ip)

		}

		panic("unreachable")
	})

	return m.SorterBase.Sort(col, order)
}

func (m *PeerModel) ResetRows(count int,ipaddr []string) {
	// Create some random data.
	vname:=ipaddr

	m.items = make([]*Peer, count-1)

	//now := time.Now()

	for i := range m.items {
		m.items[i] = &Peer{
			Index: i+1,
			ip:   vname[i],
		}
	}

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}

func serverstart(stip *walk.StatusBarItem,onoff *walk.StatusBarItem) {
var ip string="IP:"
    service := ":1200"
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    checkError(err)
         //ip주소 구하기
    addrs, err := net.InterfaceAddrs()
  if err != nil {
                 fmt.Println(err)
                
         }
for _, address := range addrs {

               // check the address type and if it is not a loopback the display it
               if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	       fmt.Println(ipnet.IP.String())
                  if ipnet.IP.To4() != nil {
                     ip+=ipnet.IP.String()
		     break
                  }

               }
         }

    stip.SetText(ip)
	    

    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err!=nil{
     return
    }

    onoff.SetText("서버상태:ON")
    for {
        conn, err := listener.Accept()
        if err != nil {
            
        }
    wr2:=&WResult{}
        encoder := json.NewEncoder(conn)
        decoder := json.NewDecoder(conn)
    
        var requ Requ
        decoder.Decode(&requ)
	if(requ.Yes==true){
	 fmt.Println("정상적 요청")
	 stest1(wr2)
	 stest2(wr2)
	 stest3(wr2)
	 stest4(wr2)
	 stest5(wr2)
	 stest6(wr2)
	 stest7(wr2)
	 stest8(wr2)
	 stest9(wr2)
	 stest10(wr2)
	 stest11(wr2)
	 stest12(wr2)
	 stest13(wr2)
	 stest14(wr2)
	 stest15(wr2)
	 stest16(wr2)
	 stest17(wr2)
	 stest18(wr2)
	 stest19(wr2)
	 encoder.Encode(wr2)
	}else{
	 fmt.Println("스캐너인거 같어")
	}
        conn.Close()
    }
}

//보고서 받아오기(클라이언트역할)
func pdown(m *PeerModel,tv *walk.TableView,done chan bool,startyn *int,pro *walk.ProgressBar){
var result int
ti:=tv.ItemChecker()
//체크된거 구하기
for i:=0;i<len(m.items);i++{
select{
case <- done:
return
default:
if ti.Checked(i){
requ := Requ{Yes:true,}
service :=fmt.Sprint(m.items[i].ip,":1200")
conn, err := net.Dial("tcp", service)

if err!=nil{
fmt.Println("error")
return
}
encoder := json.NewEncoder(conn)
decoder := json.NewDecoder(conn)

encoder.Encode(requ)

var wresult WResult
decoder.Decode(&wresult)

fmt.Println(wresult)
soutput_result(wresult,m.items[i].ip)
result=(i+1)*100/len(m.items)
pro.SetValue(result)
}else{
result=(i+1)*100/len(m.items)
pro.SetValue(result)
continue
}
}//select
}//for
*startyn=0
}


func peerscan(pro *walk.ProgressBar,startyn *int,done chan bool,m *PeerModel){
var ip string //변수선언
var naddr string
//local 아이피 주소구하기
  addrs, err := net.InterfaceAddrs()
  if err != nil {
                 fmt.Println(err)
                
         }
for _, address := range addrs {

               // check the address type and if it is not a loopback the display it
               if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	       fmt.Println(ipnet.IP.String())
                  if ipnet.IP.To4() != nil {
                     ip=ipnet.IP.String()
		     break
                  }

               }
         }
//c클래스 여부 확인
addr := net.ParseIP(ip)
mask := addr.DefaultMask()
network := addr.Mask(mask)
ones, bits := mask.Size()
fmt.Println(bits)
if(ones==24){
fmt.Println("c클래스입니다.")
fmt.Println(network.String())
naddr=network.String()
naddr=naddr[:len(naddr)-2]
fmt.Println(naddr)
}else{
fmt.Println("c클래스가 아닙니다.")
}
ct:=0//피어 활성화 개수
peerips:=""
count:=2
 for j:=2; j<255;j++{
     select{
     case <- done:
     return
     default:
     scanip:=fmt.Sprint(naddr,".",j)
     fmt.Println("ipaddress:",scanip)
     ps := portscanner.NewPortScanner(scanip, 1*time.Second, 1)

     // get opened port
     fmt.Printf("port %d-%d scan\n", 1200, 1200)

     openedPorts := ps.GetOpenedPort(1200, 1200)

     for i := 0; i < len(openedPorts); i++ {
     	port := openedPorts[i]
     	fmt.Print(" ", port, " [open]")
     	fmt.Println("  -->  ", ps.DescribePort(port))
	peerips+=scanip
	peerips+=","
	ct++
	fmt.Println("ct:",ct)
	fmt.Println("peerips:",peerips)
     }
     count++
     pro.SetValue(count)
    }
    }
    //초기화 부분
    ipresult:=strings.Split(peerips,",")
    for k:=0;k<len(ipresult);k++{
    fmt.Println(ipresult[k])
    }
m.ResetRows(len(ipresult),ipresult)
    *startyn=0
    fmt.Println("끝")
}

func stest1(w *WResult){
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta,color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","1. ","Administrator 계정 관리 #####")
d.Println("[기준] : 관리자 계정이 하나만 존재할 경우에는 양호")
cmd := exec.Command("cmd","/C","net","localgroup","administrators")
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 }
 result1=string(out)
 w.Test1s=result1
d.Println("[현황]")
d.Println(result1)
str:=strings.Split(result1,"\n")
if strings.Contains(str[7],"명령을"){
f.Println("[결과] Administrator 계정관리 : 양호합니다.\n")
w.Test1="#01DF01"
w.Yes=w.Yes+1
}else{
c.Println("[결과] Administrator 계정관리 : 취약합니다.\n")
w.Test1="#DF0101"
w.No=w.No+1
}
}

//Guest 계정관리
func stest2(w *WResult){
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta,color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","2. ","Guest 계정 관리 #####")
d.Println("[기준] : Guest 계정이 비활성화일 경우에 양호")
cmd := exec.Command("cmd","/C","net","user","guest")
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 }
 result1=string(out)
 w.Test2s=result1
d.Println("[현황]")
d.Println(result1)
str:=strings.Split(result1,"\n")
if strings.Contains(str[5],"아니요"){
f.Println("[결과] Guest 계정관리 : 양호합니다.\n")
w.Test2="#01DF01"
w.Yes=w.Yes+1
}else{
c.Println("[결과] Guest 계정관리 : 취약합니다.\n")
w.Test2="#DF0101"
w.No=w.No+1
}
}

//계정 잠금 정책 설정
func stest3(w *WResult){
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","3. ","계정 잠금 정책 설정 #####")
d.Println("[기준] : 계정 잠금 임계값 5번 이하,계정 잠금 기간,계정 잠금기간 원래대로 설정 60분이상")
cmd := exec.Command("cmd","/C","net","accounts")
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 }
 result1=string(out)
 w.Test3s=result1
d.Println("[현황]")
str:=strings.Split(result1,"\n")
d.Println(str[5])
d.Println(str[6])
d.Println(str[7])
one:=strings.Fields(str[5])
two:=strings.Fields(str[6])
three:=strings.Fields(str[7])
//d.Println(one[2])
//d.Println(two[3])
//d.Println(three[4])
num1, err :=strconv.Atoi(one[2])
num2, err :=strconv.Atoi(two[3])
num3, err :=strconv.Atoi(three[4])
if( one[2]=="아님" || num1 > 5 ){
c.Println("[결과] 계정 잠금 정책 설정 : 취약합니다.\n");
w.Test3="#DF0101"
w.No=w.No+1
}else if( num2 < 60 ){
c.Println("[결과] 계정 잠금 정책 설정 : 취약합니다.\n");
w.Test3="#DF0101"
w.No=w.No+1
}else if( num3 < 60 ){
c.Println("[결과] 계정 잠금 정책 설정 : 취약합니다.\n");
w.Test3="#DF0101"
w.No=w.No+1
}else{
f.Println("[결과] 계정 잠금 정책 설정 : 양호합니다.\n");
w.Test3="#01DF01"
w.Yes=w.Yes+1
}
}

//암호정책설정
func stest4(w *WResult){
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","4. ","암호 정책 설정 #####")
d.Println("[기준] :")
d.Println("1] 최소 암호 사용기간 1일 이상")
d.Println("2] 최대 암호 사용기간 90일 이하")
d.Println("3] 최소 암호 길이 8문자 이상")
d.Println("4] 최근 암호 기억 12개로 설정되어 있으면 양호")

cmd := exec.Command("cmd","/C","net","accounts")
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 }
 result1=string(out)
 w.Test4s=result1
d.Println("[현황]")
str:=strings.Split(result1,"\n")
d.Println(str[1])
d.Println(str[2])
d.Println(str[3])
d.Println(str[4])
one:=strings.Fields(str[1])
two:=strings.Fields(str[2])
three:=strings.Fields(str[3])
four:=strings.Fields(str[4])
//d.Println(one[5])
//d.Println(two[5])
//d.Println(three[3])
//d.Println(four[3])
num1, err :=strconv.Atoi(one[5])
num2, err :=strconv.Atoi(two[5])
num3, err :=strconv.Atoi(three[3])
num4, err :=strconv.Atoi(four[3])
if( num1 < 1 ){
c.Println("[결과] 암호 정책 설정 : 취약합니다.\n");
w.Test4="#DF0101"
w.No=w.No+1
}else if( num2 > 90 ){
c.Println("[결과] 암호 정책 설정 : 취약합니다.\n");
w.Test4="#DF0101"
w.No=w.No+1
}else if( num3 < 8 ){
c.Println("[결과] 암호 정책 설정 : 취약합니다.\n");
w.Test4="#DF0101"
w.No=w.No+1
}else if( num4 != 12){
c.Println("[결과] 암호 정책 설정 : 취약합니다.\n");
w.Test4="#DF0101"
w.No=w.No+1
}else{
f.Println("[결과] 암호 정책 설정 : 양호합니다.\n");
w.Test4="#01DF01"
w.Yes=w.Yes+1
}
}

//사용자계정 컨트롤설정
func stest5(w *WResult){
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","5. ","사용자 계정 컨트롤 설정 #####")
d.Println("[기준] : 사용자 계정 컨트롤 사용하면 양호")
k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k.Close()

s, _, err := k.GetIntegerValue("PromptOnSecureDesktop")
if err != nil {
fmt.Println(err)
}
fmt.Println("[현황]")
fmt.Println("PromptOnSecureDesktop : ",s)
str:=strconv.Itoa(int(s))
result1="PromptOnSecureDesktop : "+str
w.Test5s=result1
if(s < 1){
c.Println("[결과] 사용자 계정 컨트롤 : 취약합니다.\n");
w.Test5="#DF0101"
w.No=w.No+1
}else{
f.Println("[결과] 사용자 계정 컨트롤 : 양호합니다.\n");
w.Test5="#01DF01"
w.Yes=w.Yes+1
}
}


//CMD파일 권한 설정
func stest6(w *WResult){
yorn:=0
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","6. ","CMD 파일 권한 설정 #####")
d.Println("[기준] : Administrator와 System 과 TrustedInstaller 그룹만실행 권한이 설정되어 있으면 양호")
cmd := exec.Command("cmd","/C","cacls","%systemroot%\\System32\\cmd.exe")
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 fmt.Println(err)
 }
 result1=string(out)
 fmt.Println(result1)
 w.Test6s=result1
 str:=strings.Split(result1,"\n")

for i:=0;i<len(str);i++{
   if strings.Contains(strings.ToLower(str[i]),"administrator"){
                continue
   }else if strings.Contains(strings.ToLower(str[i]),"system:"){
                continue
   }else if strings.Contains(strings.ToLower(str[i]),"trustedinstaller:"){
                continue
   }else if len(str[i])<3 {
                continue
   }else{
                yorn=1
                break
   }
}
if yorn==0{
f.Println("[결과] CMD파일 권한 설정 : 양호합니다.\n")
w.Test6="#01DF01"
w.Yes=w.Yes+1
}else{
c.Println("[결과] CMD파일 권한 설정 : 취약합니다.\n")
w.Test6="#DF0101"
w.No=w.No+1
}
}


//사용자 디렉터리 접근제한
func stest7(w *WResult){
users:=0
everyone:=0
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","7. ","사용자 디렉터리 접근제한 #####")
d.Println("[기준] : 홈디렉터리 권한중 Users:F 와 Everyone:F가 없으면 양호")
cmd := exec.Command("cmd","/C","cacls","c:\\users\\*")
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 }
 result1=string(out)
 w.Test7s=result1
 str:=strings.Split(result1,"\n")

for i:=0;i<len(str);i++{
   if strings.Contains(str[i],"Users:(OI)(CI)F"){
          users=1
   }else if strings.Contains(str[i],"Everyone:(OI)(CI)F"){
          everyone=1
   }
}

if users==1 && everyone==1{
 c.Println("[결과] 사용자 디렉터리 접근제한 : 취약합니다. 이유:Users:F,Everyone:F\n")
w.Test7="#DF0101"
w.No=w.No+1
}else if users==1{
c.Println("[결과] 사용자 디렉터리 접근제한 : 취약합니다. 이유:Users:F\n")
w.Test7="#DF0101"
w.No=w.No+1
}else if everyone==1{
c.Println("[결과] 사용자 디렉터리 접근제한 : 취약합니다. 이유:Everyone:F\n")
w.Test7="#DF0101"
w.No=w.No+1
}else{
f.Println("[결과] 사용자 디렉터리 접근제한 : 양호합니다.")
w.Test7="#01DF01"
w.Yes=w.Yes+1
}
}


//하드디스크 기본공유 제거
func stest8(w *WResult){
yorn:=0
yorn2:=0
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","8. ","하드디스크 기본공유 제거 #####")
d.Println("[기준] : 레지스트리값 AutoShareWks가 0이며 기본공유가 존재하지 않을 경우[IPC$ 제외]")
cmd := exec.Command("cmd","/C","net","share")
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 }
 result1=string(out)
 w.Test8s=result1
 str:=strings.Split(result1,"\n")

for i:=4;i<len(str);i++{
   if strings.Contains(str[i],"IPC$"){
                continue
   }else if strings.Contains(str[i],"명령") {
                break
   }else{
                yorn=1
                break
   }
}

k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\lanmanserver\parameters`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k.Close()
s, _, err := k.GetIntegerValue("AutoShareWks")
if(s==0){
yorn2=0
}else{
yorn2=1
}

if(yorn==1||yorn2==1){
c.Println("[결과] 하드디스크 기본 공유제거 : 취약합니다.")
w.Test8="#DF0101"
w.No=w.No+1
}else{
f.Println("[결과] 하드디스크 기본 공유제거 : 양호합니다.")
w.Test8="#01DF01"
w.Yes=w.Yes+1
}
}

//SAM
func stest9(w *WResult){
yorn:=0
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","9. ","SAM 파일 접근 통제 #####")
d.Println("[기준] : SAM파일 접근권한이 Administrator,System 그룹만 모든 권한으로 등록되어 있는 경우")
d.Println("[현황] :")
cmd := exec.Command("cmd","/C","icacls","%systemroot%\\system32\\config\\SAM")
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 
 }

 result1=string(out)
 w.Test9s=result1
 d.Println(result1)
str:=strings.Split(result1,"\n")
for i:=0;i<len(str);i++{
  if strings.Contains(strings.ToLower(str[i]),"system"){
                continue
   }else if strings.Contains(strings.ToLower(str[i]),"administrator") {
                continue
   }else if len(str[i])<3 {
                continue
   }else if strings.Contains(str[i],"파일을") {
                break
   }else{
                yorn=1
                break
   }
}
if(yorn==0){
f.Println("[결과] SAM 파일 접근 통제 : 양호합니다.")
w.Test9="#01DF01"
w.Yes=w.Yes+1
}else{
c.Println("[결과] SAM 파일 접근 통제 : 취약합니다.")
w.Test9="#DF0101"
w.No=w.No+1
}
}



func stest10(w *WResult){
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta,color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","10. ","최신 서비스 팩 적용 #####")
d.Println("[기준] : 서비스팩 2 적용여부")
cmd := exec.Command("cmd","/C","systeminfo")
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 }
 result1=string(out)
w.Test10s=result1
d.Println("[현황]")
d.Println(result1)
str:=strings.Split(result1,"\n")
if strings.Contains(str[3],"Service Pack 2"){
f.Println("[결과] 최신 서비스 팩 적용 : 양호합니다.\n")
//Apendr(m,10,true)
w.Test10="#01DF01"
w.Yes=w.Yes+1
}else{
c.Println("[결과] 최신 서비스 팩 적용 : 취약합니다.\n")
//Apendr(m,10,false)
w.Test10="#DF0101"
w.No=w.No+1
}
//pro.SetValue(100)
}

func stest11(w *WResult){
yorn:=1//양호
//istr:=""
d := color.New(color.FgWhite, color.Bold)
//c := color.New(color.FgMagenta,color.Bold)
//f := color.New(color.FgGreen, color.Bold)
var result1,result2 string
d.Println("##### ","11. ","공유권한 및 사용자 그룹 설정 #####")
d.Println("[기준] : 공유디렉터리가 없거나 공유 디렉터리 접근 권한에 Everyone이 없으면 양호")
func(){
cmd := exec.Command("cmd","/C","net","share")
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 }
 result1=string(out)
 }()

w.Test11s=result1+"\n"
d.Println("[현황]")
d.Println(result1)
str:=strings.Split(result1,"\n")
for i:=4;i<len(str);i++{
fmt.Println(len(str[i]),":",str[i])
if(strings.Contains(str[i],"$")){
continue
}
if(len(str[i])==46 && strings.Contains(str[i],"Users")!=true){
continue
}
if(strings.Contains(str[i],"명령을")){
break
}
str2:=strings.Fields(str[i])
str3:=""
for k:=0;k<len(str2);k++{
if k==1{
str3=str2[k]
}

}
func(){
cmd := exec.Command("cmd","/C","icacls",str3)
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 }
 result2=string(out)
 }()
 w.Test11s+=result2
if strings.Contains(result2,"Everyone")==true{
yorn=0
}
}

if yorn==1{
fmt.Println("[결과] 공유권한 및 사용자 그룹 설정 : 양호합니다.")
//Apendr(m,11,true)
w.Test11="#01DF01"
w.Yes=w.Yes+1
}else{
fmt.Println("[결과] 공유권한 및 사용자 그룹 설정 : 취약합니다.")
//Apendr(m,11,false)
w.Test11="#DF0101"
w.No=w.No+1
}
//pro.SetValue(110)
}


//로그오프나 워크스테이션 잠김
func stest12(w *WResult){
yorn:=1
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","12. ","마지막 로그온 사용자 계정숨김 #####")
d.Println("[기준] : 화면보호기를 설정하고 암호를 사용하며 대기 시간이 5분이면 양호")

//ScreenSaveActive 1
k, err := registry.OpenKey(registry.CURRENT_USER, `Control Panel\Desktop`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k.Close()
s, _, err := k.GetStringValue("ScreenSaveActive")
if err != nil {
fmt.Println(err)
}

//ScreenSaverIsSecure 1
k2, err := registry.OpenKey(registry.CURRENT_USER, `Control Panel\Desktop`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k2.Close()

s2, _, err := k2.GetStringValue("ScreenSaveActive")
if err != nil {
fmt.Println(err)
}

//ScreenSaveTimeOut 300이상
k3, err := registry.OpenKey(registry.CURRENT_USER, `Control Panel\Desktop`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k3.Close()

s3, _, err := k3.GetStringValue("ScreenSaveActive")
if err != nil {
fmt.Println(err)
}

//현황
fmt.Println("[현황]")
fmt.Println("ScreenSaveActive : ",s)
fmt.Println("ScreenSaverIsSecure : ",s2)
fmt.Println("ScreenSaveTimeOut : ",s3)
str,_:=strconv.Atoi(string(s))
str2,_:=strconv.Atoi(string(s2))
str3,_:=strconv.Atoi(string(s3))
result1="ScreenSaveActive : "+s+"\n"
result1+="ScreenSaverIsSecure : "+s2+"\n"
result1+="ScreenSaveTimeOut : "+s3+"\n"
fmt.Println(result1)
w.Test12s=result1
if(str<1){
yorn=0
}
if(str2<1){
yorn=0
}
if(str3<300){
yorn=0
}

if(yorn == 1){
c.Println("[결과] 사용자 계정 컨트롤 : 양호합니다.\n");
//Apendr(m,12,true)
w.Test12="#01DF01"
w.Yes=w.Yes+1
}else{
f.Println("[결과] 사용자 계정 컨트롤 : 취약합니다.\n");
//Apendr(m,12,false)
w.Test12="#DF0101"
w.No=w.No+1
}
//pro.SetValue(120)
}

//이벤트 뷰어 설정
func stest13(w *WResult){
yorn:=1
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","13. ","이벤트 뷰어 설정 #####")
d.Println("[기준] : 최대 로그 크기 10240KB 이상이고, 로그 덮어쓰기 설정 옵션이 0 으로 설정되면 양호")
//응용프로그램 크기 조회
k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\Eventlog\Application`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k.Close()
s, _, err := k.GetIntegerValue("MaxSize")
if err != nil {
fmt.Println(err)
}

//보안 로그크기 조회
k2, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\Eventlog\Security`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k2.Close()
s2, _, err := k2.GetIntegerValue("MaxSize")
if err != nil {
fmt.Println(err)
}

//시스템 로그크기 조회
k3, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\Eventlog\System`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k3.Close()
s3, _, err := k3.GetIntegerValue("MaxSize")
if err != nil {
fmt.Println(err)
}

//응용 프로그램 로그 덮어 쓰기 설정 옵션
k4, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\Eventlog\Application`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k4.Close()
s4, _, err := k4.GetIntegerValue("Retention")
if err != nil {
fmt.Println(err)
}

//보안 로그 덮어쓰기 설정 옵션
k5, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\Eventlog\Security`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k5.Close()
s5, _, err := k5.GetIntegerValue("Retention")
if err != nil {
fmt.Println(err)
}

//시스템 로그 덮어쓰기 설정 옵션
k6, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\Eventlog\System`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k6.Close()
s6, _, err := k6.GetIntegerValue("Retention")
if err != nil {
fmt.Println(err)
}

fmt.Println("[현황]")
fmt.Println("Application MaxSize",s)
fmt.Println("Security MaxSize",s2)
fmt.Println("System MaxSize",s3)
fmt.Println("Application Retention",s4)
fmt.Println("Security Retention",s5)
fmt.Println("System Retention",s6)
str:=strconv.Itoa(int(s))
str2:=strconv.Itoa(int(s2))
str3:=strconv.Itoa(int(s3))
str4:=strconv.Itoa(int(s4))
str5:=strconv.Itoa(int(s5))
str6:=strconv.Itoa(int(s6))
result1="응용프로그램 로그크기 : "+str+"\n"
result1+="보안 로그크기 : "+str2+"\n"
result1+="시스템 로그크기 : "+str3+"\n"
result1+="응용프로그램 로그 덮어쓰기 설정옵션 : "+str4+"\n"
result1+="보안 로그 덮어쓰기 설정옵션 : "+str5+"\n"
result1+="시스템 로그 덮어쓰기 설정옵션 : "+str6+"\n"
fmt.Println(result1)
w.Test13s=result1
if(s<10240){
yorn=0
}
if(s2<10240){
yorn=0
}
if(s3<10240){
yorn=0
}
if(s4!=0){
yorn=0
}
if(s5!=0){
yorn=0
}
if(s6!=0){
yorn=0
}

if(yorn == 1){
f.Println("[결과] 이벤트 뷰어 설정 : 양호합니다.\n");
//Apendr(m,13,true)
w.Test13="#01DF01"
w.Yes=w.Yes+1
}else{
c.Println("[결과] 이벤트 뷰어 설정 : 취약합니다.\n");
//Apendr(m,13,false)
w.Test13="#DF0101"
w.No=w.No+1
}

//pro.SetValue(130)

}

//마지막 로그온 사용자 계정숨김
func stest14(w *WResult){
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","14. ","마지막 로그온 사용자 계정숨김 #####")
d.Println("[기준] : 마지막 로그온 사용자 숨김 설정 사용으로 설정돼 있으면 양호")
k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k.Close()

s, _, err := k.GetIntegerValue("DontDisplayLastUserName")
if err != nil {
fmt.Println(err)
}
fmt.Println("[현황]")
fmt.Println("DontDisplayLastUserName : ",s)
str:=strconv.Itoa(int(s))
result1="DontDisplayLastUserName : "+str
fmt.Println(result1)
w.Test14s=result1
if(s == 1){
f.Println("[결과] 마지막 로그온 사용자 계정숨김 : 양호합니다.\n");
//Apendr(m,14,true)
w.Test14="#01DF01"
w.Yes=w.Yes+1
}else{
c.Println("[결과] 마지막 로그온 사용자 계정숨김 : 취약합니다.\n");
//Apendr(m,14,false)
w.Test14="#DF0101"
w.No=w.No+1
}
//pro.SetValue(140)
}

//로그온하지 않은 사용자 시스템종료 방지
func stest15(w *WResult){
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","15. ","로그온하지 않은 사용자 시스템종료 방지 #####")
d.Println("[기준] : 로그온하지 않고 시스템 종료 허용이 사용안함으로 설정되어 있을 경우 양호")
k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k.Close()

s, _, err := k.GetIntegerValue("ShutdownWithoutLogon")
if err != nil {
fmt.Println(err)
}
fmt.Println("[현황]")
fmt.Println("ShutdownWithoutLogon : ",s)
str:=strconv.Itoa(int(s))
result1="ShutdownWithoutLogon : "+str
fmt.Println(result1)
w.Test15s=result1
if(s == 0){
f.Println("[결과] 마지막 로그온 사용자 계정숨김 : 양호합니다.\n");
//Apendr(m,15,true)
w.Test15="#01DF01"
w.Yes=w.Yes+1
}else{
c.Println("[결과] 마지막 로그온 사용자 계정숨김 : 취약합니다.\n");
//Apendr(m,15,false)
w.Test15="#DF0101"
w.No=w.No+1
}
//pro.SetValue(150)

}

//백신 프로그램 설치
func stest16(w *WResult){
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta,color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","16. ","백신프로그램 설치 #####")
d.Println("[기준] : 백신 프로그램 설치여부")
cmd := exec.Command("cmd","/C","tasklist")
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 }
 result1=string(out)
 w.Test16s=result1
d.Println("[현황]")
d.Println(result1)
if strings.Contains(result1,"AYCAgentSrvayc"){
f.Println("[결과] 백신프로그램 설치 : 양호합니다.\n")
//Apendr(m,16,true)
w.Test16="#01DF01"
w.Yes=w.Yes+1
}else{
c.Println("[결과] 백신프로그램 설치 : 취약합니다.\n")
//Apendr(m,16,false)
w.Test16="#DF0101"
w.No=w.No+1
}
//pro.SetValue(160)

}

//Null Session 설정
func stest17(w *WResult){
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","17. ","Null Session 설정 #####")
d.Println("[기준] : 해당 레지스트리 값이 설정되어 있으면 양호")
k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\LSA`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k.Close()

s, _, err := k.GetIntegerValue("restrictanonymous")
if err != nil {
fmt.Println(err)
}
fmt.Println("[현황]")
fmt.Println("restrictanonymous : ",s)
str:=strconv.Itoa(int(s))
result1="restrictanonymous : "+str
fmt.Println(result1)
w.Test17s=result1
if(s == 2){
f.Println("[결과] 마지막 로그온 사용자 계정숨김 : 양호합니다.\n");
//Apendr(m,17,true)
w.Test17="#01DF01"
w.Yes=w.Yes+1
}else{
c.Println("[결과] 마지막 로그온 사용자 계정숨김 : 취약합니다.\n");
//Apendr(m,17,false)
w.Test17="#DF0101"
w.No=w.No+1
}
//pro.SetValue(170)

}

//레지스트리 보호 차단
func stest18(w *WResult){
yorn:=0
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","18. ","레지스트리 보호 차단 #####")
d.Println("[기준] : Remote Registry Service 가 중지되어 있으면 양호")
cmd := exec.Command("cmd","/C","net start")
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
var o bytes.Buffer
cmd.Stdout=&o
cmd.Run()
out,err := korean.EUCKR.NewDecoder().Bytes(o.Bytes())
 if err != nil {
 }
 result1=string(out)
 w.Test18s=result1+"\n"
d.Println("[현황]")
d.Println(result1)
if strings.Contains(result1,"Remote Registry")==false{
f.Println("[결과] 레지스트리 보호 차단 : 양호합니다.\n");
//Apendr(m,18,true)
w.Test18="#01DF01"
w.Yes=w.Yes+1
}else{

k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\SecurePipeServers\winreg`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
}
defer k.Close()

s, _, err := k.GetStringValue("Description")
if err != nil {
fmt.Println(err)
}
fmt.Println("[현황]")
fmt.Println("Description : ",s)
w.Test18s+="Description: "+s
if(strings.Contains(s,"Registry Server")){
yorn=1
}

if yorn==0{
f.Println("[결과] 레지스트리 보호 차단 : 취약합니다.\n");
//Apendr(m,18,false)
w.Test18="#DF0101"
w.No=w.No+1
}else{
c.Println("[결과] 레지스트리 보호 차단 : 양호합니다.\n");
//Apendr(m,18,true)
w.Test18="#01DF01"
w.Yes=w.Yes+1
}
}
//pro.SetValue(180)


}

//AutoLogon 기능제어
func stest19(w *WResult){
d := color.New(color.FgWhite, color.Bold)
c := color.New(color.FgMagenta, color.Bold)
f := color.New(color.FgGreen, color.Bold)
var result1 string
d.Println("##### ","19. ","AutoLogon 기능제어 #####")
d.Println("[기준] : AutoAdminLogon 값이 없거나 0으로 설정되어 있으면 양호")
k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFRWARE\Microsoft\Windows NT\CurrentVersion\Winlogon`, registry.QUERY_VALUE)
if err != nil {
fmt.Println(err)
w.Test19s="접근권한 막힘"
c.Println("[결과] AutoLogon 기능제어 : 취약합니다.\n");
//Apendr(m,19,false)
w.Test19="#DF0101"
w.No=w.No+1
//pro.SetValue(190)
return
}
defer k.Close()

s, _, err := k.GetIntegerValue("AutoAdminLogon")
if err != nil {
fmt.Println(err)
}
fmt.Println("[현황]")
fmt.Println("AutoAdminLogon : ",s)
str:=strconv.Itoa(int(s))
result1="AutoAdminLogon : "+str
fmt.Println(result1)
w.Test5s=result1
if(s == 1){
c.Println("[결과] AutoLogon 기능제어 : 취약합니다.\n");
//Apendr(m,19,false)
w.Test19="#DF0101"
w.No=w.No+1
}else{
f.Println("[결과] AutoLogon 기능제어 : 양호합니다.\n");
//Apendr(m,19,true)
w.Test19="#01DF01"
w.Yes=w.Yes+1
}
//pro.SetValue(190)

}







//출력
func soutput_result(w WResult,num string){
ipstr:=strings.Replace(num,".","_",-1)
ipstr2:=fmt.Sprint(ipstr,".html")
file, err := os.OpenFile(
		ipstr2,
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, // 파일이 없으면 생성,
                                                  // 읽기/쓰기, 파일을 연 뒤 내용 삭제
		os.FileMode(0644),                // 파일 권한은 644
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음
fw := bufio.NewWriter(file)

t:=template.New("Result template")
t,err=t.Parse(temp)
checkError(err)

err = t.Execute(fw,w)
checkError(err)
     // io.Writer 인터페이스를 따르는 file로
                                       // io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성
	fw.Flush()
}