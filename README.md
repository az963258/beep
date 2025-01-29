# beep
A very simple Golang project that beeps your pc. That's all

it sends "alert" symbol to tty1 (```echo -en "\a" > /dev/tty1```), and that symbol makes your pc beep

tested on linux, not sure if it works on windows

# how to build
```git clone https://github.com/az963258/beep.git/```

``` cd beep```

```go build main.go```

or

```git clone https://github.com/az963258/beep.git/ && cd beep && go build main.go && mv main gobeep && sudo cp gobeep /usr/bin/ && echo "success, now you can launch it (just type gobeep)"```
