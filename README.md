# ChainMaker

## 脚本搭建

1. 进入工作目录: `cd ${WORKDIR}`
2. 克隆项目: `git clone -b v2.2.1 https://git.chainmaker.org.cn/chainmaker/chainmaker-go.git`
3. 克隆项目: `git clone -b v2.2.0 https://git.chainmaker.org.cn/chainmaker/chainmaker-cryptogen.git`
4. 编译: `cd chainmaker-cryptogen && make`
5. 软链接(使用绝对路径): `ln -s ${WORKDIR}/chainmaker-cryptogen ${WORKDIR}/chainmaker-go/tools`
6. 生成单链4节点证书配置文件:
   - `cd ${WORKDIR}/chainmaker-go/scripts`
   - 全部默认回车: `./prepare.sh 4 1`
   - `tree -L 3 ../build/`
7. 编译安装包:
   - `./build_release.sh`
   - `tree ../build/release/`
8. 运行节点集群:
   - 启动: `./cluster_quick_start.sh normal`
   - 查看进程: `ps -ef|grep chainmaker | grep -v grep`
   - 查看端口:
     - (Linux)`netstat -lptn | grep 1230`
     - (macOS)`netstat -an | grep 1230`
9. 编译命令行工具(可跳过):
   - `cd ${WORKDIR}/chainmaker-go/tools/cmc`
   - `go build`
   - `cp -rf $WORKDIR/chainmaker-go/build/crypto-config $WORKDIR/chainmaker-go/tools/cmc/testdata/`
10. 安装Go-SDK: 
    - `git clone -b v2.2.1 https://git.chainmaker.org.cn/chainmaker/sdk-go.git`
    - `cd ${WORKDIR}/sdk-go/testdata`
    - `rm -rf crypto-config`
    - 软链接(使用绝对路径): `ln -s ${WORKDIR}/chainmaker-go/build/crypto-config ${WORKDIR}/sdk-go/testdata`
    - 测试存证合约: `cd ${WORKDIR}/sdk-go/examples/user_contract_claim && go run main.go`
  
## Docker 搭建
- 待更