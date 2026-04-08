package cli

import "fmt"

const banner = `
 __      __   ________  ____  ___
 \ \    / /  |  ______| \   \/  /
  \ \  / /   | |__       \     / 
   \ \/ /    |  __|      /     \ 
    \  /     | |______  /   /\  \
     \/      |________|/___/  \__\
                                   
      Stateful API Logic Breaker   
            [ Zero-Noise ]         
`

func PrintBanner() {
	fmt.Println(banner)
}
