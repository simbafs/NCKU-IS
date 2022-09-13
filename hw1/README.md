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
Enter what you want to hide, only ascii availabe: ncku infor
mation security
Enter the carrier text: sdf jdfk fsdkl ruei cnxzm
```
## 取出
```
$ cat t | go run . extract
ncku information security%
```

## t 內容
```
00000000: 7364 66e2 8084 6a64 666b e281 9f66 7364  sdf...jdfk...fsd
00000010: 6b6c e280 8472 7565 69e2 8081 636e 787a  kl...ruei...cnxz
00000020: 6de2 8084 7364 66e2 8089 6a64 666b e280  m...sdf...jdfk..
00000030: 8566 7364 6b6c e280 8372 7565 69e2 8080  .fsdkl...ruei...
00000040: 636e 787a 6d20 7364 66e2 8084 6a64 666b  cnxzm sdf...jdfk
00000050: e280 8766 7364 6b6c e280 8472 7565 69e2  ...fsdkl...ruei.
00000060: 819f 636e 787a 6de2 8084 7364 66e2 8084  ..cnxzm...sdf...
00000070: 6a64 666b e280 8466 7364 6b6c e380 8072  jdfk...fsdkl...r
00000080: 7565 69e2 8085 636e 787a 6de2 8080 7364  uei...cnxzm...sd
00000090: 66e2 8084 6a64 666b e280 af66 7364 6b6c  f...jdfk...fsdkl
000000a0: e280 8472 7565 69c2 a063 6e78 7a6d e280  ...ruei..cnxzm..
000000b0: 8573 6466 e280 826a 6466 6be2 8084 6673  .sdf...jdfk...fs
000000c0: 646b 6ce2 8087 7275 6569 e280 8463 6e78  dkl...ruei...cnx
000000d0: 7a6d e380 8073 6466 e280 846a 6466 6be2  zm...sdf...jdfk.
000000e0: 819f 6673 646b 6ce2 8080 7275 6569 2063  ..fsdkl...ruei c
000000f0: 6e78 7a6d e280 8573 6466 e280 816a 6466  nxzm...sdf...jdf
00000100: 6be2 8084 6673 646b 6ce2 8083 7275 6569  k...fsdkl...ruei
00000110: e280 8463 6e78 7a6d e280 8173 6466 e280  ...cnxzm...sdf..
00000120: 856a 6466 6be2 8083 6673 646b 6ce2 8085  .jdfk...fsdkl...
00000130: 7275 6569 e280 8063 6e78 7a6d e280 8473  ruei...cnxzm...s
00000140: 6466 e280 876a 6466 6be2 8085 6673 646b  df...jdfk...fsdk
00000150: 6ce2 8082 7275 6569 e280 8563 6e78 7a6d  l...ruei...cnxzm
00000160: e280 8773 6466 0a                        ...sdf.
```
