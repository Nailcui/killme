package killme

import (
	"fmt"
	"os"
)

const fileName string = "killme.sh"

func init() {
	pid := os.Getpid()
	f, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
	content := fmt.Sprintf("#!/bin/bash\nkill -9 %d", pid)
	f.WriteString(content)
	f.Close()
}
