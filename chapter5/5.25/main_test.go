package main

import (
	"bytes"
	"html/template"
	"sync"
	"testing"
)

func BenchmarkParallelOops(b *testing.B) {
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}

func BenchmarkParallelOk(b *testing.B) {
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	var mutex sync.Mutex
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mutex.Lock()
			t.Execute(&buf, data)
			buf.Reset()
			mutex.Unlock()
		}
	})
}