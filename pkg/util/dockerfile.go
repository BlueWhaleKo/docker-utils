package util

import (
	"fmt"
	"strings"
)

type Dockerfile struct {
	inner []string
}

func NewDockerfile() *Dockerfile {
	return &Dockerfile{}
}

func (d *Dockerfile) From(from string) *Dockerfile {
	d.inner = append(d.inner, fmt.Sprintf("FROM %s", from))
	return d
}

func (d *Dockerfile) Entrypoint(entrypoint string) *Dockerfile {
	d.inner = append(d.inner, fmt.Sprintf("ENTRYPOINT %s", entrypoint))
	return d
}

func (d *Dockerfile) Cmd(cmd string) *Dockerfile {
	d.inner = append(d.inner, fmt.Sprintf("CMD %s", cmd))
	return d
}

func (d *Dockerfile) WORKDIR(dir string) *Dockerfile {
	d.inner = append(d.inner, fmt.Sprintf("WORKDIR %s", dir))
	return d
}

func (d *Dockerfile) Run(cmd string) *Dockerfile {
	d.inner = append(d.inner, fmt.Sprintf("RUN %s", cmd))
	return d
}
func (d *Dockerfile) Add(source, target string) *Dockerfile {
	d.inner = append(d.inner, fmt.Sprintf("ADD %s %s", source, target))
	return d
}

func (d *Dockerfile) Copy(source, target string) *Dockerfile {
	d.inner = append(d.inner, fmt.Sprintf("COPY %s %s", source, target))
	return d
}

func (d *Dockerfile) Build() string {
	return strings.Join(d.inner, "\n")
}
