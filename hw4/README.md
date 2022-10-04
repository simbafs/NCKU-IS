# NCKU-IS-HW4
修改 `main.go:14` 開始的幾個參數
| name   | description               |
| :---:  | :---                      |
| ct     | ciphertext                |
| bs     | block size                |
| prefix | 明文的前綴                |
| max    | 密鑰測試的最大值          |
| fo     | file option，輸出結果用的 |
```
go run .
```

## 多核心加速
我嘗試用 goroutine 多執行緒加速，但預估時間只從 18 分鐘變成 16 分鐘，我猜應該是某個地方些錯了。
