package gologs

import (
	"errors"
)

const reset = "\033[0m"
const Red = Color("\033[31m")
const Greed = Color("\033[32m")
const Yellow = Color("\033[33m")
const Blue = Color("\033[34m")
const Purple = Color("\033[35m")
const Cyan = Color("\033[36m")
const Gray = Color("\033[37m")
const White = Color("\033[97m")

type Color string

type colors struct {
	Info    string
	Warning string
	Error   string
	Fatal   string
}

func NewColors() *colors {
	return &colors{
		Info:    string(Blue),
		Warning: string(Yellow),
		Error:   string(Red),
		Fatal:   string(Red),
	}
}

func (c *colors) SetError(newColor Color) error {
	var err error

	if newColor == "" {
		err = errors.New("Not defined color")
	}

	c.Error = string(newColor)

	return err
}

func (c *colors) SetWarn(newColor Color) error {
	var err error

	if newColor == "" {
		err = errors.New("Not defined color")
	}

	c.Warning = string(newColor)

	return err
}

func (c *colors) SetFatal(newColor Color) error {
	var err error

	if newColor == "" {
		err = errors.New("Not defined color")
	}

	c.Fatal = string(newColor)

	return err
}

func (c *colors) SetInfo(newColor Color) error {
	var err error

	if newColor == "" {
		err = errors.New("Not defined color")
	}

	c.Info = string(newColor)

	return err
}
