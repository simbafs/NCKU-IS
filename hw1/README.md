# NCKU-IS-HW1
NCKU Information Security HW1

# 原理
## 藏入
輸入 `secret` 和 `carriers`，`secret` 指要隱藏的字串，只能包含 `a-zA-Z` 和空白（`\u0020`），`carriers` 是一串用空白分隔的字串。
1. 將 `secret` 按照 ascii code 和 `spaceMap` 編碼成兩個空白，獲得一串空白字元 `secretSpace`
2. 將 `carriers` 用空白切成一串 `[]string`，稱為 `carrierList`
3. 依序將一個 `carrier`、一個空白串起來，直到 `secretSpace` 用完，在結尾在插入最後一個 `carrier`，如果 `carrier` 不夠用，則從頭開開始
## 取出
取出所有空白字元，根據 `spaceMap` 解碼，就獲得原本的 `secret` 了
