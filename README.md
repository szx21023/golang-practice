## Linux install go:
### 1. 下載壓縮檔（.tar.gz）：

```plaintext
wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz
```
### 2. 解壓並安裝到 /usr/local：

```plaintext
sudo tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
```

### 3. 在 ~/.bashrc 或 ~/.zshrc 加入以下設定(macOS / Linux)：

```plaintext
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

### 4. 儲存後執行：

```plaintext
source ~/.zshrc    # 或 source ~/.bashrc
```

### 5. 確認是否安裝成功

```plaintext
go version
```