package main

import (
	"fmt"
)

var (
	system1 string = `
  _______     _______ _______ ______ __  __ 
 / ____\ \   / / ____|__   __|  ____|  \/  |
| (___  \ \_/ / (___    | |  | |__  | \  / |
 \___ \  \   / \___ \   | |  |  __| | |\/| |
 ____) |  | |  ____) |  | |  | |____| |  | |
|_____/   |_| |_____/   |_|  |______|_|  |_|
`
	wants1 string = `
 __          __         _       
 \ \        / /        | |      
  \ \  /\  / /_ _ _ __ | |_ ___ 
   \ \/  \/ / _' | '_ \| __/ __|
    \  /\  / (_| | | | | |_\__ \
     \/  \/ \__,_|_| |_|\__|___/
`
	to1 string = `
 _        
| |       
| |_ ___  
| __/ _ \ 
| || (_) |
 \__\___/ 
`
	keep1 string = `
  __    __  ________  ________  _______  
 |  \  /  \|        \|        \|       \ 
 | $$ /  $$| $$$$$$$$| $$$$$$$$| $$$$$$$\
 | $$/  $$ | $$__    | $$__    | $$__/ $$
 | $$  $$  | $$  \   | $$  \   | $$    $$
 | $$$$$\  | $$$$$   | $$$$$   | $$$$$$$ 
 | $$ \$$\ | $$_____ | $$_____ | $$      
 | $$  \$$\| $$     \| $$     \| $$      
  \$$   \$$ \$$$$$$$$ \$$$$$$$$ \$$      
`
	you1 string = `
 __      __   ______   __    __ 
|  \    /  \ /      \ |  \  |  \
 \$$\  /  $$|  $$$$$$\| $$  | $$
  \$$\/  $$ | $$  | $$| $$  | $$
   \$$  $$  | $$  | $$| $$  | $$
    \$$$$   | $$  | $$| $$  | $$
    | $$    | $$__/ $$| $$__/ $$
    | $$     \$$    $$ \$$    $$
     \$$      \$$$$$$   \$$$$$$ 
`
	low1 string = `
  __        ______   __       __ 
 |  \      /      \ |  \  _  |  \
 | $$     |  $$$$$$\| $$ / \ | $$
 | $$     | $$  | $$| $$/  $\| $$
 | $$     | $$  | $$| $$  $$$\ $$
 | $$     | $$  | $$| $$ $$\$$\$$
 | $$_____| $$__/ $$| $$$$  \$$$$
 | $$     \\$$    $$| $$$    \$$$
  \$$$$$$$$ \$$$$$$  \$$      \$$
`

	its1 string = `
    _______________
   /  _/_  __/ ___/
   / /  / /  \__ \ 
 _/ /  / /  ___/ / 
/___/ /_/  /____/  
`

	gonna1 string = `
     __________  _   ___   _____ 
    / ____/ __ \/ | / / | / /   |
   / / __/ / / /  |/ /  |/ / /| |
  / /_/ / /_/ / /|  / /|  / ___ |
  \____/\____/_/ |_/_/ |_/_/  |_|
`

	leave1 string = `


$$\       $$$$$$$$\  $$$$$$\  $$\    $$\ $$$$$$$$\ 
$$ |      $$  _____|$$  __$$\ $$ |   $$ |$$  _____|
$$ |      $$ |      $$ /  $$ |$$ |   $$ |$$ |      
$$ |      $$$$$\    $$$$$$$$ |\$$\  $$  |$$$$$\    
$$ |      $$  __|   $$  __$$ | \$$\$$  / $$  __|   
$$ |      $$ |      $$ |  $$ |  \$$$  /  $$ |      
$$$$$$$$\ $$$$$$$$\ $$ |  $$ |   \$  /   $$$$$$$$\ 
\________|\________|\__|  \__|    \_/    \________|
                                                   
`

	leaveyou1 string = `
  /$$       /$$$$$$$$  /$$$$$$  /$$    /$$ /$$$$$$$$       /$$     /$$ /$$$$$$  /$$   /$$
 | $$      | $$_____/ /$$__  $$| $$   | $$| $$_____/      |  $$   /$$//$$__  $$| $$  | $$
 | $$      | $$      | $$  \ $$| $$   | $$| $$             \  $$ /$$/| $$  \ $$| $$  | $$
 | $$      | $$$$$   | $$$$$$$$|  $$ / $$/| $$$$$           \  $$$$/ | $$  | $$| $$  | $$
 | $$      | $$__/   | $$__  $$ \  $$ $$/ | $$__/            \  $$/  | $$  | $$| $$  | $$
 | $$      | $$      | $$  | $$  \  $$$/  | $$                | $$   | $$  | $$| $$  | $$
 | $$$$$$$$| $$$$$$$$| $$  | $$   \  $/   | $$$$$$$$          | $$   |  $$$$$$/|  $$$$$$/
 |________/|________/|__/  |__/    \_/    |________/          |__/    \______/  \______/ 
`

	leaveyouwith1 string = `
  __        ________   ______   __     __  ________        __      __   ______   __    __ 
 |  \      |        \ /      \ |  \   |  \|        \      |  \    /  \ /      \ |  \  |  \
 | $$      | $$$$$$$$|  $$$$$$\| $$   | $$| $$$$$$$$       \$$\  /  $$|  $$$$$$\| $$  | $$
 | $$      | $$__    | $$__| $$| $$   | $$| $$__            \$$\/  $$ | $$  | $$| $$  | $$
 | $$      | $$  \   | $$    $$ \$$\ /  $$| $$  \            \$$  $$  | $$  | $$| $$  | $$
 | $$      | $$$$$   | $$$$$$$$  \$$\  $$ | $$$$$             \$$$$   | $$  | $$| $$  | $$
 | $$_____ | $$_____ | $$  | $$   \$$ $$  | $$_____           | $$    | $$__/ $$| $$__/ $$
 | $$     \| $$     \| $$  | $$    \$$$   | $$     \          | $$     \$$    $$ \$$    $$
  \$$$$$$$$ \$$$$$$$$ \$$   \$$     \$     \$$$$$$$$           \$$      \$$$$$$   \$$$$$$ 
  __       __  ______  ________  __    __ 
 |  \  _  |  \|      \|        \|  \  |  \
 | $$ / \ | $$ \$$$$$$ \$$$$$$$$| $$  | $$
 | $$/  $\| $$  | $$     | $$   | $$__| $$
 | $$  $$$\ $$  | $$     | $$   | $$    $$
 | $$ $$\$$\$$  | $$     | $$   | $$$$$$$$
 | $$$$  \$$$$ _| $$_    | $$   | $$  | $$
 | $$$    \$$$|   $$ \   | $$   | $$  | $$
  \$$      \$$ \$$$$$$    \$$    \$$   \$$
`

	leaveyouwithnowhere1 string = `
  __        ________   ______   __     __  ________        __      __  ______   __    __ 
 /  |      /        | /      \ /  |   /  |/        |      /  \    /  |/      \ /  |  /  |
 $$ |      $$$$$$$$/ /$$$$$$  |$$ |   $$ |$$$$$$$$/       $$  \  /$$//$$$$$$  |$$ |  $$ |
 $$ |      $$ |__    $$ |__$$ |$$ |   $$ |$$ |__           $$  \/$$/ $$ |  $$ |$$ |  $$ |
 $$ |      $$    |   $$    $$ |$$  \ /$$/ $$    |           $$  $$/  $$ |  $$ |$$ |  $$ |
 $$ |      $$$$$/    $$$$$$$$ | $$  /$$/  $$$$$/             $$$$/   $$ |  $$ |$$ |  $$ |
 $$ |_____ $$ |_____ $$ |  $$ |  $$ $$/   $$ |_____           $$ |   $$ \__$$ |$$ \__$$ |
 $$       |$$       |$$ |  $$ |   $$$/    $$       |          $$ |   $$    $$/ $$    $$/ 
 $$$$$$$$/ $$$$$$$$/ $$/   $$/     $/     $$$$$$$$/           $$/     $$$$$$/   $$$$$$/  
 __       __  ______  ________  __    __        __    __   ______   __       __  __    __  ________  _______   ________ 
/  |  _  /  |/      |/        |/  |  /  |      /  \  /  | /      \ /  |  _  /  |/  |  /  |/        |/       \ /        |
$$ | / \ $$ |$$$$$$/ $$$$$$$$/ $$ |  $$ |      $$  \ $$ |/$$$$$$  |$$ | / \ $$ |$$ |  $$ |$$$$$$$$/ $$$$$$$  |$$$$$$$$/ 
$$ |/$  \$$ |  $$ |     $$ |   $$ |__$$ |      $$$  \$$ |$$ |  $$ |$$ |/$  \$$ |$$ |__$$ |$$ |__    $$ |__$$ |$$ |__    
$$ /$$$  $$ |  $$ |     $$ |   $$    $$ |      $$$$  $$ |$$ |  $$ |$$ /$$$  $$ |$$    $$ |$$    |   $$    $$< $$    |   
$$ $$/$$ $$ |  $$ |     $$ |   $$$$$$$$ |      $$ $$ $$ |$$ |  $$ |$$ $$/$$ $$ |$$$$$$$$ |$$$$$/    $$$$$$$  |$$$$$/    
$$$$/  $$$$ | _$$ |_    $$ |   $$ |  $$ |      $$ |$$$$ |$$ \__$$ |$$$$/  $$$$ |$$ |  $$ |$$ |_____ $$ |  $$ |$$ |_____ 
$$$/    $$$ |/ $$   |   $$ |   $$ |  $$ |      $$ | $$$ |$$    $$/ $$$/    $$$ |$$ |  $$ |$$       |$$ |  $$ |$$       |
$$/      $$/ $$$$$$/    $$/    $$/   $$/       $$/   $$/  $$$$$$/  $$/      $$/ $$/   $$/ $$$$$$$$/ $$/   $$/ $$$$$$$$/ 
`

	togo1 string = `
  ______   ___        ____   ___  
 |      | /   \      /    | /   \ 
 |      ||     |    |   __||     |
 |_|  |_||  O  |    |  |  ||  O  |
   |  |  |     |    |  |_ ||     |
   |  |  |     |    |     ||     |
   |__|   \___/     |___,_| \___/ 
`
)

