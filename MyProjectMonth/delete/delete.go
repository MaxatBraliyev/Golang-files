package delete

import (
    "os"
)

func DeleteFile(file string) {
    err := os.Remove(file)

    if err != nil {
        return
    }
}