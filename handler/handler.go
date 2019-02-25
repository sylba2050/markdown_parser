package handler

import (
    "io"
    "os"
    "fmt"
    "net/http"
    "github.com/labstack/echo"
    "os/exec"
)

type MarkdownData struct {
    MD string `json:"md" form:"md" query:"md"`
}

func Default(c echo.Context) (err error) {
    md := new(MarkdownData)
    if err = c.Bind(md); err != nil {
        return
    }

    dst, err := os.Create("mdfiles/002.md")
    if err != nil {
        return err
    }

    dst.Write(([]byte) (md.MD))
	
    out, err := exec.Command("python", "markdown.py", "mdfiles/002.md").Output()
    if err != nil {
        return err
    }

    return c.HTML(http.StatusOK, fmt.Sprintf("%s", out))
}


func File(c echo.Context) error {
    file, err := c.FormFile("file")
    if err != nil {
        return err
    }

    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    dst, err := os.Create("mdfiles/001.md")
    if err != nil {
        return err
    }

    if _, err = io.Copy(dst, src); err != nil {
        return err
    }
	
    out, err := exec.Command("python", "markdown.py", "mdfiles/001.md").Output()
    if err != nil {
        return err
    }

    return c.HTML(http.StatusOK, fmt.Sprintf("%s", out))
}
