/*
  fake_lzma.go
  
  version: 17.12.01
  Copyright (C) 2017 Jeroen P. Broks
  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.
  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:
  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
*/
package main
import "fmt"
import "github.com/itchio/lzma"
import "bytes"
import "os"
import "strconv"


func pack(i []byte,)[]byte{
	//litw := 8
	var buf bytes.Buffer
	com := lzma.NewWriterLevel(&buf,9)//, lzma.LSB, litw)
	w, err := com.Write(i)
	if err != nil {
		fmt.Println("lzma","write error:", err)
	}
	com.Close()
	if w==0{
		fmt.Println("Packed to 0. Isn't that extremely small?") // just a fake line
	}
	return buf.Bytes()
}

func unpack(i []byte,size int)[]byte{
	//litw := 8
	var buf bytes.Buffer
	buf=*bytes.NewBuffer(i)
	var output = make([]byte, size)
	dec := lzma.NewReader(&buf) //, lzma.LSB, litw)
	r, err := dec.Read(output)
	if err != nil {
		fmt.Println("read error:", err)
	}
	if r!=size {
		fmt.Println("WARNING! lzma: Size of unpacked buffer doesn't match the requested size")
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

