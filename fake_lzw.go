package main
import "fmt"
import "compress/lzw"
import "bytes"
import "io"


func pack(i []byte,)[]byte{
	litw := 8
	var buf bytes.Buffer
	com := lzw.NewWriter(&buf, lzw.LSB, litw)
	w, err := com.Write(data)
	if err != nil {
		fmt.Println("lzw","write error:", err)
	}
	com.Close()
	return buf.Bytes()
}

func unpack(i []byte,size int)[]byte{
	litw := 8
	buf:=bytes.NewBuffer(i)
	var output = make([]byte, len(data))
	dec := lzw.NewReader(&buf, lzw.LSB, litw)
	r, err := dec.Read(output)
	if err != nil {
		fmt.Println("read error:", err)
	}
	if r!=size {
		fmt.Println("WARNING! LZW: Size of unpacked buffer doesn't match the requested size")
	}
	return output
}
func Val(s string) int {
	r,e:=strconv.Atoi(s)
	if e!=nil {
		r=0
	}
	return r
}


func runpack(){
	bi:=os.Open(os.Args[2])
	s:=val(os.Args[4])
	up:=os.Read(s)
	bi.Close()
	p:=pack(up)
	bo:=os.Create(os.Args[3])
	bo.Write(p)
	bo.Close()
}

func rununpack(){
	if len(os.Args)<6{
		fmt.Println("dafuk?")
	} else {
		bi:=os.Open(os.Args[2])
		ps:=val(os.Args[4])
		us:=val(os.Args[5]
		p:=os.Read(ps)
		bi.Close()
		up:=unpack(us)
		bo:=os.Create(os.Args[3])
		bo.Write(up)
		bo.Close()
	}
}

func main(){
	if len(os.Args)<5 {
		fmt.Println("Wadayawant?")
	} else {
		switch(os.Args[1]){
			case "p": runpack()
			case "u": rununpack()
			default:  fmt.Prints("huh? I dongecha!")
		}
	}
}
