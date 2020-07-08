package routes

import (
	"fmt"
	"time"
	"github.com/go-vgo/robotgo"
)

var (
	pressing_left = false
	pressing_right = false
	pressing_up = false
	pressing_down = false
	pressing_left_press = false
	pressing_right_press = false
	pressing_scroll_up = false
	pressing_scroll_down = false
	prsssing_middle_press = false
)

func Press(code string) {
	switch code {
		case "left": 
			if !pressing_left {
				pressing_left = true
				go MoveLeft()
				fmt.Println("按下鼠标左移")
			}
		case "right": 
			if !pressing_right {
				pressing_right = true
				go MoveRight()
				fmt.Println("按下鼠标右移")
			}
		case "up": 
			if !pressing_up {
				pressing_up = true
				go MoveUp()
				fmt.Println("按下鼠标上移")
			}
		case "down": 
			if !pressing_down {
				pressing_down = true
				go MoveDown()
				fmt.Println("按下鼠标下移")
			}
		case "left_press":
			if !pressing_left_press {
				pressing_left_press = true
				left_press()
				fmt.Println("按下鼠标左键")
			}
		case "right_press":
			if !pressing_right_press {
				pressing_right_press = true
				right_press()
				fmt.Println("按下鼠标右键")
			}
		case "double_click":
			double_click()
			fmt.Println("按下鼠标右键")
		case "scroll_up":
			if !pressing_scroll_up {
				pressing_scroll_up = true
				scroll_up()
				fmt.Println("按下滑轮上滑")
			}
		case "scroll_down":
			if !pressing_scroll_down {
				pressing_scroll_down = true
				scroll_down()
				fmt.Println("按下滑轮下滑")
			}
		case "middle_press":
			if !prsssing_middle_press {
				prsssing_middle_press = true
				middle_press()
				fmt.Println("按下鼠标中键")
			}
	}
}

func Release(code string) {
	switch code {
		case "left": 
			if pressing_left {
				pressing_left = false
				fmt.Println("松开鼠标左移")
			}
		case "right": 
			if pressing_right {
				pressing_right = false
				fmt.Println("松开鼠标右移")
			}
		case "up": 
			if pressing_up {
				pressing_up = false
				fmt.Println("松开鼠标上移")
			}
			
		case "down": 
			if pressing_down {
				pressing_down = false
				fmt.Println("松开鼠标下移")
			}
		case "left_press":
			if pressing_left_press {
				pressing_left_press = false
				left_release()
				fmt.Println("松开鼠标左键")
			}
		case "right_press":
			if pressing_right_press {
				pressing_left_press = false
				right_release()
				fmt.Println("松开鼠标右键")
			}
		case "scroll_up" :
			if pressing_scroll_up {
				pressing_scroll_up = false
				fmt.Println("松开滑轮上滑")
			}
		case "scroll_down" :
			if pressing_scroll_down {
				pressing_scroll_down = false
				fmt.Println("松开滑轮下滑")
			}
		case "middle_press":
			if prsssing_middle_press {
				prsssing_middle_press = false
				middle_release()
				fmt.Println("松开鼠标中键")
			}
	}
}

func MoveLeft() {
	for pressing_left {
		time.Sleep(10 * time.Millisecond)
		x, y := robotgo.GetMousePos()
		x -= 2
		robotgo.MoveMouse(x, y)
	}
}

func MoveRight() {
	for pressing_right {
		time.Sleep(10 * time.Millisecond)
		x, y := robotgo.GetMousePos()
		x += 2
		robotgo.MoveMouse(x, y)
		
	}
}

func MoveUp() {
	for pressing_up {
		time.Sleep(10 * time.Millisecond)
		x, y := robotgo.GetMousePos()
		y -= 2
		robotgo.MoveMouse(x, y)
	}
}

func MoveDown() {
	for pressing_down {
		time.Sleep(10 * time.Millisecond)
		x, y := robotgo.GetMousePos()
		y += 2
		robotgo.MoveMouse(x, y)
	}
}

func left_press() {
	robotgo.MouseToggle("down", "left")
}

func right_press() {
	robotgo.MouseToggle("down", "right")
}

func left_release() {
	robotgo.MouseToggle("up", "left")
}

func right_release() {
	robotgo.MouseToggle("up", "right")
}

func double_click() {
	robotgo.MouseClick("left", true)
}

func scroll_up() {
	if (pressing_scroll_up) {
		time.Sleep(10 * time.Millisecond)
		robotgo.ScrollMouse(50, "up")
	}
}

func scroll_down() {
	if (pressing_scroll_down) {
		time.Sleep(10 * time.Millisecond)
		robotgo.ScrollMouse(50, "down")
	}
}

func middle_press() {
	robotgo.MouseToggle("down", "center")
}

func middle_release() {
	robotgo.MouseToggle("up", "center")
}