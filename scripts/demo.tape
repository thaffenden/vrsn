Output demo.gif
Require git
Set Theme "OneDark"
Set Height 1000

Hide
Type "alias vrsn=$(pwd)/vrsn"
Enter
Type "clear"
Enter
Show

Set TypingSpeed 0.2s
Type "vrsn"
Enter 2
Sleep 2s

Type "vrsn check"
Enter 2
Sleep 2s

Type "vrsn bump"
Enter
Sleep 1s
Down 1
Sleep 0.5s
Down 1
Sleep 0.5s
Up 1
Sleep 0.5s
Up 1
Sleep 1s
Enter 2

Type "vrsn check"
Enter
Sleep 1s

Set TypingSpeed 0s
Hide
Type "git checkout ./VERSION"
Enter
Sleep 2s
