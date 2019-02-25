package handler

import (
    "io"
    "os"
    "os/exec"
    "fmt"
    "net/http"

    "github.com/labstack/echo"
    "github.com/satori/go.uuid"
)

type MarkdownData struct {
    MD string `json:"md" form:"md" query:"md"`
}

func Default(c echo.Context) (err error) {
    md := new(MarkdownData)
    if err = c.Bind(md); err != nil {
        return
    }

    uu, err := uuid.NewV4()
    if err != nil {
        return err
    }

    dst, err := os.Create("mdfiles/" + uu.String())
    if err != nil {
        return err
    }

    dst.Write(([]byte) (md.MD))

    out, err := exec.Command("python", "markdown.py", "mdfiles/" + uu.String()).Output()
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

    uu, err := uuid.NewV4()
    if err != nil {
        return err
    }

    dst, err := os.Create("mdfiles/" + uu.String())
    if err != nil {
        return err
    }

    if _, err = io.Copy(dst, src); err != nil {
        return err
    }

    out, err := exec.Command("python", "markdown.py", "mdfiles/" + uu.String()).Output()
    if err != nil {
        return err
    }

    return c.HTML(http.StatusOK, fmt.Sprintf("%s", out))
}
