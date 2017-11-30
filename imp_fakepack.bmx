Rem
  imp_fakepack.bmx
  
  version: 17.11.30
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
End Rem
Strict

Rem

	You can import this file for quick testing.
	fakepack.bmx is for actually putting all this in a module.
	
End Rem

Import tricky_units.Dirry
Import jcr6.jcr6main
Import brl.bank

Incbin "fake_lzw.go"

Private

Global wd$ = Dirry("$AppSupport$/jcr6_fake/")
If Not CreateDir(wd,1) 
	Notify "Can't create: "+wd
	Print "Can't create: "+wd
	End
EndIf


Type JCR6_FAKEDRIVER Extends DRV_Compression
	
	Field name$
	
	Method installfake(n$)
		If Not CopyFile("incbin::"+n+".go",wd+n+".go")
			Print "GO ERROR!"
			End
		EndIf
	End Method
	
	Method compress:TBank(b:TBank)
	DeleteFile(wd+"1_packed")	
	SaveBank b,wd+"1_prepack"
	installfake name
	system_ "go run '"+wd+name+".go' p '"+wd+"1_prepack' '"+wd+"1_packed' "+BankSize(b)	
	Return LoadBank(wd+"1_packed")
	End Method

	Method Expand:TBank(b:TBank,truesize)
	DeleteFile(wd+"2_unpacked")	
	SaveBank b,wd+"2_preunpack"
	installfake name
	system_ "go run '"+wd+name+".go' p '"+wd+"1_prepack' '"+wd+"1_packed' "+BankSize(b)+" "+truesize
	Return LoadBank(wd+"2_unpacked")
	End Method

	End Type



Global d$[]=["lzw"]
Global JFD:JCR6_FAKEDRIVER

For Local dr$=EachIn d
	JFD=New JCR6_FAKEDRIVER
	JFD.name=dr
	RegisterCompDriver dr,JFD
	Print "JCR6 - Faked: "+dr
Next

