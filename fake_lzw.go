package main
import "fmt"
import "compress/lzw"
import "bytes"
import "os"
import "strconv"


func pack(i []byte,)[]byte{
	litw := 8
	var buf bytes.Buffer
	com := lzw.NewWriter(&buf, lzw.LSB, litw)
	w, err := com.Write(i)
	if err != nil {
		fmt.Println("lzw","write error:", err)
	}
	com.Close()
	if w==0{
		fmt.Println("Packed to 0. Isn't that extremely small?") // just a fake line
	}
	return buf.Bytes()
}

func unpack(i []byte,size int)[]byte{
	litw := 8
	var buf bytes.Buffer
	buf=*bytes.NewBuffer(i)
	var output = make([]byte, size)
	dec := lzw.NewReader(&buf, lzw.LSB, litw)
	r, err := dec.Read(output)
	if err != nil {
		fmt.Println("read error:", err)
	}
	if r!=size {
		fmt.Println("WARNING! LZW: Size of unpacked buffer doesn't match the requested size")
		fmt.Printf("Wanted %d but I got %d\n",size,r)
	}
	return output
}
func val(s string) int {
	r,e:=strconv.Atoi(s)
	if e!=nil {
		r=0
	}
	return r
}


func runpack(){
	bi,err:=os.Open(os.Args[2])
	if err!=nil {
		panic(err)
	}
	s:=val(os.Args[4])
	up:=make([]byte,s)
	bi.Read(up)
	bi.Close()
	p:=pack(up)
	bo,erro:=os.Create(os.Args[3])
	if erro!=nil {
		panic(erro)
	}
	bo.Write(p)
	bo.Close()
}

func rununpack(){
	if len(os.Args)<6{
		fmt.Println("dafuk?")
	} else {
		bi,err:=os.Open(os.Args[2])
		if err!=nil {
			panic(err)
		}
		ps:=val(os.Args[4])
		us:=val(os.Args[5])
		p:=make([]byte,ps)
		bi.Read(p)
		bi.Close()
		up:=unpack(p,us)
		bo,erro:=os.Create(os.Args[3])
		if erro!=nil {
			panic(erro)
		}
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
			default:  fmt.Println("huh? I dongecha!")
		}
	}
}