func chorusBig1() {
	noteRest(eightnote)
	fmt.Print(system1)
	noteRest(halfnote)
	cleanDisplay()
	fmt.Print(wants1)
	noteRest(eightnote)
	fmt.Print(to1)
	noteRest(eightnote)
	cleanDisplay()
	fmt.Print(keep1)
	noteRest(eightnote)
	fmt.Print(you1)
	noteRest(quaternote)
	fmt.Print(low1)
	noteRest(quaternote)
	cleanDisplay()
	fmt.Print(its1)
	noteRest(eightnote)
	fmt.Print(gonna1)
	noteRest(quaternote)
	cleanDisplay()
	fmt.Print(leave1)
	noteRest(quaternote)
	cleanDisplay()
	fmt.Print(leaveyou1)
	noteRest(quaternote)
	cleanDisplay()
	fmt.Print(leaveyouwith1)
	noteRest(quaternote)
	cleanDisplay()
	fmt.Print(leaveyouwithnowhere1)
	noteRest(halfnote)
	cleanDisplay()
	fmt.Print(togo1)
	noteRest(halfnote + quaternote)
	cleanDisplay()
	noteRest(eightnote)
}

func chorusSmall1() {
	noteRest(eightnote)
	fmt.Print(system1)
	noteRest(halfnote)
	cleanDisplay()
	fmt.Print(wants1)
	noteRest(eightnote)
	fmt.Print(to1)
	noteRest(eightnote)
	cleanDisplay()
	fmt.Print(keep1)
	noteRest(eightnote)
	fmt.Print(you1)
	noteRest(quaternote)
	fmt.Print(low1)
	noteRest(quaternote)
	cleanDisplay()
}
