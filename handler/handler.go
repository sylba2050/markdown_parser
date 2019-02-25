package handler

import (
    "io"
    "os"
    "fmt"
    "net/http"
    "github.com/labstack/echo"
    // "os/exec"
)

type MarkdownData struct {
    MD string `json:"md" form:"md" query:"md"`
}

func Default(c echo.Context) (err error) {
    md := new(MarkdownData)
    if err = c.Bind(md); err != nil {
        return
    }
	
    return c.JSON(http.StatusOK, md.MD)
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
	
    return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully", file.Filename))
}
