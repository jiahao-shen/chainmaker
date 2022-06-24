# ChainMaker

## TODO
- 完善合约, 添加key

## 基于脚本搭建

### 长安链
1. 进入工作目录: `cd ${WORKDIR}`
2. 克隆项目: `git clone -b v2.2.0 https://git.chainmaker.org.cn/chainmaker/chainmaker-go.git`
3. 克隆项目: `git clone -b v2.2.0 https://git.chainmaker.org.cn/chainmaker/chainmaker-cryptogen.git`
4. 编译: `cd chainmaker-cryptogen && make`
5. 软链接(使用绝对路径): `ln -s ${WORKDIR}/chainmaker-cryptogen ${WORKDIR}/chainmaker-go/tools`
6. 生成单链单节点证书配置文件:
   - `cd ${WORKDIR}/chainmaker-go/scripts`
   - `./prepare.sh 1 1`
     - `input consensus type (0-SOLO,1-TBFT(default),3-MAXBFT,4-RAFT): 0`
     - `input log level (DEBUG|INFO(default)|WARN|ERROR): INFO`
     - `enable docker vm (YES|NO(default))YES`
   - `tree -L 3 ../build`
7. 编译安装包:
   - `./build_release.sh`
   - `tree ../build/release`
8. 解压安装包并运行:
   - `cd ../build/release`
   - `tar zxvf chainmaker-xxxx.tar.gz`
   - `docker pull chainmakerofficial/chainmaker-vm-docker-go:v2.2.0.1`
   - `cd chainmaker-xxxx/bin && ./start.sh`

<!-- 8. 运行节点集群:
   - 启动: `./cluster_quick_start.sh normal`
   - 查看进程: `ps -ef|grep chainmaker | grep -v grep`
   - 查看端口:
     - (Linux)`netstat -lptn | grep 1230`
     - (macOS)`netstat -an | grep 1230` -->

### 命令行工具(暂时跳过)
1. `cd ${WORKDIR}/chainmaker-go/tools/cmc`
2. `go build`
3. `cp -rf $WORKDIR/chainmaker-go/build/crypto-config $WORKDIR/chainmaker-go/tools/cmc/testdata/`

### 安装Go-SDK(暂时跳过): 
1. 克隆项目: `git clone -b v2.2.1 https://git.chainmaker.org.cn/chainmaker/sdk-go.git`
2. `cd ${WORKDIR}/sdk-go/testdata`
3. `rm -rf crypto-config`
4. 软链接(使用绝对路径): `ln -s ${WORKDIR}/chainmaker-go/build/crypto-config ${WORKDIR}/sdk-go/testdata`
5. 测试存证合约: `cd ${WORKDIR}/sdk-go/examples/user_contract_claim && go run main.go`