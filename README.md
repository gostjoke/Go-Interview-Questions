# Golang Interview Questions & Notes

A curated collection of **Golang interview questions**, explanations, and code snippets.  
This repository is designed for **backend engineers**, **system developers**, and anyone preparing for **Go technical interviews**.

The focus is not only on *what* the answer is, but also on *why* it works.

---

## 🎯 Goals

- Cover common **Go interview questions**
- Explain **core concepts clearly**
- Emphasize **real-world reasoning**, not rote memorization
- Serve as a **long-term personal knowledge base**

---

## 📚 Topics Covered

- Go basics & language design
- Memory management & escape analysis
- Stack vs heap
- Pointers, values, and references
- Slices, maps, and strings
- Interfaces & type system
- Goroutines & channels
- Synchronization primitives
- Garbage collection & performance
- Common pitfalls & best practices

---

## 🧠 How to Use This Repo

- Read by topic when preparing for interviews
- Search specific questions quickly
- Add your own notes and examples
- Use code snippets to experiment locally

This repo is **opinionated and practical**, based on real interview experiences.

---

## 🛠 Requirements

- Go 1.20+
- Basic understanding of programming concepts

---

## Each Folder and content

### Defer 使用與時機
Save in Stack, LIFO

### EscapeAnalysis
Research the Memory Escape and example
研究Golang 的內存逃逸狀況與案例

### Golang Redis Spin Lock
Redis 自旋鎖 = 拿不到鎖就一直嘗試（sleep + retry），直到成功或超時

### Panic and defer and recover
panic 發生
↓
執行 defer
↓
defer 中呼叫 recover() 才能捕捉