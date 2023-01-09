package banner

import "fmt"

var banner string = `
				DIRECTORY MAPPER
   				    ::by ERN::
		   :: github.com/Lukmanern/DirectoryMapper ::   
`

func ShowBanner() {
	fmt.Println(banner)
}