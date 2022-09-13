# NCKU-IS-HW1
NCKU Information Security HW1

[pdf](./out.pdf)

# 原理
## 藏入
輸入 `secret` 和 `carriers`，`secret` 指要隱藏的字串，只能包含 ASCII，`carriers` 是一串用空白分隔的字串。
1. 將 `secret` 按照 ascii code 和 `spaceMap` 編碼成兩個空白，獲得一串空白字元 `secretSpace`
2. 將 `carriers` 用空白切成一串 `[]string`，稱為 `carrierList`
3. 依序將一個 `carrier`、一個空白串起來，直到 `secretSpace` 用完，在結尾在插入最後一個 `carrier`，如果 `carrier` 不夠用，則從頭開開始
## 取出
取出所有空白字元，根據 `spaceMap` 解碼，就獲得原本的 `secret` 了

# usage
## 隱藏
```
$ go run . hide > t
Enter what you want to hide, only ascii availabe: ncku information security
Enter the carrier text: this is a message that may repeat many times
```
## 取出
```
$ cat t | go run . extract
ncku information security%                                                      
```

## t 內容
```
00000000: 7468 6973 e280 8469 73e2 819f 61e2 8084  this...is...a...
00000010: 6d65 7373 6167 65e2 8081 7468 6174 e280  message...that..
00000020: 846d 6179 e280 8972 6570 6561 74e2 8085  .may...repeat...
00000030: 6d61 6e79 e280 8374 696d 6573 e280 8074  many...times...t
00000040: 6869 7320 6973 e280 8461 e280 876d 6573  his is...a...mes
00000050: 7361 6765 e280 8474 6861 74e2 819f 6d61  sage...that...ma
00000060: 79e2 8084 7265 7065 6174 e280 846d 616e  y...repeat...man
00000070: 79e2 8084 7469 6d65 73e3 8080 7468 6973  y...times...this
00000080: e280 8569 73e2 8080 61e2 8084 6d65 7373  ...is...a...mess
00000090: 6167 65e2 80af 7468 6174 e280 846d 6179  age...that...may
000000a0: c2a0 7265 7065 6174 e280 856d 616e 79e2  ..repeat...many.
000000b0: 8082 7469 6d65 73e2 8084 7468 6973 e280  ..times...this..
000000c0: 8769 73e2 8084 61e3 8080 6d65 7373 6167  .is...a...messag
000000d0: 65e2 8084 7468 6174 e281 9f6d 6179 e280  e...that...may..
000000e0: 8072 6570 6561 7420 6d61 6e79 e280 8574  .repeat many...t
000000f0: 696d 6573 e280 8174 6869 73e2 8084 6973  imes...this...is
00000100: e280 8361 e280 846d 6573 7361 6765 e280  ...a...message..
00000110: 8174 6861 74e2 8085 6d61 79e2 8083 7265  .that...may...re
00000120: 7065 6174 e280 856d 616e 79e2 8080 7469  peat...many...ti
00000130: 6d65 73e2 8084 7468 6973 e280 8769 73e2  mes...this...is.
00000140: 8085 61e2 8082 6d65 7373 6167 65e2 8085  ..a...message...
00000150: 7468 6174 e280 876d 6179 0a              that...may.
```
