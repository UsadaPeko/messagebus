# MessageBus

Simple Message Bus Written in Golang

## How to Use

```bash
go get gopkg.io/UsadaPeko/messagebus
```

## Why?

Event driven architecture에서는 이벤트 발생시에 내부 리스너들에게 전달해야 할 필요가 있습니다. 이 과정에서 외부가 아닌 내부 컴포넌트들에게
간단하게 메세지를 전달하는 경우 이 Message Bus를 사용할 수 있습니다.
