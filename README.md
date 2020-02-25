# 🚀 Triggy

[![Actions Status](https://github.com/JulienBreux/triggy/workflows/Build%20and%20test%20Go/badge.svg)](https://github.com/JulienBreux/triggy/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/JulienBreux/triggy)](https://goreportcard.com/report/github.com/JulienBreux/triggy)
[![codebeat badge](https://codebeat.co/badges/f0d1ee44-fb81-4af7-9ff5-5759bf9d074e)](https://codebeat.co/projects/github-com-julienbreux-triggy-master)
[![GitHub tag](https://img.shields.io/github/tag/JulienBreux/triggy.svg)](Tag)
[![Go version](https://img.shields.io/github/go-mod/go-version/JulienBreux/triggy)](https://golang.org/dl/#stable)

Triggy is an awesome tool to interconnect some providers via adapters and trigger actions between them. It's simply written in Go.

---

## 🔧 Installation

Triggy is available on Linux, OSX and Windows platforms.

* Binaries for Mac OS, Linux and Windows are available as tarballs in the [release](https://github.com/JulienBreux/triggy/releases) page.

* Via Homebrew (Mac OS) or LinuxBrew (Linux)

   ```shell
   brew tap JulienBreux/triggy
   brew install triggy
   ```

* Building from source
   Triggy was built using go 1.13 or above. In order to build Triggy from source you must:
   1. Clone this repository
   2. Add the following command in your go.mod file

      ```text
      replace (
        github.com/JulienBreux/triggy => CLONED_GIT_REPOSITORY
      )
      ```

   3. Build and run the executable

        ```shell
        go build -v -o ./triggy ./cmd/triggy
        ```

   4. Use it

        ```shell
        ./triggy
        ```

---

## 📮 Contact Info

1. **Email**:   julien.breux@gmail.com
2. **GitHub**:  [@JulienBreux](https://github.com/JulienBreux)
3. **Twitter**: [@JulienBreux](https://twitter.com/JulienBreux)

---

## 👮‍♂️ Security info

### GPG Signature

You can download Julien Breux's public key to verify the signature.

```shell
gpg --keyserver hkps://hkps.pool.sks-keyservers.net --recv-keys 0BD023FA
```
