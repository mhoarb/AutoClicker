# Autoclicker

This is a simple autoclicker program written in Go using the `robotgo` package. The program continuously clicks the mouse at the current cursor position every 80 microseconds.

## Prerequisites

- Go (Golang) installed on your machine. You can download it from [golang.org](https://golang.org/dl/).
- The `robotgo` package. You can install it using the following command:
  ```sh
  go get -u github.com/go-vgo/robotgo
  ```
Clone the repository (if you are using version control)

```bash
git clone https://github.com/yourusername/autoclicker.git
cd autoclicker
```
install the required dependencies for robotgo

On Ubuntu, you need to install the X11 development libraries:
```bash
sudo apt-get update
sudo apt-get install libx11-dev libx11-xcb-dev libxtst-dev libxrender-dev libxinerama-dev
```
