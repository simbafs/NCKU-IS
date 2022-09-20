# NCKU-IS-HW2
NCKU Information Security HW2
[pdf](./out.pdf)

# 原理
這是一個對稱式金鑰加密，金鑰為六個整數

## 加密
金鑰的六個數字依序為 `a` `b` `c` `d` `e` `f`，明文的每兩個字為 `x` `y`，帶入 `X=ax+by+c` `Y=dx+ey+f`，得到的數字再用 62 進位編碼（`0-9a-ZA-Z`）。在輸出的時候再加上每個字加密後的位數用 62 位元編碼當前綴，例如 `4Y` `8R` 會各加上 `2` 後再串起來，變成 `24Y28R`，全部串起來就得到密文

## 解密
依照位數前綴將密文差開、轉成 10 進位用克拉瑪公式解出 `x` `y`，就得到原本的明文

# 範例
## 加密
```
$ go run . encrypt 1 2 3 4 5 6 'ncku information security'
2512fb25y2gr23X2aD2572fq25w2gs24W2eX25j2g325o2g824h2br24S2eB25C2gQ25u2ge2322au
```

## 解密
```
$ go run . decrypt 1 2 3 4 5 6 '2512fb25y2gr23X2aD2572fq25w2gs24W2eX25j2g325o2g824h2br24S2eB25C2gQ25u2ge2322au'
ncku information security
```

# 缺陷
因為他還是一對一的加密，所以破解相對容易，未來可以加上交換或是其他方式改善
