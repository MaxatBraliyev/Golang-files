package mkrdir

import (
    "os"
	"path/filepath"
	"time"
)
func MkrDir(path string){
newdir := time.Now().Format("02.01.2006")

//Create output path
outPath:= filepath.Join(path, newdir)

//Create dir output using above code
if _, err := os.Stat(outPath); os.IsNotExist(err) {
    os.Mkdir(outPath, 0755)
}
}